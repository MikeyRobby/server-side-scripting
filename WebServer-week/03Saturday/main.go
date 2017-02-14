package main

import (
  "net/http"
  "html/template"
)

var tpl *template.Template

func init() {
  http.HandleFunc("/about", about)
  tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
  http.ListenAndServe(":8080",nil)
}

func about(w http.ResponseWriter, r *http.Request) {
  tpl.ExecuteTemplate(w,"about.html", nil)
}

