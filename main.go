package main

import (
	"fmt"
)

type Point struct {
	X, Y int
}

func (p *Point) distance(q *Point) int {
	return (p.X-q.X)*(p.X-q.X) + (p.Y-q.Y)*(p.Y-q.Y)
}

func main() {
	var x1, y1, x2, y2 int
	fmt.Scan(&x1)
	fmt.Scan(&y1)
	fmt.Scan(&x2)
	fmt.Scan(&y2)
	// Build two Points and print squared distance.
	p1 := Point{x1, y1}
	p2 := Point{x2, y2}
	fmt.Println(p1.distance(&p2))
}
