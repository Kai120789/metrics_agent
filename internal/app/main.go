package app

import (
	"agent/internal/config"
	"agent/internal/service"
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
	zapLog, err := logger.New(cfg.LogLevel)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	log := zapLog.ZapLogger

	// init service
	serv := service.New(log)

	serv.CollectMetrics()

	//start server
}
