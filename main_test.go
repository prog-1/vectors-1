package main

import "testing"

func TestIncludesPoint(t *testing.T) {
	for tc := range []struct {
		l    Line
		p    Vector
		want bool
	}{
		{Line{Vect{2, 5}, {6, -2}}, {8, 3}, true},
	} {

	}
}
