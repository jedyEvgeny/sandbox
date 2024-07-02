package main

import (
	"testing"
)

type inOunt struct {
	input    string
	expected rune
}

func TestFindMaxFrequencyRune(t *testing.T) {
	sliceTests := []inOunt{
		{"A", 'A'},
		{"AABBBC", 'B'},
		{"AABBBCCCC", 'C'},
		{"", 0},
		{"   ", ' '},
		{"1234555555", '5'},
		{"12345550555", '5'},
		{"12а345F550555", '5'},
	}

	for _, el := range sliceTests {
		t.Run(el.input, func(t *testing.T) {
			result := findMaxFrequencyRune(el.input)
			if result != el.expected {
				t.Errorf("Expected %q, but got %q for input %q", el.expected, result, el.input)
			}
		})
	}
}
