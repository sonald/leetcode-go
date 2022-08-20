package main

import "testing"

func TestStringMatching(t *testing.T) {
	var testcases = []struct {
		words    []string
		expected []string
	}{
		{[]string{"mass", "as", "hero", "superhero"}, []string{"as", "hero"}},
        {[]string{"leetcoder","leetcode","od","hamlet","am"}, []string{"leetcode", "od", "am"}},
	}

    for _, tc := range testcases {
        res := stringMatching(tc.words)
        if !compareStringSlices(res, tc.expected) {
            t.Errorf("words: %v, expected: %v, real: %v\n", tc.words, tc.expected, res)
        }
    }
}
