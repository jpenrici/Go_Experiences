// geometric_shapes.go
// Reference: https://gobyexample.com/interfaces

package main

import (
	"fmt"
	"math"
	"strings"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	line := strings.Repeat("-", 80)

	fmt.Println(line)
	fmt.Println("Type:", identify(g))
	fmt.Println("Area:", g.area())
	fmt.Println("Perimeter: ", g.perim())
}

func identify(g geometry) string {
	var label string
	switch g.(type) {
	case rect:
		label = "Rectangle"
	case circle:
		label = "Circle"
	default:
		label = "Unknow"
	}
	return label
}

func main() {
	measure(rect{width: 3, height: 4})
	measure(circle{radius: 5})
}
