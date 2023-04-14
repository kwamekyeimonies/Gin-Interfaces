package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	rectwidth, rectheight float64
}

type circle struct {
	radius float64
}

func (x rectangle) area() float64 {
	return x.rectheight * x.rectwidth
}

func (y circle) area() float64 {
	return math.Pi * y.radius
}

func main() {
	Rectangle := rectangle{rectwidth: 3.4, rectheight: 9.7}
	Circle := circle{radius: 7.8}

	fmt.Println(Rectangle)
	fmt.Println(Circle)

	calc_area := Rectangle.area()
	calc_area_circle := Circle.area()

	fmt.Println(calc_area)
	fmt.Println(calc_area_circle)
}
