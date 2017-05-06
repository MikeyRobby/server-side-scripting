package main

import (
  "html/template"
  "net/http"
  "github.com/satori/go.uuid"
)

var tpl *template.Template
var dbUsers = map[string]user{}       // user ID, user
var dbSessions = map[string]session{} // session ID, session

func init() {
  tpl = template.Must(template.ParseGlob("templates/*"))
}

type user struct {
  First string
  Last string
  Username string
  Pass string
  Email string
  Role string
}

type session struct {
  un string
}
var u user


func main() {
  http.HandleFunc("/", index)
  http.HandleFunc("/signup", signUp)
  http.HandleFunc("/login", login)
  http.HandleFunc("/logout", logOut)
  http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
  u := getUser(w, r)
  tpl.ExecuteTemplate(w, "index.html", u)

}

func signUp(w http.ResponseWriter, req *http.Request) {
  if alreadyLoggedIn(w, req) {
    http.Redirect(w, req, "/", http.StatusSeeOther)
    return
  }
  // process form submission
  if req.Method == http.MethodPost {
    // get form values
    un := req.FormValue("username")
    p := req.FormValue("password")
    f := req.FormValue("firstname")
    l := req.FormValue("lastname")
    e := req.FormValue("email")
    r := req.FormValue("role")
    // username taken?
    if _, ok := dbUsers[un]; ok {
      http.Error(w, "Username already taken", http.StatusForbidden)
      return
    }
    // create session
    sID := uuid.NewV4()
    c := &http.Cookie{
      Name:  "session",
      Value: sID.String(),
    }
    http.SetCookie(w, c)
    dbSessions[c.Value] = session{un}
    u = user{f, l, un, p, e, r}
    dbUsers[un] = u
    // redirect
    http.Redirect(w, req, "/", http.StatusSeeOther)
    return
  }
  tpl.ExecuteTemplate(w, "signup.html", u)

}

func login(w http.ResponseWriter, r *http.Request) {
  if alreadyLoggedIn(w, r) {
    http.Redirect(w, r, "/", http.StatusSeeOther)
    return
  }
  var u user
  // process form submission
  if r.Method == http.MethodPost {
    un := r.FormValue("username")
    p := r.FormValue("password")
    // is there a username?
    u, ok := dbUsers[un]
    if !ok {
      http.Error(w, "Username and/or password do not match", http.StatusForbidden)
      return
    }
    if u.Pass != p {
      http.Error(w, "Username and/or password do not match", http.StatusForbidden)
      return
    }
    // create session
    sID := uuid.NewV4()
    c := &http.Cookie{
      Name:  "session",
      Value: sID.String(),
    }
    http.SetCookie(w, c)
    dbSessions[c.Value] = session{un}
    http.Redirect(w, r, "/", http.StatusSeeOther)
    return
  }
  tpl.ExecuteTemplate(w, "login.html", u)
}

func logOut(w http.ResponseWriter, r *http.Request)  {
  if !alreadyLoggedIn(w, r) {
    http.Redirect(w, r, "/", http.StatusSeeOther)
    return
  }
  c, _ := r.Cookie("session")
  // delete the session
  delete(dbSessions, c.Value)
  // remove the cookie
  c = &http.Cookie{
    Name:   "session",
    Value:  "",
    MaxAge: -1,
  }
  http.SetCookie(w, c)

  http.Redirect(w, r, "/login", http.StatusSeeOther)

}
