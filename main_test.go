package main

import "testing"

func TestIncludesPoint(t *testing.T) {
	for _, tc := range []struct {
		l    Line
		p    Vector
		want bool
	}{
		{Line{Vector{2, 5}, Vector{6, -2}}, Vector{8, 3}, true},
		{Line{Vector{2, 5}, Vector{6, -2}}, Vector{5, 2}, false},
		{Line{Vector{2, 5}, Vector{6, -2}}, Vector{11, 2}, true},
	} {
		if got := tc.l.IncludesPoint(tc.p); got != tc.want {
			t.Errorf("got = %v, want = %v", got, tc.want)
		}
	}
}
