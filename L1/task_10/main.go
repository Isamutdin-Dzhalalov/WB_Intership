package main

import (
	"fmt"
	"math"
)

func groupTemperatures(temps []float64) map[int][]float64 {
	groups := make(map[int][]float64)
	for _, temp := range temps {
		// Вычисляем ключ.
		// Округляем текущее значение температуры вверх до ближайшего числа, кратного 10.
		key := int(math.Ceil(temp / 10.0)) * 10
		// Добавляем значение температуры по ключу.
		groups[key] = append(groups[key], temp)
	}
	return groups
}

func main() {
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	groups := groupTemperatures(temps)
	for key, values := range groups {
		fmt.Printf("%d: %v\n", key, values)
	}
}

