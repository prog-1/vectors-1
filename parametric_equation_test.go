package main

import (
	"testing"
)

type point struct {
	x, y int
}

type line struct {
	ax, ay, bx, by int
}

//Function that calculates if the line is passing through the point
func (l line) isPassing(c point) bool {
	return l.by*(l.ax-c.x) == l.bx*(l.ay-c.y)
}

// func main() {
// 	l := line{ax: 4, ay: 3, bx: -3, by: 3}
// 	A := point{x: 9, y: -2}
// 	B := point{x: 7, y: 2}
// 	fmt.Println(l.isPassing(A))
// 	fmt.Println(l.isPassing(B))
// }

//--------------Tests-------------------------
func TestIsPassing(t *testing.T) {
	for _, tc := range []struct {
		name  string
		point point
		want  bool
	}{
		{"A", point{9, -2}, true},
		{"B", point{7, 2}, false},
		{"C", point{2, 0}, false},
		{"D", point{1, -1}, false},
		{"E", point{3, 4}, true},
		{"F", point{1, 6}, true},
		{"O", point{0, 0}, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			l := line{ax: 4, ay: 3, bx: -3, by: 3}
			got := l.isPassing(tc.point)
			if got != tc.want {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}
