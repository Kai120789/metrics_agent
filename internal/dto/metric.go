package dto

import "time"

type Metric struct {
	ID        uint      // Уникальный идентификатор метрики
	Name      string    // Название метрики
	Type      string    // Тип метрики (counter или gauge)
	Value     *float64  // Значение для метрик типа gauge
	Delta     *int64    // Изменение для метрик типа counter
	CreatedAt time.Time // Время создания метрики
}
