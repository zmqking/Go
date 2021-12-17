package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["Answer"] = 42
	fmt.Println("The Value:", m["Answer"])

	m["Answer"] = 98
	fmt.Println("New Value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The Value:", m["Answer"])

	v, ok := m["Answer"] //ok 代表是否存在
	fmt.Println("The Value:", v, "Prsent?", ok)

	_, e := m["Answer"]
	fmt.Println("Exists?", e)
}
