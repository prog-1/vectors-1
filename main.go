package main

import "fmt"

type Point struct {
	x, y float64
}

func find(a, b, c Point) bool {
	return (c.x-a.x)/b.x == (c.y-a.y)/b.y
}

func main() {
	a, b, c := Point{0, 0}, Point{5, 5}, Point{3, 8}
	fmt.Println(find(a, b, c))
}
