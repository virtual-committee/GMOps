package service

import (
	"context"
	"os"
	"time"

	"GMOps/src/bi/model"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Service struct {
	r                    *gin.Engine
	unixSocketPath       string
	dbName               string
	modelIndexesCreators map[string]func(db *mongo.Database, logger *log.Logger) error

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

func NewService(ctx context.Context, cancel context.CancelFunc, unixSocketPath string, mongoConnector string, dbName string) (*Service, error) {
	s := &Service{
		r:              gin.Default(),
		unixSocketPath: unixSocketPath,
		dbName:         dbName,
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

	s.modelIndexesCreators = make(map[string]func(db *mongo.Database, logger *log.Logger) error)
	if err = model.RegisterModelIndexes(&s.modelIndexesCreators); err != nil {
		s.Logger.Error("BI Service cannot register model indexes creators")
		return nil, err
	}

	return s, nil
}

func (s *Service) InitDB() error {
	for _, f := range s.modelIndexesCreators {
		if err := f(s.MongoClient.Database(s.dbName), s.Logger); err != nil {
			return err
		}
	}
	return nil
}

func (s *Service) Run() error {
	go func() {
		s.Logger.Info("BI Service listening unix://", s.unixSocketPath)
		s.r.RunUnix(s.unixSocketPath)
	}()
	return nil
}
