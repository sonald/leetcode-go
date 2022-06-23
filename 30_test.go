package main

import (
	"fmt"
	"testing"
)

func TestSubstringWithConcatOfAllWords(t *testing.T) {
	var testcases = []struct {
		s        string
		words    []string
		expected []int
	}{
		{"barfoothefoobarman", []string{"foo", "bar"}, []int{0, 9}},
		{"wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}, nil},
		{"barfoofoobarthefoobarman", []string{"bar", "foo", "the"}, []int{6, 9, 12}},
		{"wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"}, []int{8}},
		{"a", []string{"a"}, []int{0}},
	}

	for i, v := range testcases {
		t.Run(fmt.Sprintf("case%d", i), func(t *testing.T) {
			res := findSubstringWithConcatOfAllWords(v.s, v.words)
			if !compareSlices(res, v.expected) {
				t.Errorf("expected: %+v, real: %+v\n", v.expected, res)
			}
		})
	}
}
