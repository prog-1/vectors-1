package main

import "testing"

func TestIsOnLine(t *testing.T) {
	for _, tc := range []struct {
		A, B, C point
		want    bool
	}{
		{point{1, 1}, point{3, 3}, point{5, 5}, true},    // https://www.desmos.com/calculator/fqed9n4t4q
		{point{1, 2}, point{3, 4}, point{5, 5}, false},   // https://www.desmos.com/calculator/dljppkeiko
		{point{1, 2}, point{3, 4}, point{6, 7}, true},    // https://www.desmos.com/calculator/e0cz8gir0o
		{point{0, 1}, point{2, 11}, point{1, 6}, true},   // https://www.desmos.com/calculator/fncu83qbyv
		{point{0, 1}, point{-2, 11}, point{1, 6}, false}, // https://www.desmos.com/calculator/zw4lbsogi2 //fails

	} {
		if got := IsOnLine(tc.A, tc.B, tc.C); got != tc.want {
			t.Errorf("IsOnLine(%v, %v, %v) = %v, want = %v", tc.A, tc.B, tc.C, got, tc.want)
		}
	}
}
