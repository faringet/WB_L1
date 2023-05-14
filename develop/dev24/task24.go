package main

import (
	"fmt"
	"math"
)

/*
Разработать программу нахождения расстояния между двумя точками,
которые представлены в виде структуры Point с инкапсулированными параметрами x, y и конструктором.
*/

/*
Формула вычисления расстояния между двумя точками - AB = √(Xb - Xa)^2 + (Yb - Ya)^2
*/

// Point структура
type Point struct {
	x float64
	y float64
}

// NewPoint - конструктор
func NewPoint(x float64, y float64) *Point {
	return &Point{x: x, y: y}
}

// Применяем формулу
func distPoints(p1, p2 *Point) float64 {
	ans := math.Sqrt(math.Pow(p2.x-p1.x, 2) + math.Pow(p2.y-p1.y, 2))
	return ans
}

func main() {
	p1 := NewPoint(10.5, 100)
	p2 := NewPoint(2, 25.5)

	fmt.Println(distPoints(p1, p2))
}
