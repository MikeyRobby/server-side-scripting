//Using the STRUCT “person”, using a composite literal create a value of type person and assign it to a variable with the identifier “p1”;
// print out “p1”;
// print out just the field fName for “p1”
package main

import "fmt"

type person struct{
fName string
lName string
}

func main() {
  p1 := person{fName:"Mikey", lName:"Robby"}
  fmt.Println(p1)
  fmt.Print(p1.fName)

}


