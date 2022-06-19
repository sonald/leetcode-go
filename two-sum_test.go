package main

import (
	"log"
	"testing"
)

func TestTwoSum(t *testing.T) {

	t.Run("one", func(t *testing.T) {
		nums := []int{2, 7, 11, 15}
		res := twoSum(nums, 9)
		if res == nil || res[0] != 0 || res[1] != 1 {
			t.Fatalf("expected: [0,1], real: %v", res)
		}
	})

	t.Run("two", func(t *testing.T) {
		nums := []int{3, 2, 4}
		res := twoSum(nums, 6)
		log.Printf("two: %v\n", res)

		if res == nil || res[0] != 1 || res[1] != 2 {
			t.Fatalf("expected: [1,2], real: %v", res)
		}
	})

	t.Run("three", func(t *testing.T) {
		nums := []int{3, 3}
		res := twoSum(nums, 6)
		if res == nil || res[0] != 0 || res[1] != 1 {
			t.Fatalf("expected: [0,1], real: %v", res)
		}
	})

}

func TestTwoSum2(t *testing.T) {

	t.Run("one", func(t *testing.T) {
		nums := []int{2, 7, 11, 15}
		res := twoSum2(nums, 9)
		if res == nil || res[0] != 0 || res[1] != 1 {
			t.Fatalf("expected: [0,1], real: %v", res)
		}
	})

	t.Run("two", func(t *testing.T) {
		nums := []int{3, 2, 4}
		res := twoSum2(nums, 6)

		if res == nil || res[0] != 1 || res[1] != 2 {
			t.Fatalf("expected: [1,2], real: %v", res)
		}
	})

	t.Run("three", func(t *testing.T) {
		nums := []int{3, 3}
		res := twoSum2(nums, 6)
		if res == nil || res[0] != 0 || res[1] != 1 {
			t.Fatalf("expected: [0,1], real: %v", res)
		}
	})

}
