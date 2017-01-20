package main

import "fmt"

func main() {
// strings are expressed using quotes.
  fmt.Println(concatenate("Mikey"))
}

// This function takes in a string and places it into the sentence below.
func concatenate(word string) string {
  return "Hello, " + word + "!"
}

// For more info on how to use String see my "Go Fundamental Notes
