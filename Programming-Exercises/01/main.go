//Initialize a SLICE of int using a composite literal;
// print out the slice;
// range over the slice printing out just the index;
// range over the slice printing out both the index and the value


package main

import (
  "fmt"
)

func main() {
  s := []int{10,15,20,25,30,35,40,45,50,55,60,65,70,75,80,85,90,95,100}
  fmt.Println(s)

  for i, _ := range s{
    fmt.Print(i)
  }

  for i, v := range s {
    fmt.Println(i,v)
  }
}

