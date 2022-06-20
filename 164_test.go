package main

import (
	"fmt"
	"log"
	"sort"
	"testing"
)

func FuzzRadixSort(f *testing.F) {
	f.Add(802, 35, 111, 441, 3412, 0, 43837, 2024)

	f.Fuzz(func(t *testing.T, d1, d2, d3, d4, d5, d6, d7, d8 int) {
		nums := []int{d1, d2, d3, d4, d5, d6, d7, d8}
		for i := 0; i < len(nums); i++ {
			if nums[i] < 0 {
				nums[i] = -nums[i]
			}
		}

		var expected = make([]int, len(nums))
		copy(expected, nums)
		sort.Ints(expected)
		radixSort(nums)
		log.Printf("input: %v\n", nums)
		log.Printf("expected: %v\n", expected)
		if !compareSortedSlices(nums, expected) {
			t.Errorf("sort failed")
		}
	})
}

func TestMaximumGap(t *testing.T) {
	var data = []struct {
		a []int
		b int
	}{
		{[]int{3, 6, 9, 1}, 3},
		{[]int{10}, 0},
	}

	for i, p := range data {
		t.Run(fmt.Sprintf("test%d", i), func(t *testing.T) {
			res := maximumGap(p.a)
			if res != p.b {
				t.Fatalf("expected: %d, real: %d\n", p.b, res)
			}
		})
	}
}
func TestMaximumGap2(t *testing.T) {
	var data = []struct {
		a []int
		b int
	}{
		{[]int{3, 6, 9, 1}, 3},
		{[]int{10}, 0},
	}

	for i, p := range data {
		t.Run(fmt.Sprintf("test%d", i), func(t *testing.T) {
			res := maximumGap2(p.a)
			if res != p.b {
				t.Fatalf("expected: %d, real: %d\n", p.b, res)
			}
		})
	}
}
