package main

import "testing"

func TestIsLinePassing(t *testing.T) {
	for _, tc := range []struct {
		a, b, c Point
		want    bool
	}{
		{Point{-1, 6}, Point{3, -2}, Point{2, 0}, true},
		{Point{-1, 6}, Point{3, -2}, Point{3, 0}, false},
		{Point{-2, -1}, Point{2, 3}, Point{1, 2}, true},
		{Point{-2, -1}, Point{2, 3}, Point{1, 1}, false},
		{Point{-6, 17}, Point{53, -160}, Point{45, -136}, true},
		{Point{-6, 17}, Point{53, -160}, Point{31, -90}, false},
		{Point{-4, -5}, Point{100, -5}, Point{25, -5}, true},
		{Point{-4, -5}, Point{100, -5}, Point{25, -100}, false},
		{Point{0, 4}, Point{0, 16}, Point{0, 9}, true},
		{Point{0, 4}, Point{0, 16}, Point{1, 9}, false},
	} {
		if got := isLinePassing(tc.a, tc.b, tc.c); got != tc.want {
			t.Errorf("Got = %v, Want = %v", got, tc.want)
		}
	}
}
