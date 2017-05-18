package main

import (
  "html/template"
  "net/http"
  "fmt"
)

var tpl *template.Template

func init()  {
  tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
  http.HandleFunc("/", index)
  http.HandleFunc("/foo", foo)
  http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
  tpl.ExecuteTemplate(w, "index.html", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
  s := `Here is some text from foo`
  fmt.Fprintf(w, s)
}