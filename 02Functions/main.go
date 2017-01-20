package main

import "fmt"

func main() {
  // calling the person function that was created
	person() // the () are required when calling a function
}

//creating a function
func person() {
	name := "MikeyRobby"
	age := "22"
  //printing out the value of name and age
	fmt.Println(name, age)
}
