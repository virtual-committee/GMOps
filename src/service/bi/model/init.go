package model

import (
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterModelIndexes(f *map[string]func(db *mongo.Database, logger *log.Logger) error) error {
	(*f)["User"] = createUserIndex
	(*f)["UserAuthKey"] = createUserAuthKeyIndex
	(*f)["UserRepo"] = createUserRepoIndex

	return nil
}
