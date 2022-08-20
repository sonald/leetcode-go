package main

import (
	"strings"
)

func stringMatching(words []string) []string {
    var res []string

    for i, s := range words {
        for j, t := range words {
            if i != j && strings.Contains(t, s) {
                res = append(res, s)
                break
            }
        }
    }

	return res
}
