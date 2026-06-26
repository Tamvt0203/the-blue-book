package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func walkDir(dir string, fileSize chan<- int64, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subDir := filepath.Join(dir, entry.Name())
			wg.Add(1)
			walkDir(subDir, fileSize, wg)
		} else {
			entryInfo, _ := entry.Info()
			fileSize <- entryInfo.Size()
		}
	}
}
func dirents(dir string) []os.DirEntry {
	smp <- struct{}{}
	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
	}
	defer func() {
		<-smp
	}()
	return entries
}

var verbose = flag.Bool("v", false, "show verbose progress message")
var smp = make(chan struct{}, 200)

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	fileSize := make(chan int64)
	var wg sync.WaitGroup
	for _, root := range roots {
		wg.Add(1)
		go walkDir(root, fileSize, &wg)
	}
	go func() {
		wg.Wait()
		close(fileSize)
	}()
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Microsecond)
	}
	var nfiles, nbytes int64
loop:
	for {
		select {
		case size, ok := <-fileSize:
			if !ok {
				break loop //fileSize was closed
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}
	printDiskUsage(nfiles, nbytes)

}
func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files %.3f MB\n", nfiles, float64(nbytes)/1e6)
}
