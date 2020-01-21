package model

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

type UserAuthKey struct {
	Id        primitive.ObjectID `bson:"_id"`
	User      primitive.ObjectID
	Title     string
	AuthKey   string
	Writed    bool
	Available bool
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

func NewUserAuthKey(user *User, title, key string) *UserAuthKey {
	return &UserAuthKey{
		Id:        primitive.NewObjectID(),
		User:      user.Id,
		Title:     title,
		AuthKey:   key,
		Writed:    false,
		Available: false,
	}
}

func LoadUserAuthKey(id string, db *mongo.Database, logger *log.Logger) (*UserAuthKey, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		logger.Error("BI Server failed transfer Hex to ObjectID: ", err)
		return nil, err
	}
	ret := &UserAuthKey{}
	err = db.Collection(GMOPS_COLLECTION_USER_AUTH_KEY).FindOne(context.TODO(), bson.D{{"_id", oid}}).Decode(ret)
	if err != nil {
		logger.Error("BI Server failed LoadUserAuthKey: ", err)
		return nil, err
	}
	return ret, nil
}

func GetUserAuthKey(user *User, db *mongo.Database, logger *log.Logger) ([]*UserAuthKey, error) {
	ret := make([]*UserAuthKey, 0)
	cursor, err := db.Collection(GMOPS_COLLECTION_USER_AUTH_KEY).Find(context.TODO(), bson.D{{"user", user.Id}})
	defer cursor.Close(context.TODO())
	if err != nil {
		logger.Error("BI Server GetUserAuthKey failed find UserAuthKey: ", err)
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		userAuthKey := &UserAuthKey{}
		if err = cursor.Decode(&userAuthKey); err != nil {
			logger.Error("BI Server GetUserAuthKey failed decode: ", err)
			return nil, err
		}
		ret = append(ret, userAuthKey)
	}

	return ret, nil
}

func ExistUserAuthKey(key string, db *mongo.Database, logger *log.Logger) (bool, error) {
	count, err := db.Collection(GMOPS_COLLECTION_USER_AUTH_KEY).CountDocuments(context.TODO(), bson.D{{"authKey", key}})
	if err != nil {
		logger.Error("BI Server cannot got UserAuthKey count documents: ", err)
		return false, err
	}
	return count == 1, nil
}

func (k *UserAuthKey) Save(db *mongo.Database, logger *log.Logger) error {
	ret, err := db.Collection(GMOPS_COLLECTION_USER_AUTH_KEY).InsertOne(context.TODO(), k)
	if err != nil {
		logger.Error("BI Server UserAuthKey cannot save: ", err)
		return err
	}
	logger.Info("BI Server user inserted: ", ret)
	return nil
}

func (k *UserAuthKey) Update(db *mongo.Database, logger *log.Logger) error {
	update := bson.D{
		{"$set", bson.D{
			{"title", k.Title},
			{"writed", k.Writed},
			{"available", k.Available},
		}},
	}
	result, err := db.Collection(GMOPS_COLLECTION_USER_AUTH_KEY).UpdateOne(context.TODO(), bson.D{{"_id", k.Id}}, update)
	if err != nil {
		logger.Error("BI Server User Update failed: ", err)
		return err
	}
	if result.MatchedCount != 1 {
		logger.Error("BI Server User Update failed: updated count ", result)
		return fmt.Errorf("BI Server User Update failed: updated count ", result)
	}
	return nil
}

func (k *UserAuthKey) GetUser(db *mongo.Database, logger *log.Logger) (*User, error) {
	return LoadUserById(k.User, db, logger)
}
