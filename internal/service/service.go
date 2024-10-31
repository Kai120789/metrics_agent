package service

import (
	"agent/internal/api"
	"agent/internal/config"
	"agent/internal/dto"
	"agent/internal/utils"
	"math/rand/v2"
	"runtime"

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
		Name:  "PollCount",
		Type:  "counter",
		Value: nil,
		Delta: &pollCountStartValue,
	}

	metrics[0] = pollCount

	return metrics
}

func (s *Service) CollectMetrics(metrics [31]dto.Metric) [31]dto.Metric {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	retMetrics := [31]dto.Metric{
		metrics[0],
		addMetric("Alloc", float64(memStats.Alloc)),
		addMetric("BuckHashSys", float64(memStats.BuckHashSys)),
		addMetric("Frees", float64(memStats.Frees)),
		addMetric("GCCPUFraction", memStats.GCCPUFraction),
		addMetric("GCSys", float64(memStats.GCSys)),
		addMetric("HeapAlloc", float64(memStats.HeapAlloc)),
		addMetric("HeapIdle", float64(memStats.HeapIdle)),
		addMetric("HeapInuse", float64(memStats.HeapInuse)),
		addMetric("HeapObjects", float64(memStats.HeapObjects)),
		addMetric("HeapReleased", float64(memStats.HeapReleased)),
		addMetric("HeapSys", float64(memStats.HeapSys)),
		addMetric("LastGC", float64(memStats.LastGC)),
		addMetric("Lookups", float64(memStats.Lookups)),
		addMetric("MCacheInuse", float64(memStats.MCacheInuse)),
		addMetric("MCacheSys", float64(memStats.MCacheSys)),
		addMetric("MSpanInuse", float64(memStats.MSpanInuse)),
		addMetric("MSpanSys", float64(memStats.MSpanSys)),
		addMetric("Mallocs", float64(memStats.Mallocs)),
		addMetric("NextGC", float64(memStats.NextGC)),
		addMetric("NumForcedGC", float64(memStats.NumForcedGC)),
		addMetric("NumGC", float64(memStats.NumGC)),
		addMetric("OtherSys", float64(memStats.OtherSys)),
		addMetric("PauseTotalNs", float64(memStats.PauseTotalNs)),
		addMetric("StackInuse", float64(memStats.StackInuse)),
		addMetric("StackSys", float64(memStats.StackSys)),
		addMetric("Sys", float64(memStats.Sys)),
		addMetric("TotalAlloc", float64(memStats.TotalAlloc)),
		addMetric("RandomValue", rand.Float64()),
		addMetric("TotalMemory", float64(memStats.Sys)),
		addMetric("FreeMemory", float64(memStats.Frees)),
	}

	*retMetrics[0].Delta += 1

	utils.PrintMetrics(retMetrics)

	return retMetrics
}

func (s *Service) SendMetrics(metrics [31]dto.Metric) {
	err := api.SendMetrics(metrics, s.config.ServerURL, s.config.SecretKey)
	if err != nil {
		return
	}
}

func addMetric(name string, value float64) dto.Metric {
	metric := dto.Metric{
		Name:  name,
		Type:  "gauge",
		Value: &value,
		Delta: nil,
	}
	return metric
}
