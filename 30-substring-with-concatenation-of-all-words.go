package main

func findSubstringWithConcatOfAllWords(s string, words []string) []int {
	if len(words) == 0 || len(s) == 0 {
		return nil
	}

	var wordmap = map[string]int{}
	for _, w := range words {
		wordmap[w]++
	}

	var res []int
	var K, M, N = len(words[0]), len(words), len(s)

	for i := 0; i <= N-K*M; i++ {
		cands := map[string]int{}
		for j := i; j <= N-K; j += K {
			t := s[j : j+K]
			if _, ok := wordmap[t]; !ok {
				break
			}

			if v, ok := cands[t]; ok && v >= wordmap[t] {
				//already exists, redo
				break
			}

			cands[t]++
		}

		sz := 0
		for _, v := range cands {
			sz += v
		}
		if sz == M {
			res = append(res, i)
		}
	}

	return res
}
