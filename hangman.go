package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
	"unicode"
)

var wordsToGuess = []string{
	"IPHONE",
	"MACBOOK",
	"IPAD",
	"IOS",
	"MACOS",
	"IMAC",
	"MACMINI",
	"APPLEWATCH",
	"AIRPODS",
	"ITUNES",
	"ICLOUD",
	"SIRI",
	"APPLEMUSIC",
	"APPLETV",
	"APPPRO",
	"MAGICMOUSE",
	"MAGICKB",
	"FACEID",
	"TOUCHID",
}

var guessedWord []string
var guessedLetters []string
var attempts = 11
var wordToFind string

func init() {
	rand.Seed(time.Now().UnixNano())
	// Choisissez un mot au hasard parmi la liste
	resetGame()
}

func resetGame() {
	wordToFind = chooseNewWord()
	guessedWord = make([]string, len(wordToFind))
	attempts = 11
	initGuessedWord()
}

func initGuessedWord() {
	for i := range guessedWord {
		guessedWord[i] = "_"
	}
}

func Hangman() {
	fmt.Println(wordToFind)
}

func InWord(l string) bool {
	l = strings.ToUpper(l) // Convertir la lettre en majuscules

	for _, letter := range wordToFind {
		if string(unicode.ToUpper(letter)) == l {
			return true
		}
	}
	return false
}

func revealGuessedWord(letter string) {
	for i := range wordToFind {
		if string(wordToFind[i]) == letter {
			guessedWord[i] = letter
		}
	}
}

func chooseNewWord() string {
	// Choisissez un mot au hasard parmi la nouvelle liste
	return wordsToGuess[rand.Intn(len(wordsToGuess))]
}