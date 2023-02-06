package main

type Vector struct {
	x, y int
}

func (v *Vector) IsZero() bool {
	return v.x == 0 && v.y == 0
}

type Line struct {
	a, v Vector
}

func (l *Line) IncludesPoint(p Vector) bool {
	if l.v.IsZero() {
		return true
	}
	return (p.x-l.a.x)/l.v.x == (p.y-l.a.y)/l.v.y
}

func main() {

}
