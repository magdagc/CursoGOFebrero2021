package main

import (
	"fmt"
	"math"
)

type vertice struct {
	X, Y float64
}

func (v vertice) abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func (v *vertice) agrandar(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
func main() {

	v := vertice{3, 4}

	fmt.Println(v.abs())

	v.agrandar(10)

	fmt.Println(v.abs())

}
