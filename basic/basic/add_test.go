package main

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		a, b, c int
	}{
		{1, 2, 3},
		{3, 5, 8},
		{2, 4, 6},
	}

	for _, tt := range tests {
		if actual := Add(tt.a, tt.b); actual != tt.c {
			fmt.Printf("%d + %d, except:%d, actual:%d", tt.a, tt.b, tt.c, actual)
		}
	}
}
