package main

import (
	"reflect"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	words := []string{"дaдa", "пятак", "пятка", "пятка", "тяпка", "листок", "слиток", "столик"}
	expected := map[string][]string{
		"акптя":  {"тяпка", "пятак", "пятка"},
		"иклост": {"листок", "слиток", "столик"},
	}

	result := findAnagrams(words)

	if reflect.DeepEqual(result, expected) {
		t.Errorf("Expected: %v, but got: %v", expected, result)
	}
}

func TestSortLetters(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"пятак", "акптя"},
		{"слиток", "иклост"},
	}

	for _, test := range tests {
		result := sortLetters(test.input)
		if result != test.expected {
			t.Errorf("For input %s, expected: %s, but got: %s", test.input, test.expected, result)
		}
	}
}

