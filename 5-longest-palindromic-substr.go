package main

// time: O(N*N),space: O(N)
func longestPalindrome(s string) string {
	var n = len(s)
	var m = make([][]bool, 3)
	for i := 0; i < 3; i++ {
		m[i] = make([]bool, n)
	}

	res := ""
	for i := 0; i < n; i++ {
		m[0][i] = true
		res = s[i : i+1]
	}

	for i := 0; i < n-1; i++ {
		if s[i] == s[i+1] {
			m[1][i] = true
			res = s[i : i+2]
		}
	}

	for l := 2; l < n; l++ {
		for i := 0; i+l < n; i++ {
			m[2][i] = false
			if m[0][i+1] && s[i] == s[i+l] {
				m[2][i] = true
				res = s[i : i+l+1]
			}
		}

		copy(m[0], m[1])
		copy(m[1], m[2])
	}

	return res
}

func longestPalindrome2(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}

	expand := func(c1, c2 int) (int, int) {
		for c1 >= 0 && c2 < n && s[c1] == s[c2] {
			c1--
			c2++
		}
		return c1 + 1, c2 - 1
	}

	var start, end = 0, 0
	for c := 0; c < n-1; c++ {
		if l, r := expand(c, c); end-start < r-l {
			start, end = l, r
		}
		if l, r := expand(c, c+1); end-start < r-l {
			start, end = l, r
		}
	}

	return s[start : end+1]
}

// https://leetcode.cn/problems/longest-palindromic-substring/solution/zui-chang-hui-wen-zi-chuan-by-leetcode-solution/
// manacher algorithm
