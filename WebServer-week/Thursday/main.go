package main

import (
  "net/http"
  "html/template"
)

func init() {
  tpl = template.Must(template.ParseGlob("templates/*"))

}

var tpl *template.Template

func main() {
  http.HandleFunc("/", index)
  http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, req *http.Request) {
  tpl.ExecuteTemplate(w, "index.html", nil)
}

