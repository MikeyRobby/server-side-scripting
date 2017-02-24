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
  http.Handle("/res/", http.StripPrefix("/res",http.FileServer(http.Dir("/pic"))))
  http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, req *http.Request)  {
  tpl.ExecuteTemplate(w,"index.html", nil)
}

func projects(w http.ResponseWriter, req *http.Request) {
  tpl.ExecuteTemplate(w, "projects.html", nil)
}

func blog(w http.ResponseWriter, req *http.Request) {
  r,_ := ioutil.ReadFile("posts/message.txt")
  r2 ,_ := ioutil.ReadFile("posts/listen.txt")
  r3,_ := ioutil.ReadFile("posts/fear.txt")
  r4 ,_ := ioutil.ReadFile("posts/change.txt")
  r5 ,_ := ioutil.ReadFile("posts/scared.txt")

  p := []string{string(r), string(r2), string(r3), string(r4), string(r5)}

  tpl.ExecuteTemplate(w, "blog.html", p)
}

func about(w http.ResponseWriter, req *http.Request){
  tpl.ExecuteTemplate(w, "about.html", nil)
}

func contact(w http.ResponseWriter, req *http.Request) {
  w.Header().Set("content-type", "text/html; charset=utf-8")
  tpl.ExecuteTemplate(w, "contact.html", nil)

}


