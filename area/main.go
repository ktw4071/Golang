package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sidelength float64
}

func main() {
	t := triangle{10, 5}
	s := square{20}

	printArea(t)
	printArea(s)
}
func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sidelength * s.sidelength
}
