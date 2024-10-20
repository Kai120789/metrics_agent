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

func (s *Service) AddPollCount(metrics []dto.Metric) []dto.Metric {
	pollCountStartValue := int64(0)

	pollCount := dto.Metric{
		ID:        1,
		Name:      "PollCount",
		Type:      "counter",
		Value:     nil,
		Delta:     &pollCountStartValue,
		CreatedAt: time.Now(),
	}

	metrics = append(metrics, pollCount)

	return metrics
}

func (s *Service) CollectMetrics(metrics []dto.Metric) {
	metrics = metrics[:1]

	metrics = append(metrics, utils.GetMetrics()...)

	*metrics[0].Delta += 1

	utils.PrintMetrics(metrics)

	time.Sleep(time.Second * time.Duration(s.config.PollInterval))
}

func (s *Service) SendMetrics(metrics []dto.Metric) {
	err := api.SendMetrics(metrics, s.config.ServerAddress)
	if err != nil {
		return
	}
}
