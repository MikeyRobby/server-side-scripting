package main

import (
  "net/http"
  "html/template"
)

var tpl *template.Template

func init()  {
  tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
  http.HandleFunc("/", index)
  http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("images"))))
  http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
  s := "PARADISE FOUND"
  tpl.ExecuteTemplate(w, "index.html", s)
}
