package main

import (
	"strings"
	"unicode"
)

func WordFrequency(text string) map[string]int {
	frequency := make(map[string]int)

	words := strings.FieldsFunc(text, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})

	for _, word := range words {
		word = strings.ToLower(word)
		frequency[word]++
	}

	return frequency
}
