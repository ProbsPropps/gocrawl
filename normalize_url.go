package main

import (
	"net/url"
)

func normalizeURL(inputURL string) (string, error) {
	urlData, err := url.Parse(inputURL)
	if err != nil {
		return "", err
	}
	if urlData.Path == "/" || urlData.Path == "" {
		return urlData.Host, nil
	}

	fullURL := urlData.Host + urlData.Path
	return fullURL, nil
}
