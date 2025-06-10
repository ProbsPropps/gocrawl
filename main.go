package main

import (
	"fmt"
	"os"
)

func main() {
	urlArgs := os.Args[1:]

	if len(urlArgs) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	} 
	
	if len(urlArgs) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	
	url := urlArgs[0]

	fmt.Printf("starting crawl of: %v\n", url)
	fmt.Print(getHTML(url))
}
