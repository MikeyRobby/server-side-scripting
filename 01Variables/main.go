package main

import "fmt"

// Package Scope Variable
var a int // Any function can use this variable
// Can also be done this way
var b int = 22

func main() {
  a = 10 // Setting the package scope variable 'a' to a value of 10
  // Short Declaration
  c := a * b // Can only be used inside of the function when declared like this. Type is inferred automatically.

  fmt.Println(c) // Shows the value of 'c' on the screen

}
