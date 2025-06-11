package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	resp, err := http.Get(rawURL)
	if err != nil {
		return "", fmt.Errorf("got network error: %v", err)
	}
	defer resp.Body.Close()


	if resp.StatusCode > 399 {
		return "", fmt.Errorf("Status Code Error: %d", resp.StatusCode)
	}
	
	contentType := resp.Header.Get("content-type")

	if !strings.Contains(contentType, "text/html") {
		return "", fmt.Errorf("Content type mismatch: %s", contentType)
	}
	
	html, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Error reading html file: %v", err)
	}
	
	return string(html), nil

}
