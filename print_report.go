package main

import (
	"fmt"
	"sort"
)

func printReport(pages map[string]int, baseURL string) {
	fmt.Println("====================")
	fmt.Printf("REPORT for %s\n", baseURL)
	fmt.Println("====================")
	
	sortedPages := sortPages(pages)
	
	for _, page := range sortedPages {
		fmt.Printf("Found %d internal links to %s\n", page.visited, page.url)
	}
}

func sortPages(pages map[string]int) []pageStruct {
	var pageSorted []pageStruct

	for url, visited := range pages {
		pageSorted = append(pageSorted, pageStruct{url: url, visited: visited,})
	}
	sort.Slice(pageSorted, func(i, j int) bool {
		return pageSorted[j].visited < pageSorted[i].visited
	})
	return pageSorted

}

type pageStruct struct {
	url 	string
	visited int
}
