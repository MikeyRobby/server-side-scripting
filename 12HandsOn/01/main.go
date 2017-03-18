/*
ListenAndServe on port 8080 of localhost

For the default route "/" Have a func called "foo" which writes to the response "foo ran"

For the route "/dog/" Have a func called "dog" which parses a template called "dog.gohtml" and writes to the response "

This is from dog

" and also shows a picture of a dog when the template is executed.
Use "http.ServeFile" to serve the file "dog.jpeg"
*/

package main

import (
  "net/http"
  "html/template"
  "io"
)

var tpl *template.Template

func init()  {
  tpl = template.Must(template.ParseGlob("templates/*"))
}

func main()  {
  http.HandleFunc("/", foo)
  http.HandleFunc("/dog/", dog)
  http.HandleFunc("/main.css", css)
  http.HandleFunc("/dog-coding.jpeg", pic)
  http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
  io.WriteString(w, `<h1>foo ran</h1>`)
}

func dog(w http.ResponseWriter, r *http.Request) {
  tpl.ExecuteTemplate(w, "dog.html", nil)
}

func css(w http.ResponseWriter, r *http.Request) {
  http.ServeFile(w,r,"main.css")
}

func pic(w http.ResponseWriter, r *http.Request) {
  http.ServeFile(w,r,"dog-coding.jpeg")
}
