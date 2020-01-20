package logic

import (
	"GMOps/src/bi/model"
)

func (lgc *Logic) ExistsUser(name string) (bool, error) {
	return model.ExistUser(name, lgc.db, lgc.logger)
}

func (lgc *Logic) CreateUser(name string, password string) (string, error) {
	user := model.NewUser()
	user.Username = name
	user.Password = password
	user.Available = true

	err := user.Save(lgc.db, lgc.logger)
	if err != nil {
		return "", err
	}
	return user.Id.Hex(), nil
}
