package main
import(
	"fmt"
	"net/url"
	"sync"
)


type config struct{
	pages map[string]int //tracks visit
	baseURL *url.URL //url of the site being crawled
	mu *sync.Mutex //ensure thread safe access
	concurrencyControl chan struct{} //buffered channel to control concurrent crawling
	wg *sync.WaitGroup //use to sync goroutines
	maxSetting int //controls how many goroutines
}


func (cfg *config) addPageVisit(normalURL string)(isFirst bool){ //check whether site has been visited 
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	if _,Visited := cfg.pages[normalURL]; Visited{ //if url in cfg.pages, increment by1, else add to cfg.pages
		cfg.pages[normalURL] += 1
		return false
	}
	cfg.pages[normalURL] = 1
	return true
}


func (cfg *config) lenPages()(int){ //return len of cfg.pages
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	return len(cfg.pages)
}


func configure(rawBASEURL string, maxConcurrency int, maxsetting int)(*config, error){ //make a new config struct 
	baseurl, err := url.Parse(rawBASEURL)
	if err != nil{
		return nil, fmt.Errorf("couldn't parse base URL: %v", err)
	}

	return &config{
		pages:    make(map[string]int),
		baseURL: baseurl,
		mu: &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg: &sync.WaitGroup{},
		maxSetting: maxsetting,
	}, nil
}