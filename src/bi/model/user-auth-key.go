package model

import (
	"context"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

type UserAuthKey struct {
	Id      string `bson:"_id"`
	User    string
	Title   string
	AuthKey string
	Writed  bool
}

func createUserAuthKeyIndex(db *mongo.Database, logger *log.Logger) error {
	idx := mongo.IndexModel{
		Keys:    bsonx.Doc{{"authKey", bsonx.Int32(1)}},
		Options: options.Index().SetUnique(true),
	}
	ret, err := db.Collection(GMOPS_COLLECTION_USER_AUTH_KEY).Indexes().CreateOne(context.Background(), idx)
	if err != nil {
		logger.Error("BI Server failed create UserAuthKey index: ", err)
		return err
	}
	logger.Info("BI Server success created UserAuthKey index: ", ret)
	return nil
}
