package model

import (
	"context"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

type User struct {
	Id        string `bson:"_id"`
	Username  string
	Password  string
	Available bool
}

func createUserIndex(db *mongo.Database, logger *log.Logger) error {
	idx := mongo.IndexModel{
		Keys:    bsonx.Doc{{"username", bsonx.Int32(1)}},
		Options: options.Index().SetUnique(true),
	}
	ret, err := db.Collection(GMOPS_COLLECTION_USER).Indexes().CreateOne(context.Background(), idx)
	if err != nil {
		logger.Error("BI Server failed create User index: ", err)
		return err
	}
	logger.Info("BI Server success created User index: ", ret)
	return nil
}
