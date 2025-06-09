package main

import (
	"strings"
	"net/url"

	"golang.org/x/net/html"
)


func getURLsFromHTML (htmlBody, rawBaseURL string) ([]string, error) {
	var urls []string
	
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, err
	}

	doc, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return nil, err
	}
	for n := range doc.Descendants() {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					href, err := url.Parse(a.Val)
					if err != nil {
						return nil, err
					}

					resolvedURL := baseURL.ResolveReference(href)
					urls = append(urls, resolvedURL.String())
				}
			}
		}
	}
	return urls, nil
}
