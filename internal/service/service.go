package service

import (
	"agent/internal/config"
	"agent/internal/dto"
	"agent/internal/utils"
	"sync"
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

func (s *Service) CollectMetrics() {
	var metrics []dto.Metric

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

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		for {
			metrics = metrics[:1]

			metrics = append(metrics, utils.GetMetrics()...)

			*metrics[0].Delta += 1

			utils.PrintMetrics(metrics)

			time.Sleep(time.Second * time.Duration(s.config.PollInterval))
		}
	}()

	wg.Wait()
}

func (s *Service) SendMetrics(metrics []dto.Metric) {

}
