package main

import (
	"fmt"
	"net/url"
	"sync"
)


type config struct {
	pages 				map[string]int
	baseURL 			*url.URL
	mu 					*sync.Mutex
	concurrencyChannel 	chan struct{}
	ws 					*sync.WaitGroup
	maxPages			int
}

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	if _, visited := cfg.pages[normalizedURL]; visited {
		cfg.pages[normalizedURL]++
		return false
	}
	
	cfg.pages[normalizedURL] = 1
	return true
}

func configure(rawBaseURL string, maxConcurrency, maxPages int) (*config, error) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, fmt.Errorf("Error - configure: couldn't parse URL: %v", err)
	}

	return &config{
		pages: make(map[string]int),
		baseURL: baseURL,
		mu: &sync.Mutex{},
		concurrencyChannel: make(chan struct{}, maxConcurrency),
		ws: &sync.WaitGroup{},
		maxPages: maxPages,
	}, nil
}
