package utils

import (
	"agent/internal/dto"
	"fmt"
)

func PrintMetrics(metrics [31]dto.Metric) {
	var id uint = 1
	for _, metric := range metrics {

		if metric.Value != nil {
			fmt.Println(
				"ID:", id,
				"Name:", metric.Name,
				"Type:", metric.Type,
				"Value:", *metric.Value,
			)
		} else {
			fmt.Println(
				"ID:", id,
				"Name:", metric.Name,
				"Type:", metric.Type,
				"Delta:", *metric.Delta,
			)
		}

		id += 1
	}
	fmt.Println()
}
