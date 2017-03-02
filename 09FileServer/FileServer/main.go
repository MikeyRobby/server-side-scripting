package main

import(
  "net/http"
  "io"
)
func main()  {
  http.HandleFunc("/", dog)
  http.HandleFunc("/soccer", soccerDog)
  http.HandleFunc("/face", faceDog)
  http.HandleFunc("/run", runningDog)
  http.HandleFunc("/hot", hotDog)
  http.Handle("/res/", http.StripPrefix("/res", http.FileServer(http.Dir("./res"))))
  http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request){
  w.Header().Set("Content-Type", "text/html; charset=utf-8")
  io.WriteString(w, `<img src="/res/dog-coding.jpeg">`)
}

func soccerDog(w http.ResponseWriter, req *http.Request){
  w. Header().Set("Content-Type", "text/html; charset=utf-8")
  io.WriteString(w, `<img src="/res/dog-soccer.jpeg">`)
}

func faceDog(w http.ResponseWriter, req *http.Request)  {
  w.Header().Set("Content-Type", "text/html; charset=utf-8")
  io.WriteString(w, `<img src="/res/dog-face.jpeg">`)
}

func runningDog(w http.ResponseWriter, req *http.Request)  {
  w.Header().Set("Content-Type", "text/html; charset=utf-8")
  io.WriteString(w, `<img src="/res/dog-running.jpeg">`)
}

func hotDog(w http.ResponseWriter, req *http.Request)  {
  w.Header().Set("Content-Type", "text/html; charset=utf-8")
  io.WriteString(w, `<img src="/res/dog-hot.jpeg">`)
}
