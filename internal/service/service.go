package service

import (
	"agent/internal/api"
	"agent/internal/config"
	"agent/internal/dto"
	"agent/internal/utils"
	"time"

	"go.uber.org/zap"
)

type Service struct {
	logger *zap.Logger
	config *config.Config
}

func New(l *zap.Logger, c *config.Config) *Service {
	return &Service{
		logger: l,
		config: c,
	}
}

func (s *Service) AddPollCount(metrics [31]dto.Metric) [31]dto.Metric {
	pollCountStartValue := int64(0)

	pollCount := dto.Metric{
		ID:        1,
		Name:      "PollCount",
		Type:      "counter",
		Value:     nil,
		Delta:     &pollCountStartValue,
		CreatedAt: time.Now(),
	}

	metrics[0] = pollCount

	return metrics
}

func (s *Service) CollectMetrics(metrics [31]dto.Metric) [31]dto.Metric {
	metricsAll := utils.GetMetrics(metrics)

	*metricsAll[0].Delta += 1

	utils.PrintMetrics(metricsAll)

	return metricsAll
}

func (s *Service) SendMetrics(metrics [31]dto.Metric) {
	err := api.SendMetrics(metrics, s.config.ServerURL, s.config.SecretKey)
	if err != nil {
		return
	}
}
