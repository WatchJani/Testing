package main

import (
	"fmt"
	"testing"
)

func BenchmarkFunc(b *testing.B) {

	list := [...]int{1, 2, 3, 456, 465, 65, 45, 645, 465, 4665, 465, 4}

	len := 12

	for i := 0; i < b.N; i++ {
		for index := 0; index < len; index++ {
			fmt.Println(list[index])
		}
	}
}
