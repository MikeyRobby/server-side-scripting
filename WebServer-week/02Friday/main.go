package main

import (
  "net/http"
  "html/template"
)

var tpl *template.Template

func init() {
  tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
  http.HandleFunc("/contact", contact)
  http.ListenAndServe(":8080", nil)
}

func contact(w http.ResponseWriter, r *http.Request) {
  tpl.ExecuteTemplate(w, "contact.html", nil)
}

