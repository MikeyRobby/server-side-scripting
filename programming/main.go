//test
//edit
//edit2


package main

import (
	"fmt"
)

type person struct {
	first string
}

func main() {
	p1 := person{"Mikey"}
	p1.speak()

	xi := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	m := map[string]int{
		"Roberts:": 22,
		"Jones:":   23,
		"Spry:":    50,
		"Jenkins:": 68,
		"smith:":   17,
	}

	fmt.Println(xi)
	for i, v := range xi {
		fmt.Println(i, v)
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
}

func (p person) speak() {
	fmt.Println("Hello " + p.first)
}
