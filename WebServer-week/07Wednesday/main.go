package main

import (
  "net/http"
  "html/template"
  "io/ioutil"
)

var tpl *template.Template

func init() {
  tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
  http.HandleFunc("/", index)
  http.HandleFunc("/projects.html", projects)
  http.HandleFunc("/blog.html", blog)
  http.HandleFunc("/about.html", about)
  http.HandleFunc("/contact.html", contact)
  http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request)  {
  tpl.ExecuteTemplate(w,"index.html", nil)
}

func projects(w http.ResponseWriter, req *http.Request) {
  tpl.ExecuteTemplate(w, "projects.html", nil)
}

func blog(w http.ResponseWriter, req *http.Request) {
  r ,_ := ioutil.ReadFile("posts/message")
  tpl.ExecuteTemplate(w, "blog.html", string(r))
}

func about(w http.ResponseWriter, req *http.Request){
  tpl.ExecuteTemplate(w, "about.html", nil)
}

func contact(w http.ResponseWriter, req *http.Request) {
  tpl.ExecuteTemplate(w, "contact.html", nil)
}



