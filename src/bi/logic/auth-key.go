package logic

import (
	"GMOps/src/bi/model"
)

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
