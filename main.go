package main

import (
	"fmt"
)

func main() {
	// println("hello go!")
	// println(len("hello go!")) //9字节
	// println(len("你好"))        //6字节
	// var a uint = 1
	// var b uint = 2
	// fmt.Println(a - b)
	// var s []int = make([]int, 5)
	s := []int{1, 3, 5, 7, 8}

	fmt.Println(Add(s, 2, 9))
	fmt.Println(Delete(s, 2))
	var f = 3.1
	fmt.Printf("%.2f", f)

}

func Add(s []int, index int, value int) []int {
	var stemp []int = make([]int, 0, len(s)+1)
	for i := 0; i < len(s); i++ {
		if i == index {
			stemp = append(stemp, value)
		}
		stemp = append(stemp, s[i])
	}
	return stemp
}

func Delete(s []int, index int) []int {
	var stemp []int = make([]int, 0, len(s)-1)
	for i := 0; i < len(s); i++ {
		if i != index {
			stemp = append(stemp, s[i])
		}
	}
	return stemp
}
