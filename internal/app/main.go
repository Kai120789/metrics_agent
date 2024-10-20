package app

import (
	"agent/internal/config"
	"agent/internal/dto"
	"agent/internal/utils"
	"agent/pkg/logger"
	"fmt"
	"sync"
	"time"
)

func StartApp() {
	// init config
	cfg, err := config.GetConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// init logger
	log, err := logger.New(cfg.LogLevel)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

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

	_ = log

	wg.Wait()

	// init service

	// init handler

	// init router

	//start server
}
