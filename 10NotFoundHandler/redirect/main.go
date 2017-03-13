package main

import(
  "net/http"
  "fmt"
)

func main(){
  http.HandleFunc("/", index)
  http.Handle("/favicon.ico", http.NotFoundHandler())
  http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request){
  if req.URL.Path != "/" {
    http.NotFound(w, req)
    return
  }
  fmt.Println(req.URL.Path)
  fmt.Fprintln(w, "go look at your terminal")
}
