package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
}
type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.length * s.length
}
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type shape interface {
	area() float64
}

func info(sh shape) {
	fmt.Println(sh.area())
}
func main() {
	sq := square{
		length: 25,
	}
	c := circle{
		radius: (25 / math.Sqrt(math.Pi)),
	}
	info(c)
	info(sq)
}
