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
	http.HandleFunc("/int", ints)
	http.HandleFunc("/bool", bool)
	http.HandleFunc("/map", mapAll)
	http.HandleFunc("/mapValues", mapValues)
	http.HandleFunc("/slice", slice)
	http.HandleFunc("/string", strings)
	http.HandleFunc("/struct", structP)
	http.ListenAndServe(":8080", nil)
}

func ints(w http.ResponseWriter, req *http.Request) {
	i := 12
	tpl.ExecuteTemplate(w, "int.html", i)
}

func bool(w http.ResponseWriter, req *http.Request) {
	b := true
	tpl.ExecuteTemplate(w, "bool.html", b)
}

func mapAll(w http.ResponseWriter, req *http.Request) {
	m := map[string]string{
		"Oxygen: ":   "O2",
		"Cloride: ":  "Cl2",
		"Nitrogen: ": "N2",
		"Hydrogen: ": "H",
	}
	tpl.ExecuteTemplate(w, "map.html", m)
}

func mapValues(w http.ResponseWriter, req *http.Request) {
	m := map[string]float64{
		"Hydrogen: ": 1.008,
		"Oxygen: ":   15.999,
		"Nitrogen: ": 14.007,
		"Iron: ":     55.845,
	}
	tpl.ExecuteTemplate(w, "mapValues.html", m)

}

func slice(w http.ResponseWriter, req *http.Request) {
	xs := []string{"Acid-Base", "Precipitation", "Gas Evolution", "Combustion"}
	tpl.ExecuteTemplate(w, "slice.html", xs)

}

func strings(w http.ResponseWriter, req *http.Request) {
	s := "MikeyRobby"
	tpl.ExecuteTemplate(w, "string.html", s)

}

func structP(w http.ResponseWriter, req *http.Request) {

	type Person struct {
		Fname    string
		Lname    string
		Age      int
		Username string
	}

	p1 := Person{"Michael", "Roberts", 22, "MikeyRobby"}
	p2 := Person{"Bennie", "Jones", 23, "Knock-Knock"}
	xp := []Person{}
	xp = append(xp, p1, p2)

	tpl.ExecuteTemplate(w, "struct.html", xp)
}
