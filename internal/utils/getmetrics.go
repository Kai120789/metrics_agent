package utils

import (
	"agent/internal/dto"
	"math/rand/v2"
	"runtime"
	"time"
)

func GetMetrics() []dto.Metric {
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

	addMetric(2, "Alloc", float64(memStats.Alloc))
	addMetric(3, "BuckHashSys", float64(memStats.BuckHashSys))
	addMetric(4, "Frees", float64(memStats.Frees))
	addMetric(5, "GCCPUFraction", memStats.GCCPUFraction)
	addMetric(6, "GCSys", float64(memStats.GCSys))
	addMetric(7, "HeapAlloc", float64(memStats.HeapAlloc))
	addMetric(8, "HeapIdle", float64(memStats.HeapIdle))
	addMetric(9, "HeapInuse", float64(memStats.HeapInuse))
	addMetric(10, "HeapObjects", float64(memStats.HeapObjects))
	addMetric(11, "HeapReleased", float64(memStats.HeapReleased))
	addMetric(12, "HeapSys", float64(memStats.HeapSys))
	addMetric(13, "LastGC", float64(memStats.LastGC))
	addMetric(14, "Lookups", float64(memStats.Lookups))
	addMetric(15, "MCacheInuse", float64(memStats.MCacheInuse))
	addMetric(16, "MCacheSys", float64(memStats.MCacheSys))
	addMetric(17, "MSpanInuse", float64(memStats.MSpanInuse))
	addMetric(18, "MSpanSys", float64(memStats.MSpanSys))
	addMetric(19, "Mallocs", float64(memStats.Mallocs))
	addMetric(20, "NextGC", float64(memStats.NextGC))
	addMetric(21, "NumForcedGC", float64(memStats.NumForcedGC))
	addMetric(22, "NumGC", float64(memStats.NumGC))
	addMetric(23, "OtherSys", float64(memStats.OtherSys))
	addMetric(24, "PauseTotalNs", float64(memStats.PauseTotalNs))
	addMetric(25, "StackInuse", float64(memStats.StackInuse))
	addMetric(26, "StackSys", float64(memStats.StackSys))
	addMetric(27, "Sys", float64(memStats.Sys))
	addMetric(28, "TotalAlloc", float64(memStats.TotalAlloc))
	addMetric(29, "RandomValue", rand.Float64())
	addMetric(30, "TotalMemory", float64(memStats.Sys))
	addMetric(31, "FreeMemory", float64(memStats.Frees))

	return metrics
}
