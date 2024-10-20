package app

import (
	"agent/internal/config"
	"agent/internal/dto"
	"agent/internal/utils"
	"agent/pkg/logger"
	"fmt"
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

	metrics := utils.GetMetrics()

	pollCountStartValue := int64(0)

	pollCount := dto.Metric{
		ID:        31,
		Name:      "PollCount",
		Type:      "counter",
		Value:     nil,
		Delta:     &pollCountStartValue,
		CreatedAt: time.Now(),
	}

	metrics = append(metrics, pollCount)

	_ = log

	for _, metric := range metrics {
		if metric.Value != nil {
			fmt.Println(
				"ID:", metric.ID,
				"Name:", metric.Name,
				"Type:", metric.Type,
				"Value:", *metric.Value,
				"CreatedAt:", metric.CreatedAt.String(),
			)
		} else {
			fmt.Println(
				"ID:", metric.ID,
				"Name:", metric.Name,
				"Type:", metric.Type,
				"Delta:", *metric.Delta,
				"CreatedAt:", metric.CreatedAt.String(),
			)
		}
	}

	// init service

	// init handler

	// init router

	//start server
}
