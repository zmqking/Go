package main

import (
	"fmt"
	"math"
)

type Verex struct {
	x, y float64
}

func (v Verex) Abs() float64 {
	return math.Sqrt(v.x * v.y)
}

func main() {
	v := Verex{3, 8}
	fmt.Println(v.Abs())
}
