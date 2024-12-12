package main

import (
	"fmt"
	"math"
)

func main() {
	p1 := NewPoint(1, 2)
	p2 := NewPoint(4, 5)

	coords := NewCoords(p1, p2)

	distance := coords.GetDistance()

	fmt.Println(distance)
}

type Point struct {
	x int
	y int
}

func (p *Point) getX() int {
	return p.x
}
func (p *Point) getY() int {
	return p.y
}

func NewPoint(x, y int) Point {
	return Point{x: x, y: y}
}

type Coords struct {
	point1 Point
	point2 Point
}

func NewCoords(p1, p2 Point) *Coords {
	return &Coords{point1: p1, point2: p2}
}

func (c *Coords) GetDistance() float64 {

	xPoints := math.Pow(float64(c.point2.getX()-c.point1.getX()), 2)
	yPoints := math.Pow(float64(c.point2.getY()-c.point1.getY()), 2)

	return math.Sqrt(xPoints + yPoints)
}
