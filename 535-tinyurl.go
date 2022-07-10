package main

import (
	"fmt"
	"math/rand"
)

type Codec struct {
	db map[string]string
}

func Constructor() Codec {
	return Codec{
		db: make(map[string]string),
	}
}

const (
	candidates = "abcdefghijklmnopqrstuvwxyz1234567890"
	TINY_SIZE  = 24
	PREFIX     = "http://sonald.me/"
)

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	for {
		var tiny [TINY_SIZE]byte
		for i := 0; i < TINY_SIZE; i++ {
			k := rand.Intn(len(candidates))
			tiny[i] = candidates[k]
		}
		if _, ok := this.db[string(tiny[:])]; !ok {
			res := fmt.Sprintf("%s%s", PREFIX, string(tiny[:]))
			this.db[res] = longUrl
			return res
		}
	}
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	tiny := shortUrl[len(PREFIX):]
	if res, ok := this.db[tiny]; ok {
		return res
	}

	return shortUrl
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * url := obj.encode(longUrl);
 * ans := obj.decode(url);
 */
