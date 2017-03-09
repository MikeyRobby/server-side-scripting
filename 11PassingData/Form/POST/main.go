package main

import(
  "net/http"
  "io"
)

func main(){
  http.HandleFunc("/", index)
  http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request){
  v := req.FormValue("q")
  w.Header().Set("Content-Type", "text/html; charset=utf-8")
  io.WriteString(w, `
    <form method="POST">
    <input type="text" name="q">
    <input type="submit">
    </form>
    <br>`+v)
}
