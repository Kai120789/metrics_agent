package app

import (
	"agent/internal/api"
	"agent/internal/config"
	"agent/internal/dto"
	"agent/internal/service"
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
	zapLog, err := logger.New(cfg.LogLevel)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	log := zapLog.ZapLogger

	// init api
	api := api.New()

	// init service
	serv := service.New(log, cfg, api)

	var metrics [31]dto.Metric

	metrics = serv.AddPollCount(metrics)

	// create channels for polling and reporting
	collectTicker := time.NewTicker(time.Second * time.Duration(cfg.PollInterval))
	sendTicker := time.NewTicker(time.Second * time.Duration(cfg.ReportInterval))

	// loop to collect and send metrics
	for {
		select {
		// collect metrics every PollInterval seconds
		case <-collectTicker.C:
			metrics = serv.CollectMetrics(metrics)
		// send metrics every ReportInterval seconds
		case <-sendTicker.C:
			serv.SendMetrics(metrics)
		}
	}
}
