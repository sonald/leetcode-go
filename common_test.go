package main

import "testing"

func TestAVL(t *testing.T) {
	avl := AVLInit()
	var data = []int{3, 1, 10, 4, 8, 5, 6, 11, 9}
	for _, v := range data {
		avl.Insert(v)
		avl.Dump()
	}

	if avl.head.h != 4 {
		t.Errorf("head height should be 4")
	}
}


