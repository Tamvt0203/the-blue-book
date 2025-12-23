package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Printf("error while reading file: %s", file)
				continue
			}
			counts = make(map[string]int)
			countLines(f, counts)
			for _, n := range counts {
				if n > 1 {
					fmt.Printf("File %s contains duplicates!\n", file)
					break
				}
			}
			defer f.Close()
		}
	}

}
func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
