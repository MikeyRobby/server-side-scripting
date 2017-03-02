package main

import(
	"net/http"
	"io"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/main.css", css)
	http.HandleFunc("/forest.jpeg", ocean)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	io.WriteString(w, `<!DOCTYPE html>
  <html lang="en">
  <head>
  <meta charset="utf-8">
  <link rel="stylesheet" href="main.css">
  </head>
  <body>
    <h1>Hello, Are you Here?</h1>
  </body>
  </html>`)

}

func css(w http.ResponseWriter, req *http.Request){
	http.ServeFile(w, req, "main.css")
}

func ocean(w http.ResponseWriter, req *http.Request){
	http.ServeFile(w, req, "forest.jpeg")
}
