package main

import "fmt"
import "sort"

type FoundUrls struct{
	url string
	count int
}


func printReport(pages map[string]int, baseURL string){
	fmt.Printf("=============================\n"+
		"REPORT for %s\n"+
		"=============================\n", baseURL)
	ordered := make([]FoundUrls, 0, len(pages))
	for i, v := range pages{
		Founded := FoundUrls{
			url:i,
			count:v,
		}
		ordered = append(ordered, Founded)
	}

	sort.Slice(ordered, func(i, j int) bool {
		if ordered[i].count == ordered[j].count{
			return ordered[i].url < ordered[j].url
		}
		return ordered[i].count > ordered[j].count
	})

	for _, Founds := range ordered{
		fmt.Printf("Found %d internal links to %s\n", Founds.count, Founds.url)

	}
}