package main

import "testing"

func TestGetURLFromHTML(t *testing.T) {
	tests := []struct {
		name string
		inputURL string
		inputBody string
		expected []string
	}{
		{
			name: "absolute and relative URLS",
			inputURL: "https://blog.boot.dev",
			inputBody: `
		<html>
			<body>
				<a href="/path/one">
					<span>Boot.dev</span>
				</a>
				<a href="https://other.com/path/one">
					<span>Boot.dev</span>
				</a>
			</body>
		</html>
		`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},		
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T){
			actual, err := getURLsFromHTML(tc.inputBody, tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if len(actual) != len(tc.expected) {
				t.Errorf("Test %v = '%s' FAIL: Mismatch lengths", i, tc.name)
			}
			for j, url := range actual{
				if url != tc.expected[j] {
					t.Errorf("Test %v - '%s' FAIL: expected URL(s): %v, actual: %v", i, tc.name, tc.expected, actual)
				}	
			}
			
		})
	}
}
