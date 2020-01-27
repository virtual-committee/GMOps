package model

import (
	"context"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RepoGitHook struct {
	Id        primitive.ObjectID `bson:"_id"`
	RepoId    primitive.ObjectID
	GitHookId primitive.ObjectID
}

func NewRepoGitHook() *RepoGitHook {
	return &RepoGitHook{Id: primitive.NewObjectID()}
}

func LoadGitHookById(id primitive.ObjectID, db *mongo.Database, logger *log.Logger) (*GitHook, error) {
	ret := &GitHook{}
	if err := db.Collection(GMOPS_COLLECTION_REPO_GIT_HOOK).FindOne(context.TODO(), bson.D{{"_id", id}}).Decode(ret); err != nil {
		logger.Error("BI Server load GitHook: ", err)
		return nil, err
	}
	return ret, nil
}

func LoadGitHooksByRepoId(repo *Repo, db *mongo.Database, logger *log.Logger) ([]*GitHook, error) {
	ret := make([]*GitHook, 0)
	cursor, err := db.Collection(GMOPS_COLLECTION_REPO_GIT_HOOK).Find(context.TODO(), bson.D{{"repoid", repo.Id}})
	defer cursor.Close(context.TODO())
	if err != nil {
		logger.Error("BI Server GetGitHooksByRepoId failed find RepoGitHook: ", err)
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		repoGitHook := &RepoGitHook{}
		if err = cursor.Decode(&repoGitHook); err != nil {
			logger.Error("BI Server GetGitHooksByRepoId failed decode: ", err)
			return nil, err
		}
		gitHook, err := LoadGitHook(repoGitHook.GitHookId, db, logger)
		if err != nil {
			logger.Error("BI Server GetGitHooksByRepoId failed LoadGitHook: ", err)
			return nil, err
		}
		ret = append(ret, gitHook)
	}
	return ret, nil
}

func (r *RepoGitHook) Save(db *mongo.Database, logger *log.Logger) error {
	ret, err := db.Collection(GMOPS_COLLECTION_REPO_GIT_HOOK).InsertOne(context.TODO(), r)
	if err != nil {
		logger.Error("BI Server RepoGitHook cannot save: ", err)
		return err
	}
	logger.Info("BI Server RepoGitHook inserted: ", ret)
	return nil
}
