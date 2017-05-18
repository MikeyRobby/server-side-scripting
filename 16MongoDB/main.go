package main

import (
  "html/template"
  "net/http"
  "fmt"
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
)

type Book struct {
  Title string `json:"title" bson:"title"`
  Author string `json:"author" bson:"author"`
}

var tpl *template.Template
var DB *mgo.Database
var Books *mgo.Collection


func init()  {
  tpl = template.Must(template.ParseGlob("templates/*"))
  s, err := mgo.Dial("mongodb://localhost/books")
  if err != nil{
    panic(err)
  }

  if err = s.Ping(); err != nil {
    panic(err)
  }

  DB = s.DB("bookstore")
  Books = DB.C("books")

  fmt.Println("You are connected to your mongo database.")
}

func main() {
  http.HandleFunc("/", index)
  http.ListenAndServe(":8080", nil)
}

func AllBooks() ([]Book, error) {
  books := []Book{}
  err := Books.Find(bson.M{}).All(&books)
  if err != nil {
    return nil, err
  }
  return books, nil
}

func index(w http.ResponseWriter, r *http.Request) {

  bks, err := AllBooks()
  if err != nil {
    http.Error(w, http.StatusText(500)+err.Error(), http.StatusInternalServerError)
  }
  tpl.ExecuteTemplate(w, "index.html", bks)
}

