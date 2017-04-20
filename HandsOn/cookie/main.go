// Create a Web App that writes a cookie

package main

import (
  "net/http"
  "html/template"
  "fmt"
)
// Setting up templates
var tpl *template.Template

func init()  {
  tpl = template.Must(template.ParseGlob("templates/*"))
}
// End Setting up templates
func main() {
  // Routes
  http.HandleFunc("/", index)
  // Setting Server
  http.ListenAndServe(":8080", nil)
}
// Creating Routes
func index(w http.ResponseWriter, req *http.Request) {
  // Creating Cookies
  var cook *http.Cookie
  cook, err := req.Cookie("Success")
  if err != nil {
    fmt.Println(err)
  }
  // Setting Cookie Values
  cook = &http.Cookie{
    Name: "Success",
    Value: "Cookie Created",
  }
  http.SetCookie(w,cook)
tpl.ExecuteTemplate(w, "index.html", cook)
}
