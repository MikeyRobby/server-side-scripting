package main

import (
  "io/ioutil"
	"net/http"
	"html/template"
)

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main(){
	http.HandleFunc("/", index)
  http.HandleFunc("/blog", blog)
	http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, req *http.Request){
  tpl.ExecuteTemplate(w, "index.html", nil)
}

func blog(w http.ResponseWriter, req *http.Request){
  r, _ := ioutil.ReadFile("posts/test")
  tpl.ExecuteTemplate(w, "blog.html", string(r))
}

