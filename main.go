package main

import "fmt"

type Point struct {
	x, y int
}

func isLinePassing(a, b, c Point) bool {
	v := Point{b.x - a.x, b.y - a.y}
	return v.y*c.x-v.y*a.x == v.x*c.y-v.x*a.y
}

func main() {
	a := Point{x: -1, y: 6}
	b := Point{x: 3, y: -2}
	c := Point{x: 2, y: 0}
	fmt.Println(isLinePassing(a, b, c))
}
