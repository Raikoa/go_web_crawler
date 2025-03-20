package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 4 { //check args 
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(os.Args) > 4 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	rawBaseURL := os.Args[1] //get raw url 
	maxconcur, err := strconv.Atoi(os.Args[2])  //transform maxconcut from string to int
	if err != nil{
		fmt.Printf("error: maxconcurrency not int")
		return
	} 
	maxset, err := strconv.Atoi(os.Args[3]) //transform maxset from string to int
	if err != nil{
		fmt.Println("error: maxsetting not int")
		return 
	}
	cfg, err := configure(rawBaseURL, maxconcur, maxset) //create a configure struct
	if err != nil{
		fmt.Println("error making config")
		return
	}


	fmt.Printf("starting crawl of: %s...\n", rawBaseURL)

	cfg.wg.Add(1)
	go cfg.crawlPage(rawBaseURL)
	cfg.wg.Wait()//wait until all goroutine are finished

	for normalizURL, count := range cfg.pages{
		fmt.Printf("%d - %s\n", count, normalizURL)
	}


	printReport(cfg.pages, rawBaseURL)
}
