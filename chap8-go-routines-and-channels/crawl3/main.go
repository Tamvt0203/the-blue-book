package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Tamvt0203/the-blue-book/chap5-functions/links"
)

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	worklist := make(chan []string)
	unseenLinks := make(chan string)
	go func() {
		worklist <- os.Args[1:]
	}()
	for n := 0; n < 20; n++ {
		go func() {
			for link := range unseenLinks {
				found := crawl(link)
				go func() {
					worklist <- found
				}()
			}
		}()
	}
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if seen[link] != true {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}
