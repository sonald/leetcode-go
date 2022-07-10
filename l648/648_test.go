package l648

import (
	"fmt"
	"testing"
)

func TestReplaceWords(t *testing.T) {
	var testcases = []struct {
		dict     []string
		sen      string
		expected string
	}{
		{
			[]string{"catt", "bat", "cat", "rat"},
			"the cattle was rattled by the battery",
			"the cat was rat by the bat",
		},
		{
			[]string{"cat", "bat", "rat"},
			"the cattle was rattled by the battery",
			"the cat was rat by the bat",
		},
		{
			[]string{"a", "b", "c"},
			"aadsfasf absbs bbab cadsfafs",
			"a a b c",
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("case%d", i), func(t *testing.T) {
			pt := BuildPrefixTree(tc.dict)
			if p, b := pt.PrefixExists("cattle"); !b {
				t.Errorf("should be true")
			} else {
				if p == "" {
					t.Errorf("should be cat, but %s get", p)
				}
			}

			res := replaceWords(tc.dict, tc.sen)
			if res != tc.expected {
				t.Errorf("expected: '%v', real: '%v'\n", tc.expected, res)
			}
		})
	}
}
