package main

import "fmt"

func main() {
	//longestPalindrome2("aacabdkacaa")
	//radixSort([]int{802, 63, 2, 41, 12, 814, 53, 21, 532, 428, 436})

		root := BuildTreeFromList([]int{1, 2, 3, 4, -1, 5, 6, -1, -1, 7})
		res := findBottomLeftValue(root)
		expected := 7
		if res != expected {
			fmt.Errorf("expected: %d, real: %d\n", expected, res)
		}
}

