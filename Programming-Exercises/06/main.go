// Original Code      https://play.golang.org/p/r6gWlC_MzH
//Continuing with the code from here, add a method to type “person” with the identifier “walk”.
// Have the func print this string: <person’s first name> is walking.
// Call the method.
package main

import (
  "fmt"
)

type person struct {
  last string
  age  int
}

func main() {
  p1 := person{"Moneypenny", 24}
  fmt.Println(p1)
  p1.walk()
}

func (p person) walk() {
  fmt.Println(p.last + " is walking")
}

