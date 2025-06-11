package main

import (
	"net/url"
	"strings"
)

func normalizeURL(inputURL string) (string, error) {
	urlData, err := url.Parse(inputURL)
	if err != nil {
		return "", err
	}

	fullURL := urlData.Host + urlData.Path
	
	fullURL = strings.ToLower(fullURL)

	fullURL = strings.TrimSuffix(fullURL, "/")

	return fullURL, nil
}
