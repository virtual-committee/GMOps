package logic

import (
	"fmt"
	"os"

	"GMOps/src/service/bi/model"
	"GMOps/src/util"
)

var GMOPS_HOST_AUTHORIZED_KEYS string = os.Getenv("HOME") + "/.ssh/host_authorized_keys"

func (lgc *Logic) GetUserAuthKeys(user *model.User) ([]*model.UserAuthKey, error) {
	return model.GetUserAuthKey(user, lgc.db, lgc.logger)
}

func (lgc *Logic) AddUserAuthKey(user *model.User, title, key string) (string, error) {
	existed, err := model.ExistUserAuthKey(key, lgc.db, lgc.logger)
	if err != nil {
		return "", err
	}
	if existed {
		return "", nil
	}
	authKey := model.NewUserAuthKey(user, title, key)
	err = authKey.Save(lgc.db, lgc.logger)
	return authKey.Id.Hex(), err
}

func (lgc *Logic) GetUserAuthKeyByID(id string) (*model.UserAuthKey, error) {
	return model.LoadUserAuthKey(id, lgc.db, lgc.logger)
}

func (lgc *Logic) ApplyUserAuthKey(id string) error {
	_, err := os.Stat(GMOPS_HOST_AUTHORIZED_KEYS)
	if err != nil {
		lgc.logger.Error("BI Service ApplyUserAuthKey failed, os.Stat error: ", err)
		return err
	}

	authKey, err := model.LoadUserAuthKey(id, lgc.db, lgc.logger)
	if err != nil {
		lgc.logger.Error("BI Service ApplyUserAuthKey failed, cannot LoadUserAuthKey: ", err)
		return err
	}
	if !authKey.Writed {
		addr, err := util.GetNetworkInterface()
		if err != nil {
			lgc.logger.Error("BI Service ApplyUserAuthKey failed, GetNetworkInterface failed: ", err)
			return err
		}
		user, err := authKey.GetUser(lgc.db, lgc.logger)
		if err != nil {
			lgc.logger.Error("BI Service ApplyUserAuthKey failed, authKey.GetUser failed: ", err)
			return err
		}

		f, err := os.OpenFile(GMOPS_HOST_AUTHORIZED_KEYS, os.O_APPEND|os.O_WRONLY, 0o600)
		if err != nil {
			lgc.logger.Error("BI Service ApplyUserAuthKey failed, os.OpenFile failed: ", err)
			return err
		}
		_, err = fmt.Fprintf(f, "command=\"GWOPS_HOST=%s GWOPS_PORT=9022 /opt/GMOps/bin/gmops-proxy '%s' '%s'\" %s\n",
			addr, user.Username, authKey.Id.Hex(), authKey.AuthKey)
		if err != nil {
			f.Close()
			lgc.logger.Error("BI Service ApplyUserAuthKey failed, fmt.Fprintf failed: ", err)
			return err
		}
		f.Close()
		authKey.Writed = true
	}
	authKey.Available = true
	if err = authKey.Update(lgc.db, lgc.logger); err != nil {
		return err
	}

	return nil
}

func (lgc *Logic) CancelUserAuthKey(id string) error {
	authKey, err := model.LoadUserAuthKey(id, lgc.db, lgc.logger)
	if err != nil {
		lgc.logger.Error("BI Service CancelUserAuthKey failed, cannot LoadUserAuthKey: ", err)
		return err
	}
	authKey.Available = false
	if err = authKey.Update(lgc.db, lgc.logger); err != nil {
		return err
	}

	return nil
}
