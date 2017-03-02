package main

import "net/http"

func main() {
  http.HandleFunc("/boat", boat)
  http.HandleFunc("/bridge", bridge)
  http.HandleFunc("/river", river)
  http.HandleFunc("/road/straight", sroad)
  http.HandleFunc("/road/windy", wroad)
  http.ListenAndServe(":8080", nil)
}

func boat(w http.ResponseWriter, req *http.Request){
  http.ServeFile(w, req, "./pic/boat-on-water.jpg")
}

func bridge(w http.ResponseWriter, req *http.Request){
  http.ServeFile(w, req, "./pic/Bridge-tracks.jpg")
}

func river(w http.ResponseWriter, req *http.Request) {
  http.ServeFile(w,req,"./pic/river-canyon.jpg")
}

func sroad(w http.ResponseWriter, req *http.Request) {
  http.ServeFile(w, req, "./pic/straight-road.jpg")
}

func wroad(w http.ResponseWriter, req *http.Request) {
  http.ServeFile(w, req, "./pic/windy-road.jpg")
}
