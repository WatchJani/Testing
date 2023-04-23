package main

import "testing"

func TestAdd(t *testing.T) {
	var sum int = Add(5, 5)

	if sum != 10 {
		t.Error("Error")
	}
}
