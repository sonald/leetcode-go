package main

import (
	"testing"
)

func TestMaxArea(t *testing.T) {
	t.Run("one", func(t *testing.T) {
        var res = maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
        if res != 49 {
            t.Fatalf("expect %d, real %d\n", 49, res)
        }
	})

	t.Run("two", func(t *testing.T) {
        var res = maxArea([]int{1, 1})
        if res != 1 {
            t.Fatalf("expect %d, real %d\n", 1, res)
        }
	})
}

