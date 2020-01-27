package model

import (
	"context"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type GitHook struct {
	Id        primitive.ObjectID `bson:"_id"`
	Type      string
	Name      string
	LuaSource string
}

func NewGitHook() *GitHook {
	return &GitHook{Id: primitive.NewObjectID()}
}

func LoadGitHook(id primitive.ObjectID, db *mongo.Database, logger *log.Logger) (*GitHook, error) {
	ret := &GitHook{}
	if err := db.Collection(GMOPS_COLLECTION_GIT_HOOK).FindOne(context.TODO(), bson.D{{"_id", id}}).Decode(ret); err != nil {
		logger.Error("BI Server failed LoadGitHook: ", err)
		return nil, err
	}
	return ret, nil
}

func (g *GitHook) Save(db *mongo.Database, logger *log.Logger) error {
	ret, err := db.Collection(GMOPS_COLLECTION_GIT_HOOK).InsertOne(context.TODO(), g)
	if err != nil {
		logger.Error("BI Server failed Insert GitHook: ", err)
		return err
	}
	logger.Info("BI Server GitHook inserted: ", ret)
	return nil
}
