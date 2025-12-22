package main

import (
	"testing"
)

func BenchmarkPoor(b *testing.B) {
	args := []string{"a", "b", "c", "d"}
	for i := 0; i < b.N; i++ {
		poor(args)
	}
}

func BenchmarkElegant(b *testing.B) {
	args := []string{"a", "b", "c", "d"}
	for i := 0; i < b.N; i++ {
		elegant(args)
	}
}
