package main

import "unicode"

func IsPalindrome(text string) bool {
	var cleaned []rune

	for _, r := range text {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			cleaned = append(cleaned, unicode.ToLower(r))
		}
	}

	for i, j := 0, len(cleaned)-1; i < j; i, j = i+1, j-1 {
		if cleaned[i] != cleaned[j] {
			return false
		}
	}

	return true
}
