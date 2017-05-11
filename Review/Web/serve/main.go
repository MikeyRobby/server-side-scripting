package main

import (
  "net/http"
  "html/template"
)

type person struct{
  First string
  Last string
  Saying string
}

var tpl *template.Template

func init()  {
  tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
  http.HandleFunc("/", index)
  http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
  p1 := person{
    First: "Mikey",
    Last: "Robby",
    Saying: "I do what I do and I do it well.",
  }
  p2 := person{
    First: "Michael",
    Last: "Roberts",
    Saying: "I play a lot of cards, sometimes I play them right.",
  }

xp := []person{p1, p2}

tpl.ExecuteTemplate(w, "index.html", xp)

}
