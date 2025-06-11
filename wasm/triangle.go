//go:build js && wasm
// +build js,wasm

package main

import (
	"math"
)

type triangle struct {
	points [3]point
	color  pixel
}

// checks if a given point falls inside the triangle
func (t *triangle) is_inside(p point) bool {
	big_triangle_area := t.area()
	// find the area of all 3 triangles formed by the point and the triangle vertices
	t1 := triangle{points: [3]point{t.points[0], t.points[1], p}, color: t.color}
	t2 := triangle{points: [3]point{t.points[1], t.points[2], p}, color: t.color}
	t3 := triangle{points: [3]point{t.points[2], t.points[0], p}, color: t.color}
	a1 := t1.area()
	a2 := t2.area()
	a3 := t3.area()
	// check to see if a1 + a2 + a3 == big_triangle_area
	return a1+a2+a3-big_triangle_area <= 0.001
}

func (t *triangle) area() float64 {
	// using the shoelace formula
	d1 := t.points[0].x*t.points[1].y - t.points[1].x*t.points[0].y
	d2 := t.points[1].x*t.points[2].y - t.points[2].x*t.points[1].y
	d3 := t.points[2].x*t.points[0].y - t.points[0].x*t.points[2].y
	return math.Abs(float64((d1 + d2 + d3) / 2))
}

func (t *triangle) rotate(angle float64, pivot *point) {
	// rotate each point of the triangle around the pivot
	for i := 0; i < 3; i++ {
		t.points[i].rotate(angle, pivot)
	}
}

func (t *triangle) get_center() point {
	// get the center of the triangle
	center := point{
		x: (t.points[0].x + t.points[1].x + t.points[2].x) / 3,
		y: (t.points[0].y + t.points[1].y + t.points[2].y) / 3,
	}
	return center
}
