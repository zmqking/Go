package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 将和送入 C
}

func main() {
	s := []int{7, 2, 6, 7, 0, -4}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c //从c中接收

	fmt.Println(x, y, x+y)
}
