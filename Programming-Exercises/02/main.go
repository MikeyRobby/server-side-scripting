//Initialize a MAP using a composite literal where the key is a string and the value is an int;
// print out the map;
// range over the map printing out just the key;
// range over the map printing out both the key and the value

package main

import "fmt"


func main() {
  m := map[string]int{"Go " : 98, "Chemistry " : 72, "Firewalls " : 84, "Integrated Devices ": 93, "PreCalculus ": 70,}
  fmt.Println(m)

  for k,_ := range m {
    fmt.Print(k)
  }

  for k,v := range m {
    fmt.Println(k,v)
  }

}
