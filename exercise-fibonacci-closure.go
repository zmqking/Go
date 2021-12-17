package main

import "fmt"

//返回一个“返回 int的函数”
func fibnoacci(i int) int {
	if i < 2 {
		return i
	} else {
		return fibnoacci(i-1) + fibnoacci(i-2)
	}
}

func fibonacci1() func() int {
	back1, back2 := 0, 1

	return func() int {
		temp := back1
		back1, back2 = back2, (back1 + back2)
		return temp
	}

}

func main() {
	// for i := 0; i < 10; i++ {
	f := fibonacci1()
	fmt.Println(f())
	// }
	for i := 0; i < 10; i++ {
		fmt.Println(fibnoacci(i))
	}
}
