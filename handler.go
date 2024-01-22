package main

import (
	"html/template"
	"net/http"
)

type Web struct {
	Word string
}

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home")
}

func Solo(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles("template/solo.page.html"))
	r.ParseForm()
	letter := r.Form.Get("letter")
	InWord(letter)
	azerty := Web{
		Word: wordToFind,
	}
	template.Execute(w, azerty)

}

func renderTemplate(w http.ResponseWriter, html string) {
	t, err := template.ParseFiles("./template/" + html + ".page.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}
