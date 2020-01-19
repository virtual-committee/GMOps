package service

import (
	"context"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Service struct {
	r              *gin.Engine
	unixSocketPath string

	MongoClient *mongo.Client
	Ctx         context.Context
	Cancel      context.CancelFunc
	Logger      *log.Logger
}

func (s *Service) initLogger() {
	logger := log.New()
	logger.Out = os.Stdout

	logger.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})

	s.Logger = logger
}

func NewService(ctx context.Context, cancel context.CancelFunc, unixSocketPath string, mongoConnector string) (*Service, error) {
	s := &Service{
		r:              gin.Default(),
		unixSocketPath: unixSocketPath,
		Ctx:            ctx,
		Cancel:         cancel,
	}

	s.initLogger()

	mongoClient, err := mongo.NewClient(options.Client().ApplyURI(mongoConnector))
	if err != nil {
		s.Logger.Error("BI Service cannot connected Mongo, failed mongo.NewClient. connector: ", mongoConnector)
		s.Cancel()
		return nil, err
	}
	mongoTimeoutCtx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = mongoClient.Connect(mongoTimeoutCtx)
	if err != nil {
		s.Logger.Error("BI Service cannot connected Mongo, connect timeout.")
		s.Cancel()
		return nil, err
	}
	s.MongoClient = mongoClient
	s.Logger.Info("Connected Mongo, connector: ", mongoConnector)

	return s, nil
}

func (s *Service) Run() error {
	go func() {
		s.Logger.Info("BI Service listening unix://", s.unixSocketPath)
		s.r.RunUnix(s.unixSocketPath)
	}()
	return nil
}
