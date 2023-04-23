package main

import (
	"log"
	"testing"
)

type AddData struct {
	x, y   int
	result int
}

func TestAdd(t *testing.T) {
	for _, data := range []AddData{
		{1, 2, 3},
		{6, 5, 11},
		{182, 458, 640},
	} {
		if Add(data.x, data.y) != data.result {
			t.Error("Error")
		} else {
			log.Printf("true data testing: %d", data.result)
		}
	}
}
