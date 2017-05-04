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

  j := `[{"FirstName":"Mikey","LastName":"Robby","UserName":"Elite H@xor","Age":23,"PremiumMember":false},{"FirstName":"Bennie","LastName":"Jones","UserName":"Bonez","Age":27,"PremiumMember":true}]`

  fmt.Println("JSON: ", j)

  xpl := []player{}

  err := json.Unmarshal([]byte(j), &xpl)
  if err != nil {
    fmt.Println(err)
  }

  fmt.Printf("Go: %+v \n", xpl)

}
