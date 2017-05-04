package main

import (
  "encoding/json"
  "fmt"
)
type player struct{
  FirstName string
  LastName string
  UserName string
  Age int
  PremiumMember bool
}

func main() {

  p1 := player{
    FirstName: "Mikey",
    LastName: "Robby",
    UserName: "Elite H@xor",
    Age: 23,
    PremiumMember: false,
  }
  p2 := player{
    FirstName: "Bennie",
    LastName: "Jones",
    UserName: "Bonez",
    Age: 27,
    PremiumMember: true,
  }

  xpl := []player{p1,p2}

  fmt.Printf("Go: %+v \n", xpl)

  bs, err := json.Marshal(xpl)
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println("JSON: ", string(bs))

}
