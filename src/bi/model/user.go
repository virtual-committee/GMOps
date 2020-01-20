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

type User struct {
	Id        primitive.ObjectID `bson:"_id"`
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

func NewUser() *User {
	return &User{Id: primitive.NewObjectID()}
}

func LoadUserByUsername(username string, db *mongo.Database, logger *log.Logger) (*User, error) {
	ret := &User{}
	if err := db.Collection(GMOPS_COLLECTION_USER).FindOne(context.TODO(), bson.D{{"username", username}}).Decode(ret); err != nil {
		logger.Error("BI Server cannot load user: ", err)
		return nil, err
	}
	return ret, nil
}

func LoadUserById(id primitive.ObjectID, db *mongo.Database, logger *log.Logger) (*User, error) {
	ret := &User{}
	if err := db.Collection(GMOPS_COLLECTION_USER).FindOne(context.TODO(), bson.D{{"_id", id}}).Decode(ret); err != nil {
		logger.Error("BI Server cannot load user: ", err)
		return nil, err
	}

	return ret, nil
}

func ExistUser(username string, db *mongo.Database, logger *log.Logger) (bool, error) {
	count, err := db.Collection(GMOPS_COLLECTION_USER).CountDocuments(context.TODO(), bson.D{{"username", username}})
	if err != nil {
		logger.Error("BI Server cannot got User count documents: ", err)
		return false, err
	}
	return count == 1, nil
}

func (u *User) Save(db *mongo.Database, logger *log.Logger) error {
	ret, err := db.Collection(GMOPS_COLLECTION_USER).InsertOne(context.TODO(), u)
	if err != nil {
		logger.Error("BI Server User cannot save: ", err)
		return err
	}
	logger.Info("BI Server user inserted: ", ret)
	return nil
}

func (u *User) Update(db *mongo.Database, logger *log.Logger) error {
	update := bson.D{
		{"$set", bson.D{
			{"password", u.Password},
			{"available", u.Available},
		}},
	}
	if _, err := db.Collection(GMOPS_COLLECTION_USER).UpdateOne(context.TODO(), bson.D{{"_id", u.Id}}, update); err != nil {
		logger.Error("BI Server User Update failed: ", err)
		return err
	}
	return nil
}
