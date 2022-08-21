package main

import "testing"

func TestIsPrefixWord(t *testing.T) {
	var testcases = []struct {
		sentence   string
		searchWord string
		expected   int
	}{
		{"i love eating burger", "burg", 4},
		{"this problem is an easy problem", "pro", 2},
		{"i am tired", "you", -1},
	}

	for _, tc := range testcases {
		res := isPrefixWord2(tc.sentence, tc.searchWord)
		if res != tc.expected {
			t.Errorf("search %s, expected: %d, real: %d\n", tc.searchWord, tc.expected, res)
		}
	}
}
