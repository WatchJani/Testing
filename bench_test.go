package main

import (
	"testing"
)

func BenchmarkFunc(b *testing.B) {

	list := make([]int, 10)

	len := 10

	for i := 0; i < b.N; i++ {
		for i := 0; i < len; i++ {
			list[i] = i
		}
	}
}
