/*
Serve the files in the "starting-files" folder

Use "http.FileServer"
*/

package main

import (
  "net/http"
)

func main() {
  http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))
}
