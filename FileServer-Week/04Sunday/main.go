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
  http.HandleFunc("/main.css", style)
  http.HandleFunc("/The-Exorcism-of-Emily-Rose.jpg", pic)
  http.ListenAndServe(":8080",nil)
}

func index(w http.ResponseWriter, r *http.Request){
  tpl.ExecuteTemplate(w, "index.html", nil)
}

func style(w http.ResponseWriter, r *http.Request){
  http.ServeFile(w,r,"main.css")
}

func pic(w http.ResponseWriter, r *http.Request) {
  http.ServeFile(w,r,"The-Exorcism-of-Emily-Rose.jpg")
}
