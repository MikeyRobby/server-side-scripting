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
  http.HandleFunc("/main.css", css)
  http.HandleFunc("/w640.jpg", img)
  http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
  tpl.ExecuteTemplate(w, "index.html", nil)
}

func css(w http.ResponseWriter, req *http.Request){
  http.ServeFile(w,req,"main.css")
}

func img(w http.ResponseWriter, req *http.Request){
  http.ServeFile(w, req, "w640.jpg")
}

