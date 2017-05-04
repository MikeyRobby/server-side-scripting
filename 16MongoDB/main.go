package main

import (
  "html/template"
  "net/http"
  "fmt"
  "gopkg.in/mgo.v2"
)

var tpl *template.Template
var DB *mgo.Database
var books *mgo.Collection


func init()  {
  tpl = template.Must(template.ParseGlob("templates/*"))
  s, err := mgo.Dial("mongodb://localhost/books")
}

func main() {
  http.HandleFunc("/", index)
  http.ListenAndServe(":8080", nil)
}
