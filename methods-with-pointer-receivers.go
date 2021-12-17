package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.Y = v.Y * f
	v.X = v.X * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func main() {
	v := &Vertex{8, 9}
	fmt.Printf("Before scaling:%+v,Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling:%+v,Abs:%v\n", v, v.Abs())
}
