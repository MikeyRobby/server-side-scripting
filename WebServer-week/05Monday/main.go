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
  http.HandleFunc("/", index)
  http.HandleFunc("/contact.html", contact)
  http.HandleFunc("/projects.html", projects)
  http.HandleFunc("/about.html", about)
  http.ListenAndServe(":8080",nil)
}

func index(w http.ResponseWriter, req *http.Request) {
  tpl.ExecuteTemplate(w, "index.html", nil)
}

func contact(w http.ResponseWriter, req *http.Request){
  tpl.ExecuteTemplate(w, "contact.html", nil)
}

func projects(w http.ResponseWriter, req *http.Request){
  tpl.ExecuteTemplate(w, "projects.html", nil)
}

func about(w http.ResponseWriter, req *http.Request){
  tpl.ExecuteTemplate(w, "about.html", nil)
}
