package main

type Vector struct {
	x, y int
}

type Line struct {
	a, v Vector
}

func main() {

}

func (l *Line) IncludesPoint(p Vector) bool {
	return (p.x-l.a.x)/l.v.x == (p.y-l.a.y)/l.v.y
}
