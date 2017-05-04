package main

import (
  "fmt"
  "encoding/json"
)

type person struct {
  First string
  Last string
  TopSecret bool
  Items []string
}

func main() {
  p1 := person{
    First: "Mikey",
    Last: "Robby",
    TopSecret: true,
    Items: []string{
      "Badge",
      "Weapon",
      "USB Codes",
      "Water Bottle",
    },
  }
  bs, err := json.Marshal(p1)
  if err != nil{
    fmt.Println(err)
  }

  fmt.Println(p1)
  fmt.Println("-----------------------------------------------------------")
  fmt.Println(string(bs))

}
