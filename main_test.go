package main

import "testing"

func TestBelongs(t *testing.T) {
	for _, tc := range []struct {
		name    string
		a, b, c point
		want    bool
	}{
		{"case-1", point{x: 1, y: 2}, point{x: -1, y: 2}, point{x: 2, y: 0}, true},
		{"case-2", point{x: 1, y: 2}, point{x: -1, y: 2}, point{x: 1, y: -1}, false},
		{"case-3", point{x: 1, y: 2}, point{x: -1, y: 2}, point{x: 2, y: 1}, false},
		{"case-4", point{x: 1, y: 2}, point{x: -1, y: 2}, point{x: 0, y: 4}, true},
		{"case-5", point{x: 1, y: 2}, point{x: -1, y: 2}, point{x: 1.5, y: 1}, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got := belongs(tc.a, tc.b, tc.c)
			if got != tc.want {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}
