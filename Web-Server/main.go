package main

import (
"net/http"
"fmt"
)

func main(){
  http.HandleFunc("/", index)
  http.HandleFunc("/about", about)
  http.HandleFunc("/contact", contact)
  http.HandleFunc("/web", web)
  http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, r *http.Request){
  fmt.Fprint(w, `<!DOCTYPE html>
    <html lang="en">
    <head>
      <meta charset="UTF-8">
      <title>Index</title>
      <style>
      html {
      font-size: 48px;
      }
      </style>
    </head>
    <body>
    Hello from index
    <br><br>
    <a href="/">Index</a>
    <a href="/about">about</a>
    <a href="/contact">contact</a>
    </body>
    </html>`)
}
func about(w http.ResponseWriter, r *http.Request){
  fmt.Fprint(w, `<!DOCTYPE html>
    <html lang="en">
    <head>
      <meta charset="UTF-8">
      <title>about</title>
      <style>
      html {
      font-size: 48px;
      }
      </style>
    </head>
    <body>
    Hello from index
    <br><br>
    <a href="/">Index</a>
    <a href="/about">about</a>
    <a href="/contact">contact</a>
    </body>
    </html>`)
}
func contact(w http.ResponseWriter, r *http.Request){
  fmt.Fprint(w, `<!DOCTYPE html>
    <html lang="en">
    <head>
      <meta charset="UTF-8">
      <title>Contact</title>
      <style>
      html {
      font-size: 48px;
      }
      </style>
    </head>
    <body>
    Hello from index
    <br><br>
    <a href="/">Index</a>
    <a href="/about">about</a>
    <a href="/contact">contact</a>
    </body>
    </html>`)
}
func web(w http.ResponseWriter, r *http.Request){
  fmt.Fprint(w, `<!DOCTYPE html>
    <html lang="en">
    <head>
      <meta charset="UTF-8">
      <title>Index</title>
      <style>
      html, body, .hero-img {
          padding: 0;
          margin: 0;
          box-sizing: border-box;
          height: 100%;
      }

      .hero-img {
          display: flex;
          background-image: url("http://www.planwallpaper.com/static/images/6-940622110b39cad584098e6749eac708.jpg");
          background-size: cover;
          background-color: #444444;
          background-repeat: no-repeat;
          background-position: top center;
          flex-grow: 1;
          justify-content: center;
          align-items: center;


      }

      .btf {
          margin: 0;
          padding: 0;
          box-sizing: border-box;
          font-family: cursive;
          text-align: center ;
          font-size: x-large;
      }

      </style>
    </head>
    <body>

    <div class="hero-img">
    </div>

    <div class="btf">
        <p> I saw the birth of the universe and watched as time ran out, moment by moment, until nothing remained. No time, no space. Just me! I walked in universes where the laws of physics were devised by the mind of a madman! And I watched universes freeze and creation burn! I have seen things you wouldn't believe! I have lost things you will never understand! And I know things, secrets that must never be told, knowledge that must never be spoken! </p>


    </div>



    </body>`)
}
