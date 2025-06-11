package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	urlArgs := os.Args[1:]

	if len(urlArgs) < 1 {
		fmt.Println("no website provided")
		return
	} 
	
	if len(urlArgs) != 3 {
		fmt.Println("Format: <url> <max concurrency size> <max pages>")
		return
	}
		
	
	rawBaseURL := urlArgs[0]
	conMax, err := strconv.Atoi(urlArgs[1])
	if err != nil {
		fmt.Println("Error when reading second value: must be a number")
		return
	}
	pageMax, err := strconv.Atoi(urlArgs[2])
	if err != nil {
		fmt.Println("Error when reading third value: must be a number")
		return
	}

	cfg, err := configure(rawBaseURL, conMax, pageMax)
	if err != nil {
		msg := fmt.Errorf("Error when creating config struct, %v", err)
		fmt.Println(msg)
		return
	}
	
	fmt.Printf("starting crawl of: %v\n", rawBaseURL)
	
	cfg.ws.Add(1)
	cfg.crawlPage(rawBaseURL)
	cfg.ws.Wait()

	printReport(cfg.pages, rawBaseURL)

}
