package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
	"unicode"
)

var wordToFind = "IPHONE"
var guessedWord []string
var guessedLetters []string
var attempts = 6

func initGame() {
	rand.Seed(time.Now().UnixNano())
	guessedWord = make([]string, len(wordToFind))
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