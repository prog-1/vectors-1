package main

type point struct {
	x, y int
}

func calculate(a, b, c point) bool {
	v := point{a.x - b.x, a.y - b.y}
	return (c.x-a.x)*v.y == (c.y-a.y)*v.x
}
