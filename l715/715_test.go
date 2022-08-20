package l715

import (
	"testing"
)

func TestRangeModule(t *testing.T) {
	type Interval struct {
		start, end int
	}

	/*
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
	*/

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

	t.Run("Remove4", func(t *testing.T) {
		st := Constructor()
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
				st.AddRange(tc.arg.start, tc.arg.end)
			case 'r':
				st.RemoveRange(tc.arg.start, tc.arg.end)
			case 'q':
				res := st.QueryRange(tc.arg.start, tc.arg.end)
				if res != tc.result {
					t.Fatalf("query: %v, expected: %v, real: %v\n", tc.arg, tc.result, res)
				}
			}
		}
	})

	t.Run("Remove5", func(t *testing.T) {
		st := Constructor()
		var testcases = []struct {
			cmd    byte
			arg    Interval
			result bool // for query
		}{
			{'u', Interval{5, 8}, true},
			{'q', Interval{3, 4}, false},
			{'r', Interval{5, 6}, true},
			{'r', Interval{3, 6}, true},
			{'u', Interval{1, 3}, true},
			{'q', Interval{2, 3}, true},
			{'a', Interval{4, 8}, true},
			{'q', Interval{2, 3}, true},
			{'r', Interval{4, 9}, true},
		}

		for _, tc := range testcases {
			switch tc.cmd {
			case 'u':
				st.AddRange(tc.arg.start, tc.arg.end)
			case 'r':
				st.RemoveRange(tc.arg.start, tc.arg.end)
			case 'q':
				res := st.QueryRange(tc.arg.start, tc.arg.end)
				if res != tc.result {
					t.Fatalf("query: %v, expected: %v, real: %v\n", tc.arg, tc.result, res)
				}
			}
		}
	})

	t.Run("Remove6", func(t *testing.T) {
		st := Constructor()
		var testcases = []struct {
			cmd    byte
			arg    Interval
			result bool // for query
		}{
			{'u', Interval{6, 8}, true},
			{'r', Interval{7, 8}, true},
			{'r', Interval{8, 9}, true},
			{'u', Interval{8, 9}, true},
			{'r', Interval{1, 3}, true},
			{'u', Interval{1, 8}, true},
			{'q', Interval{2, 4}, true},
			{'q', Interval{2, 9}, true},
			{'q', Interval{4, 6}, true},
		}

		for _, tc := range testcases {
			switch tc.cmd {
			case 'u':
				st.AddRange(tc.arg.start, tc.arg.end)
			case 'r':
				st.RemoveRange(tc.arg.start, tc.arg.end)
			case 'q':
				res := st.QueryRange(tc.arg.start, tc.arg.end)
				if res != tc.result {
					t.Fatalf("query: %v, expected: %v, real: %v\n", tc.arg, tc.result, res)
				}
			}
		}
	})

	t.Run("Remove7", func(t *testing.T) {
		st := Constructor()
		var testcases = []struct {
			cmd    byte
			arg    Interval
			result bool // for query
		}{
			{'u', Interval{14, 100}, true},
			{'r', Interval{1, 8}, true},
			{'q', Interval{77, 80}, true},
			{'q', Interval{8, 43}, false},
			{'q', Interval{4, 13}, false},
			{'r', Interval{3, 9}, true},
			{'r', Interval{45, 49}, true},
			{'r', Interval{41, 90}, true},
			{'u', Interval{58, 79}, true},
			{'u', Interval{4, 83}, true},
			{'u', Interval{34, 39}, true},
			{'r', Interval{84, 100}, true},
			{'u', Interval{8, 9}, true},
			{'q', Interval{32, 56}, true},
			{'u', Interval{35, 46}, true},
			{'u', Interval{9, 100}, true},
			{'q', Interval{85, 99}, true},
			{'q', Interval{23, 33}, true},
			{'u', Interval{10, 31}, true},
			{'r', Interval{15, 45}, true},
			{'r', Interval{52, 70}, true},
			{'r', Interval{26, 42}, true},
			{'q', Interval{30, 70}, false},
			{'q', Interval{60, 69}, false},
			{'u', Interval{10, 94}, true},
			{'u', Interval{2, 89}, true},
			{'q', Interval{26, 39}, true},
			{'u', Interval{46, 93}, true},
			{'u', Interval{30, 83}, true},
			{'r', Interval{42, 48}, true},
			{'u', Interval{47, 74}, true},
			{'u', Interval{39, 45}, true},
			{'q', Interval{14, 64}, false},
			{'r', Interval{3, 97}, true},
			{'q', Interval{16, 34}, false},
			{'r', Interval{28, 100}, true},
			{'u', Interval{19, 37}, true},
			{'u', Interval{27, 91}, true},
			{'q', Interval{55, 62}, true},
			{'r', Interval{64, 65}, true},
			{'r', Interval{2, 48}, true},
			{'u', Interval{55, 78}, true},
			{'q', Interval{21, 89}, false},
			{'q', Interval{31, 76}, false},
			{'r', Interval{13, 32}, true},
			{'r', Interval{2, 84}, true},
			{'r', Interval{21, 88}, true},
			{'q', Interval{12, 31}, false},
			{'u', Interval{89, 97}, true},
			{'r', Interval{56, 72}, true},
			{'r', Interval{16, 75}, true},
			{'q', Interval{18, 90}, false},
			{'r', Interval{46, 60}, true},
			{'r', Interval{20, 62}, true},
			{'q', Interval{28, 77}, false},
			{'u', Interval{5, 78}, true},
			{'u', Interval{58, 61}, true},
			{'r', Interval{38, 70}, true},
			{'q', Interval{24, 73}, false},
			{'q', Interval{72, 96}, false},
			{'u', Interval{5, 24}, true},
			{'r', Interval{43, 49}, true},
			{'r', Interval{2, 20}, true},
			{'u', Interval{4, 69}, true},
			{'u', Interval{18, 98}, true},
			{'u', Interval{26, 42}, true},
			{'u', Interval{14, 18}, true},
			{'q', Interval{46, 58}, true},
			{'r', Interval{16, 90}, true},
			{'u', Interval{32, 47}, true},
			{'u', Interval{19, 36}, true},
			{'u', Interval{26, 78}, true},
			{'q', Interval{7, 58}, false},
			{'u', Interval{42, 54}, true},
			{'r', Interval{42, 83}, true},
			{'q', Interval{3, 83}, false},
			{'r', Interval{54, 82}, true},
			{'r', Interval{71, 91}, true},
			{'r', Interval{22, 37}, true},
			{'q', Interval{38, 94}, false},
			{'q', Interval{20, 44}, false},
			{'q', Interval{37, 89}, false},
			{'q', Interval{15, 54}, false},
			{'q', Interval{1, 64}, false},
			{'r', Interval{63, 65}, true},
			{'q', Interval{55, 58}, false},
			{'r', Interval{23, 44}, true},
			{'q', Interval{25, 87}, false},
			{'u', Interval{38, 85}, true},
			{'q', Interval{27, 71}, false},
		}
		for _, tc := range testcases {
			switch tc.cmd {
			case 'u':
				st.AddRange(tc.arg.start, tc.arg.end)
			case 'r':
				st.RemoveRange(tc.arg.start, tc.arg.end)
			case 'q':
				res := st.QueryRange(tc.arg.start, tc.arg.end)
				if res != tc.result {
					t.Fatalf("query: %v, expected: %v, real: %v\n", tc.arg, tc.result, res)
				}
			}
		}
	})

	t.Run("Remove8", func(t *testing.T) {
		st := Constructor()
		var testcases = []struct {
			cmd    byte
			arg    Interval
			result bool // for query
		}{
			{'u', Interval{55, 62}, true},
			{'u', Interval{1, 29}, true},
			{'r', Interval{18, 49}, true},
			{'q', Interval{6, 98}, false},
			{'q', Interval{59, 71}, false},
			{'r', Interval{40, 45}, true},
			{'r', Interval{4, 58}, true},
			{'r', Interval{57, 69}, true},
			{'r', Interval{20, 30}, true},
			{'r', Interval{1, 40}, true},
			{'q', Interval{73, 93}, false},
			{'r', Interval{32, 93}, true},
			{'u', Interval{38, 100}, true},
			{'r', Interval{50, 64}, true},
			{'u', Interval{26, 72}, true},
			{'q', Interval{8, 74}, false},
			{'q', Interval{15, 53}, false},
			{'u', Interval{44, 85}, true},
			{'u', Interval{10, 71}, true},
			{'q', Interval{54, 70}, true},
			{'r', Interval{10, 45}, true},
			{'q', Interval{30, 66}, false},
			{'u', Interval{47, 98}, true},
			{'q', Interval{1, 7}, false},
			{'r', Interval{44, 78}, true},
			{'r', Interval{31, 49}, true},
			{'u', Interval{62, 63}, true},
			{'u', Interval{49, 88}, true},
			{'r', Interval{47, 72}, true},
			{'r', Interval{8, 50}, true},
			{'r', Interval{49, 79}, true},
			{'u', Interval{31, 47}, true},
			{'u', Interval{54, 87}, true},
			{'q', Interval{77, 78}, true},
			{'q', Interval{59, 100}, true},
			{'q', Interval{8, 9}, false},
			{'q', Interval{50, 51}, false},
			{'q', Interval{67, 93}, true},
			{'r', Interval{25, 86}, true},
			{'r', Interval{8, 92}, true},
			{'q', Interval{31, 87}, false},
			{'u', Interval{90, 95}, true},
			{'u', Interval{28, 56}, true},
			{'u', Interval{10, 42}, true},
			{'q', Interval{27, 34}, true},
			{'u', Interval{75, 81}, true},
			{'u', Interval{17, 63}, true},
			{'r', Interval{78, 90}, true},
			{'u', Interval{9, 18}, true},
			{'q', Interval{51, 74}, false},
			{'r', Interval{20, 54}, true},
			{'u', Interval{35, 72}, true},
			{'q', Interval{2, 29}, false},
			{'u', Interval{28, 41}, true},
			{'u', Interval{17, 95}, true},
			{'u', Interval{73, 75}, true},
			{'q', Interval{34, 43}, true},
			{'u', Interval{57, 96}, true},
			{'q', Interval{51, 72}, true},
			{'r', Interval{21, 67}, true},
			{'r', Interval{40, 73}, true},
			{'r', Interval{14, 26}, true},
			{'r', Interval{71, 86}, true},
			{'q', Interval{34, 41}, false},
			{'r', Interval{10, 25}, true},
			{'q', Interval{27, 68}, false},
			{'q', Interval{18, 32}, false},
			{'r', Interval{30, 31}, true},
			{'q', Interval{45, 61}, false},
			{'u', Interval{64, 66}, true},
			{'u', Interval{18, 93}, true},
			{'q', Interval{13, 21}, false},
			{'r', Interval{13, 46}, true},
			{'r', Interval{56, 99}, true},
			{'q', Interval{6, 93}, false},
			{'u', Interval{25, 36}, true},
			{'r', Interval{27, 88}, true},
			{'r', Interval{82, 83}, true},
			{'u', Interval{30, 71}, true},
			{'u', Interval{31, 73}, true},
			{'u', Interval{10, 41}, true},
			{'q', Interval{71, 72}, true},
			{'q', Interval{9, 56}, true},
			{'u', Interval{22, 76}, true},
			{'q', Interval{38, 74}, true},
			{'r', Interval{2, 77}, true},
			{'q', Interval{33, 61}, false},
			{'r', Interval{74, 75}, true},
			{'u', Interval{11, 43}, true},
			{'q', Interval{27, 75}, false},
		}
		for _, tc := range testcases {
			switch tc.cmd {
			case 'u':
				st.AddRange(tc.arg.start, tc.arg.end)
			case 'r':
				st.RemoveRange(tc.arg.start, tc.arg.end)
			case 'q':
				res := st.QueryRange(tc.arg.start, tc.arg.end)
				if res != tc.result {
					t.Fatalf("query: %v, expected: %v, real: %v\n", tc.arg, tc.result, res)
				}
			}
		}
	})
}

func TestRangeModule2(t *testing.T) {
	t.Run("Remove1", func(t *testing.T) {
		st := Constructor2()
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
				st.AddRange(tc.arg.start, tc.arg.end)
			case 'r':
				st.RemoveRange(tc.arg.start, tc.arg.end)
			case 'q':
				res := st.QueryRange(tc.arg.start, tc.arg.end)
				if res != tc.result {
					t.Fatalf("query: %v, expected: %v, real: %v\n", tc.arg, tc.result, res)
				}
			}
		}
	})

	t.Run("Remove2", func(t *testing.T) {
		st := Constructor2()
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
				st.AddRange(tc.arg.start, tc.arg.end)
			case 'r':
				st.RemoveRange(tc.arg.start, tc.arg.end)
			case 'q':
				res := st.QueryRange(tc.arg.start, tc.arg.end)
				if res != tc.result {
					t.Fatalf("query: %v, expected: %v, real: %v\n", tc.arg, tc.result, res)
				}
			}
		}
	})

	t.Run("Remove3", func(t *testing.T) {
		st := Constructor2()
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
				st.AddRange(tc.arg.start, tc.arg.end)
			case 'r':
				st.RemoveRange(tc.arg.start, tc.arg.end)
			case 'q':
				res := st.QueryRange(tc.arg.start, tc.arg.end)
				if res != tc.result {
					t.Fatalf("query: %v, expected: %v, real: %v\n", tc.arg, tc.result, res)
				}
			}
		}
	})

	t.Run("Remove4", func(t *testing.T) {
		st := Constructor2()
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
				st.AddRange(tc.arg.start, tc.arg.end)
			case 'r':
				st.RemoveRange(tc.arg.start, tc.arg.end)
			case 'q':
				res := st.QueryRange(tc.arg.start, tc.arg.end)
				if res != tc.result {
					t.Fatalf("query: %v, expected: %v, real: %v\n", tc.arg, tc.result, res)
				}
			}
		}
	})

	t.Run("Remove5", func(t *testing.T) {
		st := Constructor2()
		var testcases = []struct {
			cmd    byte
			arg    Interval
			result bool // for query
		}{
			{'u', Interval{5, 8}, true},
			{'q', Interval{3, 4}, false},
			{'r', Interval{5, 6}, true},
			{'r', Interval{3, 6}, true},
			{'u', Interval{1, 3}, true},
			{'q', Interval{2, 3}, true},
			{'a', Interval{4, 8}, true},
			{'q', Interval{2, 3}, true},
			{'r', Interval{4, 9}, true},
		}

		for _, tc := range testcases {
			switch tc.cmd {
			case 'u':
				st.AddRange(tc.arg.start, tc.arg.end)
			case 'r':
				st.RemoveRange(tc.arg.start, tc.arg.end)
			case 'q':
				res := st.QueryRange(tc.arg.start, tc.arg.end)
				if res != tc.result {
					t.Fatalf("query: %v, expected: %v, real: %v\n", tc.arg, tc.result, res)
				}
			}
		}
	})

	t.Run("Remove6", func(t *testing.T) {
		st := Constructor2()
		var testcases = []struct {
			cmd    byte
			arg    Interval
			result bool // for query
		}{
			{'u', Interval{6, 8}, true},
			{'r', Interval{7, 8}, true},
			{'r', Interval{8, 9}, true},
			{'u', Interval{8, 9}, true},
			{'r', Interval{1, 3}, true},
			{'u', Interval{1, 8}, true},
			{'q', Interval{2, 4}, true},
			{'q', Interval{2, 9}, true},
			{'q', Interval{4, 6}, true},
		}

		for _, tc := range testcases {
			switch tc.cmd {
			case 'u':
				st.AddRange(tc.arg.start, tc.arg.end)
			case 'r':
				st.RemoveRange(tc.arg.start, tc.arg.end)
			case 'q':
				res := st.QueryRange(tc.arg.start, tc.arg.end)
				if res != tc.result {
					t.Fatalf("query: %v, expected: %v, real: %v\n", tc.arg, tc.result, res)
				}
			}
		}
	})

	t.Run("Remove7", func(t *testing.T) {
		st := Constructor2()
		var testcases = []struct {
			cmd    byte
			arg    Interval
			result bool // for query
		}{
			{'u', Interval{14, 100}, true},
			{'r', Interval{1, 8}, true},
			{'q', Interval{77, 80}, true},
			{'q', Interval{8, 43}, false},
			{'q', Interval{4, 13}, false},
			{'r', Interval{3, 9}, true},
			{'r', Interval{45, 49}, true},
			{'r', Interval{41, 90}, true},
			{'u', Interval{58, 79}, true},
			{'u', Interval{4, 83}, true},
			{'u', Interval{34, 39}, true},
			{'r', Interval{84, 100}, true},
			{'u', Interval{8, 9}, true},
			{'q', Interval{32, 56}, true},
			{'u', Interval{35, 46}, true},
			{'u', Interval{9, 100}, true},
			{'q', Interval{85, 99}, true},
			{'q', Interval{23, 33}, true},
			{'u', Interval{10, 31}, true},
			{'r', Interval{15, 45}, true},
			{'r', Interval{52, 70}, true},
			{'r', Interval{26, 42}, true},
			{'q', Interval{30, 70}, false},
			{'q', Interval{60, 69}, false},
			{'u', Interval{10, 94}, true},
			{'u', Interval{2, 89}, true},
			{'q', Interval{26, 39}, true},
			{'u', Interval{46, 93}, true},
			{'u', Interval{30, 83}, true},
			{'r', Interval{42, 48}, true},
			{'u', Interval{47, 74}, true},
			{'u', Interval{39, 45}, true},
			{'q', Interval{14, 64}, false},
			{'r', Interval{3, 97}, true},
			{'q', Interval{16, 34}, false},
			{'r', Interval{28, 100}, true},
			{'u', Interval{19, 37}, true},
			{'u', Interval{27, 91}, true},
			{'q', Interval{55, 62}, true},
			{'r', Interval{64, 65}, true},
			{'r', Interval{2, 48}, true},
			{'u', Interval{55, 78}, true},
			{'q', Interval{21, 89}, false},
			{'q', Interval{31, 76}, false},
			{'r', Interval{13, 32}, true},
			{'r', Interval{2, 84}, true},
			{'r', Interval{21, 88}, true},
			{'q', Interval{12, 31}, false},
			{'u', Interval{89, 97}, true},
			{'r', Interval{56, 72}, true},
			{'r', Interval{16, 75}, true},
			{'q', Interval{18, 90}, false},
			{'r', Interval{46, 60}, true},
			{'r', Interval{20, 62}, true},
			{'q', Interval{28, 77}, false},
			{'u', Interval{5, 78}, true},
			{'u', Interval{58, 61}, true},
			{'r', Interval{38, 70}, true},
			{'q', Interval{24, 73}, false},
			{'q', Interval{72, 96}, false},
			{'u', Interval{5, 24}, true},
			{'r', Interval{43, 49}, true},
			{'r', Interval{2, 20}, true},
			{'u', Interval{4, 69}, true},
			{'u', Interval{18, 98}, true},
			{'u', Interval{26, 42}, true},
			{'u', Interval{14, 18}, true},
			{'q', Interval{46, 58}, true},
			{'r', Interval{16, 90}, true},
			{'u', Interval{32, 47}, true},
			{'u', Interval{19, 36}, true},
			{'u', Interval{26, 78}, true},
			{'q', Interval{7, 58}, false},
			{'u', Interval{42, 54}, true},
			{'r', Interval{42, 83}, true},
			{'q', Interval{3, 83}, false},
			{'r', Interval{54, 82}, true},
			{'r', Interval{71, 91}, true},
			{'r', Interval{22, 37}, true},
			{'q', Interval{38, 94}, false},
			{'q', Interval{20, 44}, false},
			{'q', Interval{37, 89}, false},
			{'q', Interval{15, 54}, false},
			{'q', Interval{1, 64}, false},
			{'r', Interval{63, 65}, true},
			{'q', Interval{55, 58}, false},
			{'r', Interval{23, 44}, true},
			{'q', Interval{25, 87}, false},
			{'u', Interval{38, 85}, true},
			{'q', Interval{27, 71}, false},
		}
		for _, tc := range testcases {
			switch tc.cmd {
			case 'u':
				st.AddRange(tc.arg.start, tc.arg.end)
			case 'r':
				st.RemoveRange(tc.arg.start, tc.arg.end)
			case 'q':
				res := st.QueryRange(tc.arg.start, tc.arg.end)
				if res != tc.result {
					t.Fatalf("query: %v, expected: %v, real: %v\n", tc.arg, tc.result, res)
				}
			}
		}
	})

	t.Run("Remove8", func(t *testing.T) {
		st := Constructor2()
		var testcases = []struct {
			cmd    byte
			arg    Interval
			result bool // for query
		}{
			{'u', Interval{55, 62}, true},
			{'u', Interval{1, 29}, true},
			{'r', Interval{18, 49}, true},
			{'q', Interval{6, 98}, false},
			{'q', Interval{59, 71}, false},
			{'r', Interval{40, 45}, true},
			{'r', Interval{4, 58}, true},
			{'r', Interval{57, 69}, true},
			{'r', Interval{20, 30}, true},
			{'r', Interval{1, 40}, true},
			{'q', Interval{73, 93}, false},
			{'r', Interval{32, 93}, true},
			{'u', Interval{38, 100}, true},
			{'r', Interval{50, 64}, true},
			{'u', Interval{26, 72}, true},
			{'q', Interval{8, 74}, false},
			{'q', Interval{15, 53}, false},
			{'u', Interval{44, 85}, true},
			{'u', Interval{10, 71}, true},
			{'q', Interval{54, 70}, true},
			{'r', Interval{10, 45}, true},
			{'q', Interval{30, 66}, false},
			{'u', Interval{47, 98}, true},
			{'q', Interval{1, 7}, false},
			{'r', Interval{44, 78}, true},
			{'r', Interval{31, 49}, true},
			{'u', Interval{62, 63}, true},
			{'u', Interval{49, 88}, true},
			{'r', Interval{47, 72}, true},
			{'r', Interval{8, 50}, true},
			{'r', Interval{49, 79}, true},
			{'u', Interval{31, 47}, true},
			{'u', Interval{54, 87}, true},
			{'q', Interval{77, 78}, true},
			{'q', Interval{59, 100}, true},
			{'q', Interval{8, 9}, false},
			{'q', Interval{50, 51}, false},
			{'q', Interval{67, 93}, true},
			{'r', Interval{25, 86}, true},
			{'r', Interval{8, 92}, true},
			{'q', Interval{31, 87}, false},
			{'u', Interval{90, 95}, true},
			{'u', Interval{28, 56}, true},
			{'u', Interval{10, 42}, true},
			{'q', Interval{27, 34}, true},
			{'u', Interval{75, 81}, true},
			{'u', Interval{17, 63}, true},
			{'r', Interval{78, 90}, true},
			{'u', Interval{9, 18}, true},
			{'q', Interval{51, 74}, false},
			{'r', Interval{20, 54}, true},
			{'u', Interval{35, 72}, true},
			{'q', Interval{2, 29}, false},
			{'u', Interval{28, 41}, true},
			{'u', Interval{17, 95}, true},
			{'u', Interval{73, 75}, true},
			{'q', Interval{34, 43}, true},
			{'u', Interval{57, 96}, true},
			{'q', Interval{51, 72}, true},
			{'r', Interval{21, 67}, true},
			{'r', Interval{40, 73}, true},
			{'r', Interval{14, 26}, true},
			{'r', Interval{71, 86}, true},
			{'q', Interval{34, 41}, false},
			{'r', Interval{10, 25}, true},
			{'q', Interval{27, 68}, false},
			{'q', Interval{18, 32}, false},
			{'r', Interval{30, 31}, true},
			{'q', Interval{45, 61}, false},
			{'u', Interval{64, 66}, true},
			{'u', Interval{18, 93}, true},
			{'q', Interval{13, 21}, false},
			{'r', Interval{13, 46}, true},
			{'r', Interval{56, 99}, true},
			{'q', Interval{6, 93}, false},
			{'u', Interval{25, 36}, true},
			{'r', Interval{27, 88}, true},
			{'r', Interval{82, 83}, true},
			{'u', Interval{30, 71}, true},
			{'u', Interval{31, 73}, true},
			{'u', Interval{10, 41}, true},
			{'q', Interval{71, 72}, true},
			{'q', Interval{9, 56}, true},
			{'u', Interval{22, 76}, true},
			{'q', Interval{38, 74}, true},
			{'r', Interval{2, 77}, true},
			{'q', Interval{33, 61}, false},
			{'r', Interval{74, 75}, true},
			{'u', Interval{11, 43}, true},
			{'q', Interval{27, 75}, false},
		}
		for _, tc := range testcases {
			switch tc.cmd {
			case 'u':
				st.AddRange(tc.arg.start, tc.arg.end)
			case 'r':
				st.RemoveRange(tc.arg.start, tc.arg.end)
			case 'q':
				res := st.QueryRange(tc.arg.start, tc.arg.end)
				if res != tc.result {
					t.Fatalf("query: %v, expected: %v, real: %v\n", tc.arg, tc.result, res)
				}
			}
		}
	})
}
