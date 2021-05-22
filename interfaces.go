package main

import (
	"fmt"
	"math"
)

func main() {

	//old way
	var r geometry = rectangle{3, 4}
	fmt.Println("rect area", r.area())
	fmt.Println("rect perimeter", r.perimeter())

	var c geometry = circle{5}
	fmt.Println("circle area", c.area())
	fmt.Println("circle perimeter", c.perimeter())

}

//interfaces are collection of behaviours
//has type
type geometry interface {
	area() float64
	perimeter() float64
}

//structs
type circle struct {
	radious float64
}

type rectangle struct {
	height, width float64
}

//methods to implement for rectangle
func (r rectangle) area() float64 {
	return r.height * r.width
}

func (r rectangle) perimeter() float64 {
	return 2*r.height + 2*r.width
}

//methods to implement for circle
func (c circle) area() float64 {
	return math.Pi * c.radious * c.radious
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radious
}
