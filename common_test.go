package main

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

func testAVL(t *testing.T, data []int) {
	avl := AVLInit()
	for _, v := range data {
		avl.Insert(v)
		avl.Dump()
	}

	var N = len(data)
	var sorted = make([]int, N)
	copy(sorted, data)
	sort.Ints(sorted)

	i := 0
	for {
		if avl.getMin() != sorted[i] {
			t.Errorf("expected: %d, real: %d\n", sorted[i], avl.head.Data)
		}
		avl.DeleteMin()
		if avl.head != nil {
			break
		}
		i++
	}
}

func TestAVL(t *testing.T) {
	var data = []int{3, 1, 10, 4, 8, 5, 6, 11, 9}
	testAVL(t, data)
}

func FuzzAVL(f *testing.F) {
	f.Add(100, 98, 237, 432, 120, 56, 99, 323, 402)
	f.Fuzz(func(t *testing.T, d1, d2, d3, d4, d5, d6, d7, d8, d9 int) {
		testAVL(t, []int{d1, d2, d3, d4, d5, d6, d7, d8, d9})
	})
}

func TestAVL_Search(t *testing.T) {
	var data = []int{3, 1, 10, 4, 8, 5, 6, 11, 9}
	avl := AVLInit()
	for _, v := range data {
		avl.Insert(v)
	}
	avl.Dump()

	var testcases = []struct {
		bound    int
		expected int
	}{
		{2, 3},
		{7, 8},
		{1, 1},
		{0, 1},
		{12, -1},
		{-12, 1},
		{5, 5},
		{4, 4},
	}

	for _, p := range testcases {
		np := avl.Search(func(v int) bool {
			return v >= p.bound
		})

		if p.expected < 0 {
			if np != nil {
				t.Errorf("expected: nil, real: %+v\n", np)
			}
		} else {
			if np == nil || np.Data != p.expected {
				t.Errorf("bound: %d, expected: %d, real: %+v\n", p.bound, p.expected, np)
			}
		}
	}
}

func BenchmarkAVL_Insert(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	avl := AVLInit()
	for i := 0; i < b.N; i++ {
		avl.Insert(rand.Intn(1000_000))
	}

	fmt.Printf("size: %d, height: %d\n", avl.Size(), avl.head.h)
}

func BenchmarkAVL_Remove(b *testing.B) {
	avl := AVLInit()
	for i := 0; i < b.N; i++ {
		avl.Insert(i)
	}
	fmt.Printf("size: %d, height: %d\n", avl.Size(), avl.head.h)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		avl.Remove(rand.Intn(b.N))
	}

	var h = 0
	if avl.head != nil {
		h = avl.head.h
	}
	fmt.Printf("end: size: %d, height: %d\n", avl.Size(), h)
}

func BenchmarkAVL_DeleteMin(b *testing.B) {
	avl := AVLInit()
	for i := 0; i < b.N; i++ {
		avl.Insert(rand.Intn(1000_000))
	}
	fmt.Printf("size: %d, height: %d\n", avl.Size(), avl.head.h)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		avl.DeleteMin()
	}
}

func TestHeapSort(t *testing.T) {
	var testcases = []struct {
		nums     []int
		expected []int
	}{
		{[]int{7, 1, 4, 2, 5, 3, 6, 9, 3}, []int{1, 2, 3, 3, 4, 5, 6, 7, 9}},
	}

	for i, a := range testcases {
		t.Run(fmt.Sprintf("case%d", i), func(t *testing.T) {
			HeapSort(a.nums)
			if !compareSlicesNoSort(a.nums, a.expected) {
				t.Errorf("expected: %v, real: %v\n", a.expected, a.nums)
			}
		})
	}
}
