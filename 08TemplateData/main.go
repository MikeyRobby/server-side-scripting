package main

import (
  "fmt"
  "html/template"
  "net/http"
)

var tpl *template.Template

func init() {
  tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
  // TODO: Create HandleFuncs
  http.ListenAndServe(":8080",nil)
}

func int(w http.ResponseWriter, req *http.Request){
  // TODO: Create int and pass in data
}

func bool(w http.ResponseWriter, req *http.Request){
  // TODO: Create bool value and pass in data

}

func mapAll(w http.ResponseWriter, req *http.Request){
  // TODO: Create map and pass in all data

}

func mapValues(w http.ResponseWriter, req *http.Request){
  // TODO: Create map and pass in key and value data

}

func slice(w http.ResponseWriter, req *http.Request){
  // TODO: Create slice and pass in data

}

func string(w http.ResponseWriter, req *http.Request){
  // TODO: Create string and pass in data

}

func structP(w http.ResponseWriter, req *http.Request){
  // TODO: Create struct and pass in data

}
