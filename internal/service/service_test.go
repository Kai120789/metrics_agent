package service_test

import (
	"agent/internal/api"
	"agent/internal/config"
	"agent/internal/dto"
	"agent/internal/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

type MockApi struct {
	mock.Mock
}

func (m *MockApi) SendMetrics(metrics [31]dto.Metric, serverURL, secretKey string) error {
	args := m.Called(metrics, serverURL, secretKey)
	return args.Error(0)
}

func TestAddPollCount(t *testing.T) {
	logger := zap.NewNop() // No-op logger
	cfg := &config.Config{}
	api := &api.Api{}
	srv := service.New(logger, cfg, api)

	// Arrange
	var metrics [31]dto.Metric

	// Act
	result := srv.AddPollCount(metrics)

	// Assert
	assert.Equal(t, "PollCount", result[0].Name)
	assert.Equal(t, "counter", result[0].Type)
	assert.NotNil(t, result[0].Delta)
	assert.Equal(t, int64(0), *result[0].Delta)
}

func TestCollectMetrics(t *testing.T) {
	logger := zap.NewNop()
	cfg := &config.Config{}
	api := &api.Api{}
	srv := service.New(logger, cfg, api)

	// Arrange
	var metrics [31]dto.Metric
	metrics = srv.AddPollCount(metrics)

	// Act
	collectedMetrics := srv.CollectMetrics(metrics)

	// Assert
	require.Len(t, collectedMetrics, 31)
	assert.Equal(t, "PollCount", collectedMetrics[0].Name)
	assert.Equal(t, "counter", collectedMetrics[0].Type)
	assert.NotNil(t, collectedMetrics[0].Delta)
	assert.Equal(t, int64(1), *collectedMetrics[0].Delta) // Delta должна быть увеличена на 1

	// Проверка метрик памяти
	for _, metric := range collectedMetrics[1:] {
		assert.NotEmpty(t, metric.Name)
		assert.Equal(t, "gauge", metric.Type)
		assert.NotNil(t, metric.Value)
	}
}

func TestSendMetrics(t *testing.T) {

}
