package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {

	cfg.concurrencyChannel <- struct{}{}
	defer func() {
		<-cfg.concurrencyChannel
		cfg.ws.Done()
	}()
	
	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't parse URL %s : %v\n", rawCurrentURL, err)
		return
	}

	if currentURL.Hostname() != cfg.baseURL.Hostname() {
		return
	}

	normCurrentURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't normalize URL: %v", err)
		return
	}

	fmt.Printf("crawling %s\n", rawCurrentURL)

	firstVisit := cfg.addPageVisit(normCurrentURL)
	if !firstVisit {
		return
	}

	cfg.mu.Lock()
	if len(cfg.pages) >= cfg.maxPages {
		cfg.mu.Unlock()
		return
	}
	cfg.mu.Unlock()
	
	htmlBody, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't retrieve HTML: %v", err)
		return
	}
	
	urls, err := getURLsFromHTML(htmlBody, cfg.baseURL.String())
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't get URLs from HTML: %v", err)
		return
	}
	
	for _, url := range urls {
		cfg.ws.Add(1)
		go cfg.crawlPage(url)
	}
}
