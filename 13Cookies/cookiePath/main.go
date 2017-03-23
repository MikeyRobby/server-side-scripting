package main

import (
  "fmt"
  "net/http"
)

func main()  {
  http.HandleFunc("/", set)
  http.HandleFunc("/read", read)
  http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, r *http.Request) {
  http.SetCookie(w, &http.Cookie{
    Name: "my-cookie",
    Value: "some value",
    Path: "/",
  })

  http.SetCookie(w, &http.Cookie{
    Name: "Payment",
    Value: "some amount of currency",
    Path: "/buy",
  })
  fmt.Fprintln(w, "COOKIE WRITTEN - CHECK YOUR BROWSER")
  fmt.Println(w, "in chrome go to: dev tools / application / cookies")
}

func read(w http.ResponseWriter, r *http.Request) {
  c, err := r.Cookie("my-cookie")
  if err != nil {
    http.Error(w, err.Error(), http.StatusNotFound)
    return
  }

  c2, err2 := r.Cookie("Payment")
  if err2 != nil {
    http.Error(w, err2.Error(), http.StatusNotFound)
    return
  }

  fmt.Fprintln(w, "YOUR COOKIE:", c)
  fmt.Fprintln(w, "YOUR OTHER COOKIE:", c2)
}
