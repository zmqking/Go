package main

import "fmt"

type vertex struct {
	Lat, Long float64
}

var m = map[string]vertex{
	"Bell Labs": {
		40.58741, -74.23654,
	},
	"Google": {
		23.87654, -58.25641,
	},
}

func main() {
	fmt.Println(m)
}
