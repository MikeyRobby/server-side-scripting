package main

import(
  "html/template"
  "log"
  "net/http"
)

var tpl *template.Template

func init()  {
  tpl = template.Must(template.ParseGlob("templates/*"))
}

type person struct {
  UserName string
  Email string
  Member bool
}

func main()  {
  http.HandleFunc("/", index)
  http.Handle("/favicon.ico", http.NotFoundHandler())
  http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {

  u := req.FormValue("user")
  e := req.FormValue("email")
  m := req.FormValue("member") == "on"

  err := tpl.ExecuteTemplate(w, "index.html", person{u, e, m})
  if err != nil{
    http.Error(w, err.Error(), 500)
    log.Fatalln(err)
  }
}
