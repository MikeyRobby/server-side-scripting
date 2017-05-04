package main

// Go to localhost:8080 to see the JSON data
// localhost:8080/send send the data to our system through an ajax script
//localhost:8080/catch shows the data and sends it to the terminal

import (
  "net/http"
  "encoding/json"
  "fmt"
  "html/template"
)
// creates the struct that our JSON is based on
type player struct{
  FirstName string
  LastName string
  UserName string
  Age int
  PremiumMember bool
}
// Global variables
// the xpl variable needs to be here so that all functions can access it
var tpl *template.Template
var xpl []player
// initializing the templates folder
func init()  {
  tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
  // Routes
  http.HandleFunc("/", index)
  http.HandleFunc("/send", send)
  http.HandleFunc("/catch", catch)
  // default uses localhost, specifying port as :8080
  http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
  // Creating players
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

  // setting players to a xpl [slice of players]
  xpl = []player{p1,p2}
// JSON encoding
  err := json.NewEncoder(w).Encode(xpl)
  if err != nil {
    fmt.Println(err)
  }
  // No need to execute a template because we use our ResponseWriter "w" in the Encoder
}

func send(w http.ResponseWriter, r *http.Request) {
  // Executing index web page
  tpl.ExecuteTemplate(w, "index.html", nil)
}


func catch(w http.ResponseWriter, r *http.Request) {
// If the data is sent using POST decode the JSON using the body of the Request
  if r.Method == "POST"{

    err := json.NewDecoder(r.Body).Decode(&xpl)
    if err != nil {
      fmt.Println(err)
      return
    }
}
// Print data to the terminal
  fmt.Println(xpl)
// Executing the catch web page and passing in our xpl variable as data
  tpl.ExecuteTemplate(w, "catch.html", xpl)
}
