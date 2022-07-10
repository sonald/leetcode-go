package l729

import (
	"fmt"
	"testing"
)

func TestMyCalendar_Book(t *testing.T) {
	t.Run("round1", func(t *testing.T) {

		c := Constructor()
		var testcases = []struct {
			interval Interval729
			result   bool
		}{
			{Interval729{10, 20}, true},
			{Interval729{15, 25}, false},
			{Interval729{20, 30}, true},
			{Interval729{40, 50}, true},
			{Interval729{30, 42}, false},
		}

		for i, tc := range testcases {
			t.Run(fmt.Sprintf("test%d", i), func(t *testing.T) {
				res := c.Book(tc.interval.start, tc.interval.end)
				if res != tc.result {
					t.Errorf("expected: %v, real: %v\n", tc.result, res)
				}
			})
		}
	})

	t.Run("round2", func(t *testing.T) {
		c := Constructor()
		var testcases = []struct {
			interval Interval729
			result   bool
		}{
			{Interval729{37, 50}, true},
			{Interval729{33, 50}, false},
			{Interval729{4, 17}, true},
			{Interval729{35, 48}, false},
			{Interval729{8, 25}, false},
		}

		for i, tc := range testcases {
			t.Run(fmt.Sprintf("test%d", i), func(t *testing.T) {
				res := c.Book(tc.interval.start, tc.interval.end)
				if res != tc.result {
					t.Errorf("expected: %v, real: %v\n", tc.result, res)
				}
			})
		}
	})

	t.Run("round3", func(t *testing.T) {
		c := Constructor()
		var testcases = []struct {
			interval Interval729
			result   bool
		}{
			{Interval729{47, 50}, true},
			{Interval729{33, 41}, true},
			{Interval729{39, 45}, false},
			{Interval729{33, 42}, false},
			{Interval729{25, 32}, true},
			{Interval729{26, 35}, false},
			{Interval729{19, 25}, true},
			{Interval729{3, 8}, true},
			{Interval729{8, 13}, true},
			{Interval729{18, 27}, false},
		}

		for i, tc := range testcases {
			t.Run(fmt.Sprintf("test%d", i), func(t *testing.T) {
				res := c.Book(tc.interval.start, tc.interval.end)
				if res != tc.result {
					t.Errorf("expected: %v, real: %v\n", tc.result, res)
				}
			})
		}
	})
}

func TestMyCalendar_Book2(t *testing.T) {
	t.Run("round1", func(t *testing.T) {
		c := Constructor()
		var testcases = []struct {
			interval Interval729
			result   bool
		}{
			{Interval729{10, 20}, true},
			{Interval729{15, 25}, false},
			{Interval729{20, 30}, true},
			{Interval729{40, 50}, true},
			{Interval729{30, 42}, false},
		}

		for i, tc := range testcases {
			t.Run(fmt.Sprintf("test%d", i), func(t *testing.T) {
				res := c.Book2(tc.interval.start, tc.interval.end)
				if res != tc.result {
					t.Errorf("it: %v, expected: %v, real: %v\n", tc.interval, tc.result, res)
				}
			})
		}
	})

	t.Run("round2", func(t *testing.T) {
		c := Constructor()
		var testcases = []struct {
			interval Interval729
			result   bool
		}{
			{Interval729{37, 50}, true},
			{Interval729{33, 50}, false},
			{Interval729{4, 17}, true},
			{Interval729{35, 48}, false},
			{Interval729{8, 25}, false},
		}

		for i, tc := range testcases {
			t.Run(fmt.Sprintf("test%d", i), func(t *testing.T) {
				res := c.Book2(tc.interval.start, tc.interval.end)
				if res != tc.result {
					t.Errorf("it: %v, expected: %v, real: %v\n", tc.interval, tc.result, res)
				}
			})
		}
	})

	t.Run("round3", func(t *testing.T) {
		c := Constructor()
		var testcases = []struct {
			interval Interval729
			result   bool
		}{
			{Interval729{47, 50}, true},
			{Interval729{33, 41}, true},
			{Interval729{39, 45}, false},
			{Interval729{33, 42}, false},
			{Interval729{25, 32}, true},
			{Interval729{26, 35}, false},
			{Interval729{19, 25}, true},
			{Interval729{3, 8}, true},
			{Interval729{8, 13}, true},
			{Interval729{18, 27}, false},
		}

		for i, tc := range testcases {
			t.Run(fmt.Sprintf("test%d", i), func(t *testing.T) {
				c.bst.Dump()
				res := c.Book2(tc.interval.start, tc.interval.end)
				if res != tc.result {
					t.Errorf("it: %v, expected: %v, real: %v\n", tc.interval, tc.result, res)
				}
			})
		}
	})
}

func TestMyCalendar_Book3(t *testing.T) {
	t.Run("round1", func(t *testing.T) {
		c := Constructor()
		var testcases = []struct {
			interval Interval729
			result   bool
		}{
			{Interval729{10, 20}, true},
			{Interval729{15, 25}, false},
			{Interval729{20, 30}, true},
			{Interval729{40, 50}, true},
			{Interval729{30, 42}, false},
		}

		for i, tc := range testcases {
			t.Run(fmt.Sprintf("test%d", i), func(t *testing.T) {
				res := c.Book3(tc.interval.start, tc.interval.end)
				if res != tc.result {
					t.Errorf("it: %v, expected: %v, real: %v\n", tc.interval, tc.result, res)
				}
			})
		}
	})

	t.Run("round2", func(t *testing.T) {
		c := Constructor()
		var testcases = []struct {
			interval Interval729
			result   bool
		}{
			{Interval729{37, 50}, true},
			{Interval729{33, 50}, false},
			{Interval729{4, 17}, true},
			{Interval729{35, 48}, false},
			{Interval729{8, 25}, false},
		}

		for i, tc := range testcases {
			t.Run(fmt.Sprintf("test%d", i), func(t *testing.T) {
				res := c.Book3(tc.interval.start, tc.interval.end)
				if res != tc.result {
					t.Errorf("it: %v, expected: %v, real: %v\n", tc.interval, tc.result, res)
				}
			})
		}
	})

	t.Run("round3", func(t *testing.T) {
		c := Constructor()
		var testcases = []struct {
			interval Interval729
			result   bool
		}{
			{Interval729{47, 50}, true},
			{Interval729{33, 41}, true},
			{Interval729{39, 45}, false},
			{Interval729{33, 42}, false},
			{Interval729{25, 32}, true},
			{Interval729{26, 35}, false},
			{Interval729{19, 25}, true},
			{Interval729{3, 8}, true},
			{Interval729{8, 13}, true},
			{Interval729{18, 27}, false},
		}

		for i, tc := range testcases {
			t.Run(fmt.Sprintf("test%d", i), func(t *testing.T) {
				c.bst.Dump()
				res := c.Book3(tc.interval.start, tc.interval.end)
				if res != tc.result {
					t.Errorf("it: %v, expected: %v, real: %v\n", tc.interval, tc.result, res)
				}
			})
		}
	})
}
