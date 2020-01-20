package model

import (
	"context"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

type UserRepo struct {
	Id   string `bson:"_id"`
	User string
	Repo string
}

func createUserRepoIndex(db *mongo.Database, logger *log.Logger) error {
	idx := mongo.IndexModel{
		Keys: bsonx.Doc{
			{"user", bsonx.Int32(1)},
			{"repo", bsonx.Int32(1)},
		},
		Options: options.Index().SetUnique(true),
	}
	ret, err := db.Collection(GMOPS_COLLECTION_USER_REPO).Indexes().CreateOne(context.Background(), idx)
	if err != nil {
		logger.Error("BI Server failed create UserRepo index: ", err)
		return err
	}
	logger.Info("BI Server success created UserRepo index: ", ret)
	return nil
}
