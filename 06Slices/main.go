package main

import "fmt"

func main() {

  // setting the variable 'n' equal to a Slice of Integers
  n := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

  // For more information on For Loops look in the "Go Fundamental Notes"
  // For Loop using range
  for i,v := range n {
    fmt.Println(i, v)
    // prints the index and the value

  }
// For Loop
  for i := 0; i < len(n); i++ {
    fmt.Println(n[i])
  }

  // The For Loops are printing out the values of the slice 'n' at the index i
  // index position is denoted using []
  // The i in the for loop is used for the index position
  // 'i' increases by 1 every time the loop runs until it reaches the length of 'n'
}
