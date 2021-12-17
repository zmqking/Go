package main

import "math"

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X / v.Y)
}

func main() {
	v := Vertex{3, 8}
	println(Abs(v))
}
