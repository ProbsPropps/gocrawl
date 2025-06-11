package main

import (
	"fmt"
	"os"
)

func main() {
	urlArgs := os.Args[1:]

	if len(urlArgs) < 1 {
		fmt.Println("no website provided")
		return
	} 
	
	if len(urlArgs) > 1 {
		fmt.Println("too many arguments provided")
		return
	}	
	
	rawBaseURL := urlArgs[0]
	cfg, err := configure(rawBaseURL, 1)
	if err != nil {
		msg := fmt.Errorf("Error when creating config struct, %v", err)
		fmt.Println(msg)
		return
	}
	fmt.Printf("starting crawl of: %v\n", rawBaseURL)
	
	cfg.ws.Add(1)
	cfg.crawlPage(rawBaseURL)
	cfg.ws.Wait()

	for normalizedURL, count := range cfg.pages {
		fmt.Printf("%d -%s\n", count, normalizedURL)
	}

}
