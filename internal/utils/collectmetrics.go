package utils

import (
	"agent/internal/dto"
	"math/rand/v2"
	"runtime"
	"time"
)

func GetMetrics() *[]dto.Metric {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	var metrics []dto.Metric

	// Helper function to add a metric
	addMetric := func(id uint, name string, value float64) {
		metric := dto.Metric{
			ID:        id,
			Name:      name,
			Type:      "gauge",
			Value:     &value,
			Delta:     nil, // No delta for gauge
			CreatedAt: time.Now(),
		}
		metrics = append(metrics, metric)
	}

	addMetric(1, "Alloc", float64(memStats.Alloc))
	addMetric(2, "BuckHashSys", float64(memStats.BuckHashSys))
	addMetric(3, "Frees", float64(memStats.Frees))
	addMetric(4, "GCCPUFraction", memStats.GCCPUFraction)
	addMetric(5, "GCSys", float64(memStats.GCSys))
	addMetric(6, "HeapAlloc", float64(memStats.HeapAlloc))
	addMetric(7, "HeapIdle", float64(memStats.HeapIdle))
	addMetric(8, "HeapInuse", float64(memStats.HeapInuse))
	addMetric(9, "HeapObjects", float64(memStats.HeapObjects))
	addMetric(10, "HeapReleased", float64(memStats.HeapReleased))
	addMetric(11, "HeapSys", float64(memStats.HeapSys))
	addMetric(12, "LastGC", float64(memStats.LastGC))
	addMetric(13, "Lookups", float64(memStats.Lookups))
	addMetric(14, "MCacheInuse", float64(memStats.MCacheInuse))
	addMetric(15, "MCacheSys", float64(memStats.MCacheSys))
	addMetric(16, "MSpanInuse", float64(memStats.MSpanInuse))
	addMetric(17, "MSpanSys", float64(memStats.MSpanSys))
	addMetric(18, "Mallocs", float64(memStats.Mallocs))
	addMetric(19, "NextGC", float64(memStats.NextGC))
	addMetric(20, "NumForcedGC", float64(memStats.NumForcedGC))
	addMetric(21, "NumGC", float64(memStats.NumGC))
	addMetric(22, "OtherSys", float64(memStats.OtherSys))
	addMetric(23, "PauseTotalNs", float64(memStats.PauseTotalNs))
	addMetric(24, "StackInuse", float64(memStats.StackInuse))
	addMetric(25, "StackSys", float64(memStats.StackSys))
	addMetric(26, "Sys", float64(memStats.Sys))
	addMetric(27, "TotalAlloc", float64(memStats.TotalAlloc))
	addMetric(28, "RandomValue", rand.Float64())
	addMetric(29, "TotalMemory", float64(memStats.Sys))
	addMetric(30, "FreeMemory", float64(memStats.Frees))

	return &metrics
}
