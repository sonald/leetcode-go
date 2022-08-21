package main

import "bytes"

func isPrefixWord(sentence string, searchWord string) int {
    s, w := []byte(sentence), []byte(searchWord)
    id := 1
    for i := bytes.IndexByte(s, ' '); i != -1; i = bytes.IndexByte(s, ' ') {
        if bytes.HasPrefix(s, w) {
            return id
        }

        id++
        s = s[i+1:]
    }

    if len(s) > 0 && bytes.HasPrefix(s, w) {
        return id
    }
	return -1
}

func isPrefixWord2(sentence string, searchWord string) int {
    id, p := 1, 0
    miss := false
    s, w := []byte(sentence), []byte(searchWord)
    for i := 0; i < len(s); i++ {
        if s[i] == ' ' {
            if !miss && p == len(w) {
                return id
            }
            miss = false
            p = 0
            id++
            continue
        } else if miss {
            continue
        }

        if p == len(w) {
            return id
        }

        if w[p] != s[i] {
            miss = true
        }
        p++
    }

    if !miss && p == len(w) {
        return id
    }
    return -1
}
