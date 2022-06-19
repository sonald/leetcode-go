package main

import (
	"fmt"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	var data = [][]string{
		{"cbbd", "bb"},
		{"babad", "aba"},
		{"aaaaa", "aaaaa"},
		{"aacabdkacaa", "aca"},
	}

	for i, p := range data {
		t.Run(fmt.Sprintf("test%d", i), func(t *testing.T) {
			res := longestPalindrome(p[0])
			if res != p[1] {
				t.Fatalf("expect %s, real %s\n", p[1], res)
			}
		})
	}
}

func TestLongestPalindrome2(t *testing.T) {
	var data = [][]string{
		{"cbbd", "bb"},
		{"bb", "bb"},
		{"babad", "bab"},
		{"aaaaa", "aaaaa"},
		{"aacabdkacaa", "aca"},
	}

	for i, p := range data {
		t.Run(fmt.Sprintf("test%d", i), func(t *testing.T) {
			res := longestPalindrome2(p[0])
			if res != p[1] {
				t.Fatalf("expect %s, real %s\n", p[1], res)
			}
		})
	}
}
