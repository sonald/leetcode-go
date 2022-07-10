package l648

import (
	"strings"
)

// contains only lower-case letters
type PrefixNode struct {
	D    [26]*PrefixNode
	Tail bool
}
type PrefixTree struct {
	head PrefixNode
}

func BuildPrefixTree(dict []string) *PrefixTree {
	pt := &PrefixTree{}

	var dfs func(node *PrefixNode, w string, i int)
	dfs = func(node *PrefixNode, w string, i int) {
		if i == len(w) {
			node.Tail = true
			return
		}

		if node.D[w[i]-'a'] == nil {
			node.D[w[i]-'a'] = &PrefixNode{}
		}
		dfs(node.D[w[i]-'a'], w, i+1)
	}

	for _, w := range dict {
		dfs(&pt.head, w, 0)
	}

	return pt
}

//assuming ascii string
func (pt *PrefixTree) PrefixExists(s string) (string, bool) {
	np := &pt.head
	for i := 0; i < len(s); i++ {
		next := np.D[s[i]-'a']
		if next == nil {
			return "", false
		}
		if next.Tail {
			return s[:i+1], true
		}

		np = next
	}

	return "", false
}

func replaceWords(dictionary []string, sentence string) string {
	pt := BuildPrefixTree(dictionary)
	words := strings.Split(sentence, " ")
	var res = make([]string, 0, len(words))
	for _, w := range words {
		t := ""
		if prefix, b := pt.PrefixExists(w); b {
			t = prefix
		} else {
			t = w
		}

		res = append(res, t)
	}

	return strings.Join(res, " ")
}
