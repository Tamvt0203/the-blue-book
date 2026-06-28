package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type Root struct {
	name   string
	nfiles int64
	nbytes int64
}

type FileSize struct {
	rootIndex int
	fileSize  int64
}

func walkDir(dir string, fileSize chan<- FileSize, wg *sync.WaitGroup, rootIndex int) {
	defer wg.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subDir := filepath.Join(dir, entry.Name())
			wg.Add(1)
			go walkDir(subDir, fileSize, wg, rootIndex)
		} else {
			entryInfo, _ := entry.Info()
			rootFileSize := FileSize{
				rootIndex: rootIndex,
				fileSize:  entryInfo.Size(),
			}
			fileSize <- rootFileSize
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
	rootsInfo := make([]Root, 0)

	fileSize := make(chan FileSize)
	var wg sync.WaitGroup
	for index, root := range roots {
		rootInfo := Root{
			name:   root,
			nfiles: 0,
			nbytes: 0,
		}
		rootsInfo = append(rootsInfo, rootInfo)
		wg.Add(1)
		go walkDir(root, fileSize, &wg, index)
	}
	go func() {
		wg.Wait()
		close(fileSize)
	}()
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(1000 * time.Microsecond)
	}
loop:
	for {
		select {
		case fileSize, ok := <-fileSize:
			if !ok {
				break loop //fileSize was closed
			}
			rootsInfo[fileSize.rootIndex].nfiles++
			rootsInfo[fileSize.rootIndex].nbytes += fileSize.fileSize
		case <-tick:
			printDiskUsage(rootsInfo)
		}
	}
	printDiskUsage(rootsInfo)

}
func printDiskUsage(rootsInfor []Root) {
	for _, rootInfo := range rootsInfor {
		fmt.Printf("Dir %s: %d files %.3f MB\n", rootInfo.name, rootInfo.nfiles, float64(rootInfo.nbytes)/1e6)

	}
}
