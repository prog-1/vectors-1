package main

import "testing"

func TestCalculate(t *testing.T) {
	for _, tc := range []struct {
		a, b, c point
		want    bool
	}{
		{point{0, 4}, point{2, 2}, point{4, 0}, true},
		{point{0, 4}, point{4, 0}, point{2, 2}, true},
		{point{0, 4}, point{4, 0}, point{2, 1}, false},
		{point{-6, 2}, point{-2, 0}, point{2, -2}, true},
		{point{-2, -3}, point{1, 2}, point{4, 7}, true},
		{point{-67, -3}, point{1, 32}, point{4, 32}, false},
		{point{2, 5}, point{6, -2}, point{5, 2}, false},
	} {
		if got := find(tc.a, tc.b, tc.c); got != tc.want {
			t.Errorf("got = %v, want = %v", got, tc.want)
		}
	}
}
