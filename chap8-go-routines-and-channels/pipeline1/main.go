package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for x := 0; x < 100; x += 1 {
			naturals <- x
		}
		close(naturals)
	}()

	go func() {
		for {
			x, ok := <-naturals
			if ok {
				squares <- x * x
			} else {
				close(squares)
				break
			}

		}
	}()

	for {
		x, ok := <-squares
		if ok {
			fmt.Println(x)

		} else {
			break
		}
	}
}
