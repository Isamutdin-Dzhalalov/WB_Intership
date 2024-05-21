package main

import (
	"testing"
)

func TestUnpackingString(t *testing.T) {
	testCases := []struct {
		input       string
		expected    string
		shouldError bool
	}{
		{"a4bc2d5e", "aaaabccddddde", false},
		{"abcd", "abcd", false},              
		{"45", "", true},                     
		{"", "", false},                      
	}

	for _, tc := range testCases {
		result, err := UnpackingString(tc.input)
		if tc.shouldError {
			if err == nil {
				t.Errorf("Expected an error for input '%s' but got nil", tc.input)
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error for input '%s': %v", tc.input, err)
			}
			if result != tc.expected {
				t.Errorf("Unexpected result for input '%s': got '%s', expected '%s'", tc.input, result, tc.expected)
			}
		}
	}
}

