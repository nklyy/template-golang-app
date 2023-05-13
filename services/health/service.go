package health

import (
	"errors"

	"go.uber.org/zap"
)

//go:generate mockgen -source=service.go -destination=mocks/service_mock.go
type Service interface {
}

type service struct {
	repository Repository
	logger     *zap.SugaredLogger
	salt       int
}

func NewService(repository Repository, logger *zap.SugaredLogger, salt *int) (Service, error) {
	if repository == nil {
		return nil, errors.New("[health_service] invalid repository")
	}
	if logger == nil {
		return nil, errors.New("[health_service] invalid logger")
	}
	if salt == nil {
		return nil, errors.New("[health_service] invalid salt")
	}

	return &service{repository: repository, logger: logger, salt: *salt}, nil
}
