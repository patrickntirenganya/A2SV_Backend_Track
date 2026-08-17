package main

import (
	"reflect"
	"testing"
)

func TestWordFrequency(t *testing.T) {
	input := "Hello, hello world! Go is great, and Go is fast."

	expected := map[string]int{
		"hello": 2,
		"world": 1,
		"go":    2,
		"is":    2,
		"great": 1,
		"and":   1,
		"fast":  1,
	}

	result := WordFrequency(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}
