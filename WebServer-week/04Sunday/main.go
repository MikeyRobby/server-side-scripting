package main

import (
  "net/http"
  "html/template"
)

var tpl *template.Template

func init() {
  http.HandleFunc("/projects", projects)
  tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
  http.ListenAndServe(":8080", nil)
}

func projects(w http.ResponseWriter, req *http.Request){
  tpl.ExecuteTemplate(w, "projects.html", nil)
}

