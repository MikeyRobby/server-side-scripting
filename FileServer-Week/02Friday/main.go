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
  http.HandleFunc("/burger.jpg", pic)
  http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
  tpl.ExecuteTemplate(w, "index.html", nil)
}

func style(w http.ResponseWriter, req *http.Request) {
  http.ServeFile(w, req, "main.css")
}

func pic(w http.ResponseWriter, req *http.Request) {
  http.ServeFile(w, req, "burger.jpg")
}
