package main

import "fmt"

type shape interface {
	getName() string
	getArea() float64
}

type shape1 struct{}

type square struct {
	name       string
	sideLength float64
}
type triangle struct {
	name   string
	height float64
	base   float64
}

func main() {
	sq := square{
		name:       "square",
		sideLength: 11.1,
	}
	tr := triangle{
		name:   "triangle",
		height: 4.2,
		base:   4.9,
	}
	printArea(sq)
	printArea(tr)
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}
func (s square) getName() string {
	return s.name
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}
func (t triangle) getName() string {
	return t.name
}

func printArea(s shape) {
	fmt.Println("The area of the", s.getName(), "is", s.getArea())
}
