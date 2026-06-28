package main

import (
	"fmt"
	"log"
	"os"
)

var tokens = make(chan struct{}, 10)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{}
	defer func() {
		<-tokens // Release a token
	}()

	list, err := Extract(url)

	if err != nil {
		log.Print(err)
	}
	return list
}

type Work struct {
	url   string
	depth int
}

func doCrawl(list []Work) {
	for _, work := range list {
		if seen[work.url] != true && work.depth <= maxDepth {
			seen[work.url] = true
			n++
			go func(link string, currentDepth int) {
				urls := crawl(link)
				var result []Work
				for _, url := range urls {
					work := Work{
						url:   url,
						depth: currentDepth + 1,
					}
					result = append(result, work)
				}
				worklist <- result
			}(work.url, work.depth)
		}
	}
}

var maxDepth = 3
var seen = make(map[string]bool)
var n int
var worklist = make(chan []Work)

func main() {
	n++
	go func() {
		var initWork []Work
		for _, url := range os.Args[1:] {
			work := Work{
				url:   url,
				depth: 0,
			}
			initWork = append(initWork, work)
		}
		worklist <- initWork
	}()

	for ; n > 0; n-- {
		list := <-worklist
		doCrawl(list)

	}
}
