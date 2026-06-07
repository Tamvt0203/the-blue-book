package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Tamvt0203/the-blue-book/chap5-functions/links"
)

var tokens = make(chan struct{}, 20)

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

func main() {
	worklist := make(chan []string)
	var n int
	n++
	go func() {
		worklist <- os.Args[1:]
	}()

	seen := make(map[string]bool)
	// use n to keep track on the number of work to do
	for depth := 0; depth < 3; depth++ {
		list := <-worklist
		fmt.Println(list)
		for _, link := range list {
			if seen[link] != true {
				seen[link] = true
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}
