package main

import (
	"fmt"
	"math"
)

// Структура Point с инкапсулированными параметрами x, y.
type Point struct {
	x, y float64
}

// Конструктор для создания новой точки.
func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

// Метод для расчета расстояния между двумя точками.
func (p *Point) DistanceTo(other *Point) float64 {
	return math.Sqrt(math.Pow(other.x-p.x, 2) + math.Pow(other.y-p.y, 2))
}

func main() {
	// Создание двух точек.
	pointA := NewPoint(1, 2)
	pointB := NewPoint(4, 6)

	// Вычисление расстояния между точками.
	distance := pointA.DistanceTo(pointB)

	fmt.Printf("Расстояние между точками A и B: %f\n", distance)
}

