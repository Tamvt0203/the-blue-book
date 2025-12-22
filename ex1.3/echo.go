package main

import (
	"strings" 
)

func poor(args []string) string {
	var s, sep string
	for i := 1; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	return s
}

func elegant(args []string) string {
	return strings.Join(args[1:], " ")
}
