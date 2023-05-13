package health

import (
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

//go:generate mockgen -source=repository.go -destination=mocks/repository_mock.go
type Repository interface {
}

type repository struct {
	db     *mongo.Client
	dbName string
	logger *zap.SugaredLogger
}

func NewRepository(db *mongo.Client, dbName string, logger *zap.SugaredLogger) (Repository, error) {
	if db == nil {
		return nil, errors.New("[health_repository] invalid user database")
	}
	if dbName == "" {
		return nil, errors.New("[health_repository] invalid database name")
	}
	if logger == nil {
		return nil, errors.New("[health_repository] invalid logger")
	}

	return &repository{db: db, dbName: dbName, logger: logger}, nil
}
