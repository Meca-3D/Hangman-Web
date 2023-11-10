package main

import (
	"html/template"
	"net/http"
)

var tmpl *template.Template

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTeamplate(w, "index.html", nil)
}

func main() {
	http.HandleFunc("/", HomeHandler)
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		return
	}
}

func Init() {
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
}
