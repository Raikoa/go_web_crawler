package main
import (
	"net/url"
	"fmt"
)




func (cfg *config) crawlPage(currentURL string){
	
	cfg.concurrencyControl <- struct{}{}
	defer func(){
		<- cfg.concurrencyControl
		cfg.wg.Done()
	}()
	if cfg.lenPages() >= cfg.maxSetting{
		return 
	}
	normal_current, err := url.Parse(currentURL)
	if err != nil{
		fmt.Printf("unable to parse current url")
		return
	}

	if cfg.baseURL.Hostname() != normal_current.Hostname(){
		return
	}

	normalize_current, err := normalizeURL(currentURL)
	if err != nil{
		fmt.Printf("Error - normalizedURL: %v", err)
		return 
	}

	/*if _, exists := cfg.pages[normalize_current]; exists{
		cfg.pages[normalize_current] += 1
		fmt.Printf("Already crawled: %s\n", normalize_current)
		return
	}else{
		cfg.pages[normalize_current] = 1
		fmt.Printf("Crawling: %s\n", normalize_current)
	}*/
	
	isFirst := cfg.addPageVisit(normalize_current)
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
	urls, err := getURLsFromHTML(htmlBody, cfg.baseURL)
	if err != nil{
		fmt.Printf("error getting url from html")
		return 
	}
	for _, u := range urls{
		cfg.wg.Add(1)
		go cfg.crawlPage(u)
	}
}