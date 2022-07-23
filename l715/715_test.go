package l715

import "testing"

func TestRangeModule(t *testing.T) {
	type Interval struct {
		start, end int
	}

	t.Run("Remove1", func(t *testing.T) {
		st := NewIntervalQueryTree(1, 1e9)
		var testcases = []struct {
			cmd    byte
			arg    Interval
			result bool // for query
		}{
			{'u', Interval{10, 20}, true},
			{'r', Interval{14, 16}, true},
			{'q', Interval{10, 14}, true},
			{'q', Interval{13, 15}, false},
			{'q', Interval{16, 17}, true},
		}

		for _, tc := range testcases {
			switch tc.cmd {
			case 'u':
				st.Update(tc.arg.start, tc.arg.end-1)
			case 'r':
				st.Remove(tc.arg.start, tc.arg.end-1)
			case 'q':
				res := st.Query(tc.arg.start, tc.arg.end-1)
				if res != tc.result {
					t.Fatalf("query: %v, expected: %v, real: %v\n", tc.arg, tc.result, res)
				}
			}
		}
	})
	t.Run("Remove2", func(t *testing.T) {
		st := NewIntervalQueryTree(1, 1e9)
		var testcases = []struct {
			cmd    byte
			arg    Interval
			result bool // for query
		}{
			{'u', Interval{10, 180}, true},
			{'u', Interval{150, 200}, true},
			{'u', Interval{250, 500}, true},
			{'q', Interval{50, 100}, true},
			{'q', Interval{180, 300}, false},
			{'q', Interval{600, 1000}, false},
			{'r', Interval{50, 150}, true},
			{'q', Interval{50, 100}, false},
		}

		for _, tc := range testcases {
			switch tc.cmd {
			case 'u':
				st.Update(tc.arg.start, tc.arg.end-1)
			case 'r':
				st.Remove(tc.arg.start, tc.arg.end-1)
			case 'q':
				res := st.Query(tc.arg.start, tc.arg.end-1)
				if res != tc.result {
					t.Fatalf("query: %v, expected: %v, real: %v\n", tc.arg, tc.result, res)
				}
			}
		}
	})

	t.Run("Remove3", func(t *testing.T) {
		st := NewIntervalQueryTree(1, 1e9)
		var testcases = []struct {
			cmd    byte
			arg    Interval
			result bool // for query
		}{
			{'q', Interval{21, 34}, false},
			{'q', Interval{27, 87}, false},
			{'u', Interval{44, 53}, true},
			{'u', Interval{69, 89}, true},
			{'q', Interval{23, 26}, false},

			{'q', Interval{80, 84}, true},
			{'q', Interval{11, 12}, false},
			{'r', Interval{86, 91}, true},
			{'u', Interval{87, 94}, true},
			{'r', Interval{34, 52}, true},

			{'u', Interval{1, 59}, true},
			{'r', Interval{62, 96}, true},
			{'r', Interval{34, 83}, true},
			{'q', Interval{11, 59}, false},
			{'q', Interval{59, 79}, false},

			{'q', Interval{1, 13}, true},
			{'q', Interval{21, 23}, true},
			{'r', Interval{9, 61}, true},
			{'u', Interval{17, 21}, true},
			{'r', Interval{4, 8}, true},

			{'q', Interval{19, 25}, false},
		}

		for _, tc := range testcases {
			switch tc.cmd {
			case 'u':
				st.Update(tc.arg.start, tc.arg.end-1)
			case 'r':
				st.Remove(tc.arg.start, tc.arg.end-1)
			case 'q':
				res := st.Query(tc.arg.start, tc.arg.end-1)
				if res != tc.result {
					t.Fatalf("query: %v, expected: %v, real: %v\n", tc.arg, tc.result, res)
				}
			}
		}
	})
}
