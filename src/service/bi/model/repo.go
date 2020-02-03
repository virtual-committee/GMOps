package model

import (
	"context"
	"fmt"

	git "github.com/libgit2/git2go"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repo struct {
	Id       primitive.ObjectID `bson:"_id"`
	Spec     bool
	Name     string
	Descript string
}

func NewRepo() *Repo {
	return &Repo{Id: primitive.NewObjectID()}
}

func LoadRepoById(id primitive.ObjectID, db *mongo.Database, logger *log.Logger) (*Repo, error) {
	ret := &Repo{}
	if err := db.Collection(GMOPS_COLLECTION_REPO).FindOne(context.TODO(), bson.D{{"_id", id}}).Decode(ret); err != nil {
		logger.Error("BI Server load Repo: ", err)
		return nil, err
	}
	return ret, nil
}

func (r *Repo) Save(db *mongo.Database, logger *log.Logger) error {
	ret, err := db.Collection(GMOPS_COLLECTION_REPO).InsertOne(context.TODO(), r)
	if err != nil {
		logger.Error("BI Server Repo cannot save: ", err)
		return err
	}
	logger.Info("BI Server Repo inserted: ", ret)
	return nil
}

func (r *Repo) OpenGitRepo(basePath string, logger *log.Logger) (*git.Repository, error) {
	realPath := fmt.Sprintf("%s%s", basePath, r.Id.Hex())
	gitRepo, err := git.OpenRepository(realPath)
	if err != nil {
		logger.Error("BI Server Repo open git repository error: ", err)
		return nil, err
	}
	return gitRepo, nil
}
