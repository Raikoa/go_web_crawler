package main
import (
	"net/url"
	"fmt"
)




func (cfg *config) crawlPage(currentURL string){
	
	cfg.concurrencyControl <- struct{}{}//start one goroutine
	defer func(){
		<- cfg.concurrencyControl //end goroutine
		cfg.wg.Done() //mark eork as complete
	}()
	if cfg.lenPages() >= cfg.maxSetting{ //check if crawling limit has been reached
		return 
	}
	normal_current, err := url.Parse(currentURL)
	if err != nil{
		fmt.Printf("unable to parse current url")
		return
	}

	if cfg.baseURL.Hostname() != normal_current.Hostname(){ //check correct domain
		return
	}

	normalize_current, err := normalizeURL(currentURL)
	if err != nil{
		fmt.Printf("Error - normalizedURL: %v", err)
		return 
	}

	/*if _, exists := cfg.pages[normalize_current]; exists{    debug message
		cfg.pages[normalize_current] += 1
		fmt.Printf("Already crawled: %s\n", normalize_current)
		return
	}else{
		cfg.pages[normalize_current] = 1
		fmt.Printf("Crawling: %s\n", normalize_current)   
	}*/
	
	isFirst := cfg.addPageVisit(normalize_current) //prevent duplicated crawling
	if !isFirst{
		return
	}
	fmt.Printf("Crawling: %s\n", currentURL)


	htmlBody, err := GetHTML(currentURL)
	if err != nil{
		fmt.Printf("Error - getHTML: %v", err)
		return
	}

	//fmt.Print(htmlBody)
	urls, err := getURLsFromHTML(htmlBody, cfg.baseURL) //recursively crawl extracted url
	if err != nil{
		fmt.Printf("error getting url from html")
		return 
	}
	for _, u := range urls{
		cfg.wg.Add(1) //add a new go routine to wait group
		go cfg.crawlPage(u) //recursively crawl links 
	}
}