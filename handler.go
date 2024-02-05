package main

import (
	"html/template"
	"net/http"
	"strings"
)

type Web struct {
	Word          string
	Attempts      int
	LettersTapped string // Historique des lettres tapées
	Message       string // Message personnalisé
}

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home")
}

func Solo(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles("template/solo.page.html"))
	r.ParseForm()
	letter := r.Form.Get("letter")

	if letter != "" {
		// Convertir la lettre en majuscules
		letter = strings.ToUpper(letter)

		// Vérifier si la lettre a déjà été tapée
		if !contains(guessedLetters, letter) {
			if InWord(letter) {
				revealGuessedWord(letter)
			} else {
				attempts--
			}
			guessedLetters = append(guessedLetters, letter) // Ajouter la lettre à la liste des lettres déjà tapées
		}
	}

	lettersTapped := strings.Join(guessedLetters, " ") // Utilisez les lettres déjà tapées

	gameState := Web{
		Word:          strings.Join(guessedWord, " "),
		Attempts:      attempts,
		LettersTapped: lettersTapped,
		Message:       getMessage(), // Obtenez un message personnalisé
	}
	template.Execute(w, gameState)
}

// Fonction utilitaire pour obtenir des messages personnalisés
func getMessage() string {
	if isWordGuessed() {
		return "Félicitations, vous avez deviné le mot !"
	} else if attempts == 0 {
		return "Dommage, vous n'avez plus d'essais. Le mot était : " + wordToFind
	} else {
		return "Bonne chance !"
	}
}

// Fonction utilitaire pour vérifier si le mot a été deviné
func isWordGuessed() bool {
	for _, letter := range guessedWord {
		if letter == "_" {
			return false
		}
	}
	return true
}

// Fonction utilitaire pour vérifier si un élément est dans une liste
func contains(list []string, element string) bool {
	for _, value := range list {
		if value == element {
			return true
		}
	}
	return false
}

func renderTemplate(w http.ResponseWriter, html string) {
	t, err := template.ParseFiles("./template/" + html + ".page.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}