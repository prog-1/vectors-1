package main

import "fmt"

type point struct {
	x, y int
}

func IsOnLine(A, B, C point) bool {
	if B.x == A.x || B.y == A.y { //devision by zero
		return true
	}
	return (C.x-A.x)/(B.x-A.x) == (C.y-A.y)/(B.y-A.y)
}

func main() {
	fmt.Println(IsOnLine(point{1, 1}, point{3, 3}, point{5, 5})) // https://www.desmos.com/calculator/fqed9n4t4q
}
