package main

import "testing"

func TestPasses(t *testing.T) {
	for _, tc := range []struct {
		name    string
		A, B, C Point
		want    bool
	}{
		{"1", Point{4, -1}, Point{0, 5}, Point{4, -1}, true},
		{"2", Point{4, -1}, Point{0, 5}, Point{0, 5}, true},
		{"3", Point{4, -1}, Point{0, 5}, Point{2, -2}, false},
		{"4", Point{4, -1}, Point{0, 5}, Point{2, 2}, true},
		{"5", Point{0, 4}, Point{2, 0}, Point{0, 4}, true},
		{"6", Point{0, 4}, Point{2, 0}, Point{2, 0}, true},
		{"7", Point{0, 4}, Point{2, 0}, Point{1, -1}, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got := Passes(tc.A, tc.B, tc.C)
			if got != tc.want {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}
