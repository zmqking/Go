package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.Y = v.X * f
	v.X = v.Y * f
}

func main() {
	v := Vertex{3, 5}
	v.Scale(10)
	fmt.Println(v.Abs())
}
