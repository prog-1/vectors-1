package main

type point struct {
	x, y float64
}

func Calculate(a, b, c point) bool {
	return (c.x-a.x)/b.x == (c.y-a.y)/b.y
}
