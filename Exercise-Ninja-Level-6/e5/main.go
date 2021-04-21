package main

import "fmt"

const pi float64 = 3.14

type square struct {
	length float64
	width  float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.length * s.width
}

func (c circle) area() float64 {
	return pi * c.radius * c.radius
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	circ := circle{
		radius: 12.345,
	}
	squa := square{
		length: 15,
		width:  3,
	}
	info(circ)
	info(squa)
}
