package main

import "fmt"

func main() {
  // Passing in arguments
  person("MikeyRobby", 22)

}

// Parameters are set in this function so that arguments are taken when called
// instead of stored in variables
func person(name string, age int) {
  fmt.Println(name, age)
}

// Parameters are set when creating a function
// Arguments are given when calling a function
