package main
import(
	"fmt"
	"net/url"
	"sync"
)


type config struct{
	pages map[string]int
	baseURL *url.URL
	mu *sync.Mutex
	concurrencyControl chan struct{}
	wg *sync.WaitGroup
	maxSetting int
}


func (cfg *config) addPageVisit(normalURL string)(isFirst bool){
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	if _,Visited := cfg.pages[normalURL]; Visited{
		cfg.pages[normalURL] += 1
		return false
	}
	cfg.pages[normalURL] = 1
	return true
}


func (cfg *config) lenPages()(int){
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	return len(cfg.pages)
}


func configure(rawBASEURL string, maxConcurrency int, maxsetting int)(*config, error){
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