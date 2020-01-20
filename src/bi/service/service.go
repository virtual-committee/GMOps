package service

import (
	"context"
	"os"
	"time"

	"GMOps/src/bi/logic"
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
	mongoClient          *mongo.Client
	lgc                  *logic.Logic
	ctx                  context.Context
	cancel               context.CancelFunc
	logger               *log.Logger
}

func (s *Service) initLogger() {
	logger := log.New()
	logger.Out = os.Stdout

	logger.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})

	s.logger = logger
}

func NewService(ctx context.Context, cancel context.CancelFunc, unixSocketPath string, mongoConnector string, dbName string) (*Service, error) {
	s := &Service{
		r:              gin.Default(),
		unixSocketPath: unixSocketPath,
		dbName:         dbName,
		ctx:            ctx,
		cancel:         cancel,
	}

	s.initLogger()

	mongoClient, err := mongo.NewClient(options.Client().ApplyURI(mongoConnector))
	if err != nil {
		s.logger.Error("BI Service cannot connected Mongo, failed mongo.NewClient. connector: ", mongoConnector)
		s.cancel()
		return nil, err
	}
	mongoTimeoutCtx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = mongoClient.Connect(mongoTimeoutCtx)
	if err != nil {
		s.logger.Error("BI Service cannot connected Mongo, connect timeout.")
		s.cancel()
		return nil, err
	}
	s.mongoClient = mongoClient
	s.logger.Info("Connected Mongo, connector: ", mongoConnector)

	s.modelIndexesCreators = make(map[string]func(db *mongo.Database, logger *log.Logger) error)
	if err = model.RegisterModelIndexes(&s.modelIndexesCreators); err != nil {
		s.logger.Error("BI Service cannot register model indexes creators")
		return nil, err
	}

	s.initRoute()

	s.lgc = logic.NewLogic(s.mongoClient.Database(s.dbName), s.logger)

	return s, nil
}

func (s *Service) InitDB() error {
	for _, f := range s.modelIndexesCreators {
		if err := f(s.mongoClient.Database(s.dbName), s.logger); err != nil {
			return err
		}
	}
	return nil
}

func (s *Service) Run() error {
	go func() {
		s.logger.Info("BI Service listening unix://", s.unixSocketPath)
		s.r.RunUnix(s.unixSocketPath)
	}()
	return nil
}
