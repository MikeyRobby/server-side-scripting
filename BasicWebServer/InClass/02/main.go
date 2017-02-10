package main

import (
  "net/http"
  "html/template"
)

var tpl *template.Template

func init() {
  tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
  http.HandleFunc("/", index)
  http.HandleFunc("/contact", contact)
  http.HandleFunc("/about", about)
  http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
  tpl.ExecuteTemplate(w, "index.html", nil)
}

func contact(w http.ResponseWriter, r *http.Request){
  tpl.ExecuteTemplate(w, "contact.html", nil)
}

func about(w http.ResponseWriter, r *http.Request) {
  tpl.ExecuteTemplate(w,"about.html",nil)
}

