package main

import (
	"strings"
)

var profaneWords = []string{"kerfuffle", "sharbert", "fornax"}

func replaceProfanity(input string) string {
	words := strings.Split(input, " ")
	for i, word := range words {
		// Check if the word (case insensitive) matches any profane word
		lowerWord := strings.ToLower(word)
		for _, profane := range profaneWords {
			if lowerWord == profane {
				words[i] = "****"
			}
		}
	}
	return strings.Join(words, " ")
}
