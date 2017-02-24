package main

import (
  "html/template"
  "net/http"
)

var tpl *template.Template

func init() {
  tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
  http.HandleFunc("/int", int)
  http.HandleFunc("/bool", bool)
  http.HandleFunc("/map", mapAll)
  http.HandleFunc("/mapValues", mapValues)
  http.HandleFunc("/slice", slice)
  http.HandleFunc("/string", string)
  http.HandleFunc("/struct", structP)
  http.ListenAndServe(":8080",nil)
}

func int(w http.ResponseWriter, req *http.Request){
  i := 12
  tpl.ExecuteTemplate(w, "int.html", i)
}

func bool(w http.ResponseWriter, req *http.Request){
  b := true
  tpl.ExecuteTemplate(w, "bool.html", b)
}

func mapAll(w http.ResponseWriter, req *http.Request){
  m := map[string]string{
    "Oxygen: ": "O2",
    "Cloride: ": "Cl2",
    "Nitrogen: ": "N2",
    "Hydrogen: " : "H",
  }
tpl.ExecuteTemplate(w, "map.html", m)
}

func mapValues(w http.ResponseWriter, req *http.Request){
  m := map[string]float64{
    "Hydrogen: ": 1.008,
    "Oxygen: ": 15.999,
    "Nitrogen: ": 14.007,
    "Iron: ": 55.845,
  }
  tpl.ExecuteTemplate(w, "mapValues.html", m)

}

func slice(w http.ResponseWriter, req *http.Request){
  xs := []string{"Acid-Base", "Precipitation", "Gas Evolution", "Combustion"}
  tpl.ExecuteTemplate(w, "slice.html", xs)

}

func string(w http.ResponseWriter, req *http.Request){
  s := "MikeyRobby"
  tpl.ExecuteTemplate(w, "string.html", s)

}

func structP(w http.ResponseWriter, req *http.Request){
  // TODO: Create struct and pass in data
  //create a struct person.
  // create some values of type person.
  // add those values to a slice of person. pass the value of type []person to a template and display all of the values in the template

}
