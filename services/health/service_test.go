package health_test

import (
	"template-golang-app/services/health"
	mock_health "template-golang-app/services/health/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"testing"
)

func TestNewService(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	salt := 10

	tests := []struct {
		name       string
		repository health.Repository
		logger     *zap.SugaredLogger
		salt       *int
		expect     func(*testing.T, health.Service, error)
	}{
		{
			name:       "should return service",
			repository: mock_health.NewMockRepository(controller),
			logger:     &zap.SugaredLogger{},
			salt:       &salt,
			expect: func(t *testing.T, s health.Service, err error) {
				assert.NotNil(t, s)
				assert.Nil(t, err)
			},
		},
		{
			name:       "should return invalid repository",
			repository: nil,
			logger:     &zap.SugaredLogger{},
			salt:       &salt,
			expect: func(t *testing.T, s health.Service, err error) {
				assert.Nil(t, s)
				assert.NotNil(t, err)
				assert.EqualError(t, err, "[health_service] invalid repository")
			},
		},
		{
			name:       "should return invalid logger",
			repository: mock_health.NewMockRepository(controller),
			logger:     nil,
			salt:       &salt,
			expect: func(t *testing.T, s health.Service, err error) {
				assert.Nil(t, s)
				assert.NotNil(t, err)
				assert.EqualError(t, err, "[health_service] invalid logger")
			},
		},
		{
			name:       "should return invalid salt",
			repository: mock_health.NewMockRepository(controller),
			logger:     &zap.SugaredLogger{},
			salt:       nil,
			expect: func(t *testing.T, s health.Service, err error) {
				assert.Nil(t, s)
				assert.NotNil(t, err)
				assert.EqualError(t, err, "[health_service] invalid salt")
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			svc, err := health.NewService(tc.repository, tc.logger, tc.salt)
			tc.expect(t, svc, err)
		})
	}
}
