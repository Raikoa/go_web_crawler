package main
import "net/url"
import "golang.org/x/net/html"
import "strings"
import "fmt"



func normalizeURL(Url string) (string, error){
	URL, err := url.Parse(Url)
	if err != nil{
		return "", err
	}
	return URL.Host + URL.Path, nil


}


func getURLsFromHTML(htmlBody string, BaseURL *url.URL) ([]string, error) {
	
	htmlReader := strings.NewReader(htmlBody)
	doc, err := html.Parse(htmlReader)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse HTML: %v", err)
	}

	var urls []string
	var traverseNodes func(*html.Node)
	traverseNodes = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "a" {
			for _, anchor := range node.Attr {
				if anchor.Key == "href" {
					href, err := url.Parse(anchor.Val)
					if err != nil {
						fmt.Printf("couldn't parse href '%v': %v\n", anchor.Val, err)
						continue
					}

					resolvedURL := BaseURL.ResolveReference(href)
					urls = append(urls, resolvedURL.String())
				}
			}
		}

		for child := node.FirstChild; child != nil; child = child.NextSibling {
			traverseNodes(child)
		}
	}
	traverseNodes(doc)

	return urls, nil
}




