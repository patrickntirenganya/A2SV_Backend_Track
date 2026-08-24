package main

import "fmt"

func main() {
	tests := []string{
		"madam",
		"racecar",
		"Hello",
		"A man, a plan, a canal: Panama",
		"Was it a car or a cat I saw?",
	}

	for _, text := range tests {
		fmt.Printf("%q -> %t\n", text, IsPalindrome(text))
	}
}
