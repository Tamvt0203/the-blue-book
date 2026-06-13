package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Tamvt0203/the-blue-book/chap5-functions/links"
)

var tokens = make(chan struct{}, 10)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // Aqurire a token
	list, err := links.Extract(url)
	<-tokens // Release a token
	if err != nil {
		log.Print(err)
	}
	return list
}

type Work struct {
	url   string
	depth int
}

func main() {
	maxDepth := 3
	worklist := make(chan []Work)
	var n int
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

	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
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
}
