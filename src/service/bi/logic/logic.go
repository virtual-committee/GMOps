package logic

import (
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type Logic struct {
	db     *mongo.Database
	logger *log.Logger
}

func NewLogic(db *mongo.Database, logger *log.Logger) *Logic {
	return &Logic{
		logger: logger,
		db:     db,
	}
}
