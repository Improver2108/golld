package main

import (
	"fmt"

	"github.com/improver2108/golld/solid/isp/shape/cube"
	"github.com/improver2108/golld/solid/isp/shape/rectangle"
	"github.com/improver2108/golld/solid/isp/shape/square"
)

// // better to break in two shape types. because Not all shapes have Volume and also some shapes will force to implement these functions
// type Shape interface {
// 	Area() float64
// 	Volume() float64
// }

type TwoDimensionalShape interface {
	Area() float64
}

type ThreeDimensionalShape interface {
	Area() float64
	Volume() float64
}

func main() {
	sqr := square.Init(5)
	rect := rectangle.Init(5, 4)
	cub := cube.Init(5)

	fmt.Println("Square Area: ", sqr.Area())
	fmt.Println("Rectangle Area: ", rect.Area())
	fmt.Println("Cube Area: ", cub.Area())
	fmt.Println("Cube Volume: ", cub.Volume())
}
