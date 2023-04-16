package main

import "fmt"

type Point struct {
	x, y int
}

func Passes(A, B, C Point) bool { // C is the point
	v := Point{B.x - A.x, B.y - A.y}
	return (C.x-A.x)*v.y == (C.y-A.y)*v.x
}

func main() {
	// Based on my example (attached photos):
	fmt.Println(Passes(Point{4, -1}, Point{0, 5}, Point{4, -1})) // true
	fmt.Println(Passes(Point{4, -1}, Point{0, 5}, Point{0, 5}))  // true
	fmt.Println(Passes(Point{4, -1}, Point{0, 5}, Point{2, -2})) // false
}
