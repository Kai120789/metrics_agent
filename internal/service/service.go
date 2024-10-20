package service

import (
	"agent/internal/dto"
	"agent/internal/utils"
	"sync"
	"time"

	"go.uber.org/zap"
)

type Service struct {
	logger *zap.Logger
}

func New(l *zap.Logger) *Service {
	return &Service{
		logger: l,
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

			time.Sleep(time.Second * 2)
		}
	}()

	wg.Wait()
}
