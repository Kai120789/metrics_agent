package app

import (
	"agent/internal/config"
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

	_ = log
	// init service

	// init handler

	// init router

	//start server
}
