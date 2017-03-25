package main

import (
  "net/http"
  "github.com/satori/go.uuid"
  "html/template"
)

type user struct {
  UserName string
  First    string
  Last     string
}

var tpl *template.Template
var dbUsers = map[string]user{} // user ID, user
var dbSessions = map[string]string{} // session ID, user ID

func init() {
  tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
  http.HandleFunc("/", index)
  http.HandleFunc("/home.html", home)
  http.HandleFunc("/account.html", account)
  http.HandleFunc("/matrix.html", matrix)
  http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
  c, err := r.Cookie("session")
  if err != nil {
    sID := uuid.NewV4()
    c = &http.Cookie{
      Name: "session",
      Value: sID.String(),
    }
    http.SetCookie(w,c)
  }

  var u user
  if un, ok := dbSessions[c.Value]; ok {
    u = dbUsers[un]
  }

  // process form submission
  if r.Method == http.MethodPost {
    un := r.FormValue("username")
    f := r.FormValue("firstname")
    l := r.FormValue("lastname")
    u = user{un, f, l}
    dbSessions[c.Value] = un
    dbUsers[un] = u
  }

  tpl.ExecuteTemplate(w, "index.html", u)
}

func home(w http.ResponseWriter, req *http.Request) {
  // get cookie
  c, err := req.Cookie("session")
  if err != nil {
    http.Redirect(w, req, "/", http.StatusSeeOther)
    return
  }
  un, ok := dbSessions[c.Value]
  if !ok {
    http.Redirect(w, req, "/", http.StatusSeeOther)
    return
  }
  u := dbUsers[un]
  tpl.ExecuteTemplate(w, "home.html", u)
}

func account(w http.ResponseWriter, req *http.Request) {
  // get cookie
  c, err := req.Cookie("session")
  if err != nil {
    http.Redirect(w, req, "/", http.StatusSeeOther)
    return
  }
  un, ok := dbSessions[c.Value]
  if !ok {
    http.Redirect(w, req, "/", http.StatusSeeOther)
    return
  }
  u := dbUsers[un]
  tpl.ExecuteTemplate(w, "account.html", u)
}

func matrix(w http.ResponseWriter, req *http.Request) {
  // get cookie
  c, err := req.Cookie("session")
  if err != nil {
    http.Redirect(w, req, "/", http.StatusSeeOther)
    return
  }
  un, ok := dbSessions[c.Value]
  if !ok {
    http.Redirect(w, req, "/", http.StatusSeeOther)
    return
  }
  u := dbUsers[un]
  tpl.ExecuteTemplate(w, "matrix.html", u)
}

