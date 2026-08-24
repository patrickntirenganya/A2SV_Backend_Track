package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"madam", true},
		{"racecar", true},
		{"Hello", false},
		{"A man, a plan, a canal: Panama", true},
		{"Was it a car or a cat I saw?", true},
		{"12321", true},
		{"12345", false},
		{"", true},
	}

	for _, test := range tests {
		result := IsPalindrome(test.input)

		if result != test.expected {
			t.Errorf(
				"IsPalindrome(%q) = %v, expected %v",
				test.input,
				result,
				test.expected,
			)
		}
	}
}
