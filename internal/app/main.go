package app

import (
	"agent/internal/config"
	"agent/internal/utils"
	"agent/pkg/logger"
	"fmt"
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

	_ = log

	for _, metric := range *metrics {
		fmt.Println(
			"ID:", metric.ID,
			"Name:", metric.Name,
			"Type:", metric.Type,
			"Value:", *metric.Value,
			"CreatedAt:", metric.CreatedAt.String(),
		)
	}

	// init service

	// init handler

	// init router

	//start server
}
