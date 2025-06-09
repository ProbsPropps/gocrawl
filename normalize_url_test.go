package main

import (
	"testing"
)

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name 		string
		inputURL 	string
		expected 	string
	}{
		{
			name: "remove scheme",
			inputURL: "https://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name: "remove beginning",
			inputURL: "https://github.com/ProbsPropps/gocrawl",
			expected: "github.com/ProbsPropps/gocrawl",
		},
		{
			name: "remove extra forward slash",
			inputURL: "https://youtube.com/",
			expected: "youtube.com",
		},
		{
			name: "handle no port",
			inputURL: "google.com",
			expected: "google.com",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T){
			actual, err := normalizeURL(tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: uexpected error: %v", i, tc.name, err)
				return
			}
			if actual != tc.expected {
				t.Errorf("Test %v - '%s' FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
	
}
