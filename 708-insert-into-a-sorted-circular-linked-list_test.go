package main

import (
	"testing"
)

func TestInsert(t *testing.T) {
	t.Run("zero", func(t *testing.T) {
		np := insert(nil, 2)
		if np == nil || np.Val != 2 {
			t.Fatalf("return wrong node: %v\n", np)
		}

		if np.Next != np {
			t.Fatalf("return wrong node: %v\n", np.Next)
		}
	})

	t.Run("one", func(t *testing.T) {
		nd3 := &Node{Val: 4}
		nd2 := &Node{Val: 3, Next: nd3}
		nd1 := &Node{Val: 1, Next: nd2}
		nd3.Next = nd1

		np := insert(nd3, 2)
		if np == nil || np.Val != 4 {
			t.Fatalf("return wrong node: %v\n", np)
		}

		if np.Next.Val != nd1.Val {
			t.Fatalf("return wrong node: %v\n", np.Next)
		}

		if np.Next.Next.Val != 2 {
			t.Fatalf("return wrong node: %v\n", np)
		}

		if np.Next.Next.Next.Val != nd2.Val {
			t.Fatalf("return wrong node: %v\n", np)
		}
	})

	t.Run("two", func(t *testing.T) {
		nd3 := &Node{Val: 4}
		nd2 := &Node{Val: 3, Next: nd3}
		nd1 := &Node{Val: 2, Next: nd2}
		nd3.Next = nd1

		np := insert(nd1, 1)
		if np == nil || np.Val != 2 {
			t.Fatalf("return wrong node: %v\n", np)
		}

		if np.Next.Val != nd2.Val {
			t.Fatalf("return wrong node: %v\n", np)
		}

		if np.Next.Next.Val != nd3.Val {
			t.Fatalf("return wrong node: %v\n", np)
		}

		if np.Next.Next.Next.Val != 1 {
			t.Fatalf("return wrong node: %v\n", np)
		}
	})

	t.Run("three", func(t *testing.T) {
		nd3 := &Node{Val: 3}
		nd2 := &Node{Val: 3, Next: nd3}
		nd1 := &Node{Val: 3, Next: nd2}
		nd3.Next = nd1

		np := insert(nd2, 0)
		if np == nil || np.Val != 3 {
			t.Fatalf("return wrong node: %v\n", np)
		}

		if np.Next.Val != 3 && np.Next.Val != 0 {
			t.Fatalf("return wrong node: %v\n", np.Next)
		}

		if np.Next.Next.Val != 0 && np.Next.Next.Val != 3 {
			t.Fatalf("return wrong node: %v\n", np)
		}

		if np.Next.Next.Next.Val != 0 && np.Next.Next.Next.Val != 3 {
			t.Fatalf("return wrong node: %v\n", np)
		}
	})

	t.Run("four", func(t *testing.T) {
		nd3 := &Node{Val: 5}
		nd2 := &Node{Val: 3, Next: nd3}
		nd1 := &Node{Val: 3, Next: nd2}
		nd3.Next = nd1

		np := insert(nd1, 0)
		if np == nil || np.Val != 3 {
			t.Fatalf("return wrong node: %v\n", np)
		}

		if np.Next.Val != 3 {
			t.Fatalf("return wrong node: %v\n", np.Next)
		}

		if np.Next.Next.Val != 5 {
			t.Fatalf("return wrong node: %v\n", np)
		}

		if np.Next.Next.Next.Val != 0 {
			t.Fatalf("return wrong node: %v\n", np)
		}
	})

	t.Run("five", func(t *testing.T) {
		nd3 := &Node{Val: 3}
		nd2 := &Node{Val: 3, Next: nd3}
		nd1 := &Node{Val: 1, Next: nd2}
		nd3.Next = nd1

		np := insert(nd1, 4)
		if np == nil || np.Val != 1 {
			t.Fatalf("return wrong node: %v\n", np)
		}

		if np.Next.Val != 3 {
			t.Fatalf("return wrong node: %v\n", np.Next)
		}

		if np.Next.Next.Val != 3 {
			t.Fatalf("return wrong node: %v\n", np.Next.Next)
		}

		if np.Next.Next.Next.Val != 4 {
			t.Fatalf("return wrong node: %v\n", np.Next.Next.Next)
		}
	})
}
