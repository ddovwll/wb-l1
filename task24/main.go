package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

func (p Point) Distance(q Point) float64 {
	dx := math.Abs(p.x - q.x)
	dy := math.Abs(p.y - q.y)

	return math.Hypot(dx, dy)
}

func main() {
	p1 := NewPoint(3, 4)
	p2 := NewPoint(0, 0)
	fmt.Println(p1.Distance(*p2))
}
