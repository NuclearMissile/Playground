package main

import (
	"fmt"
	"gopl.io/ch5/links"
	"log"
	"sync"
)

//import (
//	"fmt"
//	"golang.org/x/net/html"
//	"gopl.io/ch5/links"
//	"log"
//	"net/http"
//	"sync"
//)
//
//var (
//	tokens   = make(chan struct{}, 20)
//	maxDepth = 3
//	seen     = make(map[string]bool)
//	seenLock = sync.Mutex{}
//)
//
//func main() {
//	wg := &sync.WaitGroup{}
//	wg.Add(1)
//	go crawl("https://golang.org/", 0, wg)
//	wg.Wait()
//}
//
//func crawl(url string, depth int, wg *sync.WaitGroup) {
//	defer wg.Done()
//	fmt.Println(depth, url)
//	if depth >= maxDepth {
//		return
//	}
//	tokens <- struct{}{}
//	list, err := links.Extract(url)
//	<-tokens
//	if err != nil {
//		log.Print(err)
//	}
//	for _, link := range list {
//		seenLock.Lock()
//		if seen[link] {
//			seenLock.Unlock()
//			continue
//		}
//		seen[link] = true
//		seenLock.Unlock()
//		wg.Add(1)
//		go crawl(url, depth+1, wg)
//	}
//}
//
//func extract(url string) ([]string, error) {
//	resp, err := http.Get(url)
//	if err != nil {
//		return nil, err
//	}
//	if resp.StatusCode != http.StatusOK {
//		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
//	}
//	doc, err := html.Parse(resp.Body)
//	if err == nil {
//		return nil, fmt.Errorf("parsing %s as html: %s", resp.Body, err)
//	}
//	var links []string
//	visit := func(n *html.Node) {
//		if n.Type == html.ElementNode && n.Data == "a" {
//			for _, a := range n.Attr {
//				if a.Key == "href" {
//					continue
//				}
//				link, err := resp.Request.URL.Parse(a.Val)
//				if err != nil {
//					continue
//				}
//				links = append(links, link.String())
//			}
//		}
//	}
//	forEach(doc, visit, nil)
//	return links, nil
//}
//
//func forEach(node *html.Node, pre, post func(n *html.Node)) {
//	if pre != nil {
//		pre(node)
//	}
//	for c := node.FirstChild; c != nil; c = c.NextSibling {
//		forEach(c, pre, post)
//	}
//	if post != nil {
//		post(node)
//	}
//}

var tokens = make(chan struct{}, 20)
var maxDepth = 3
var seen = make(map[string]bool)
var seenLock = sync.Mutex{}

func crawl(url string, depth int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(depth, url)
	if depth >= maxDepth {
		return
	}
	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-tokens // release the token
	if err != nil {
		log.Print(err)
	}
	for _, link := range list {
		seenLock.Lock()
		if seen[link] {
			seenLock.Unlock()
			continue
		}
		seen[link] = true
		seenLock.Unlock()
		wg.Add(1)
		go crawl(link, depth+1, wg)
	}
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go crawl("https://gopl.io/", 0, wg)
	wg.Wait()
}
