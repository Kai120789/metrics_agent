package utils

import (
	"agent/internal/dto"
	"fmt"
)

func PrintMetrics(metrics [31]dto.Metric) {
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
	fmt.Println()
}
