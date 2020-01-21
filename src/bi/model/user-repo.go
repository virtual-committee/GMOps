package model

import (
	"context"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

type UserRepo struct {
	Id   primitive.ObjectID `bson:"_id"`
	User primitive.ObjectID
	Repo primitive.ObjectID
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

func NewUserRepo() *UserRepo {
	return &UserRepo{Id: primitive.NewObjectID()}
}

func LoadUserReposByUser(user *User, db *mongo.Database, logger *log.Logger) ([]*UserRepo, error) {
	ret := make([]*UserRepo, 0)
	cursor, err := db.Collection(GMOPS_COLLECTION_USER_REPO).Find(context.TODO(), bson.D{{"_id", user.Id}})
	defer cursor.Close(context.TODO())
	if err != nil {
		logger.Error("BI Server LoadUserReposByUser failed find UserRepo: ", err)
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		userRepo := &UserRepo{}
		if err = cursor.Decode(&userRepo); err != nil {
			logger.Error("BI Server LoadUserReposByUser failed decode: ", err)
			return nil, err
		}
		ret = append(ret, userRepo)
	}

	return ret, nil
}

func (ur *UserRepo) Save(db *mongo.Database, logger *log.Logger) error {
	ret, err := db.Collection(GMOPS_COLLECTION_USER_REPO).InsertOne(context.TODO(), ur)
	if err != nil {
		logger.Error("BI Server UserRepo cannot save: ", err)
		return err
	}
	logger.Info("BI Server UserRepo inserted: ", ret)
	return nil
}
