package main

import "testing"

func TestExclusiveTime(t *testing.T) {
	var testcases = []struct {
		n        int
		logs     []string
		expected []int
	}{
		{2, []string{"0:start:0", "1:start:2", "1:end:5", "0:end:6"}, []int{3, 4}},
		{1, []string{"0:start:0", "0:start:2", "0:end:5", "0:start:6", "0:end:6", "0:end:7"}, []int{8}},
		{2, []string{"0:start:0", "0:start:2", "0:end:5", "1:start:6", "1:end:6", "0:end:7"}, []int{7, 1}},
	}

	for _, tc := range testcases {
		res := exclusiveTime2(tc.n, tc.logs)
		if !compareSlicesNoSort(res, tc.expected) {
			t.Errorf("logs: %v, expected: %v, real: %v\n", tc.logs, tc.expected, res)
		}
	}
}
