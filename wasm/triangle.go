//go:build js && wasm
// +build js,wasm

package main

import (
	"math"
)

type triangle struct {
	points [3]point3d
	color  pixel
}

func (t *triangle) draw(canvasdata *canvas) {
	// flatten triangle
	flat_t := [3]point3d{
		{x: t.points[0].x, y: t.points[0].y, z: 0},
		{x: t.points[1].x, y: t.points[1].y, z: 0},
		{x: t.points[2].x, y: t.points[2].y, z: 0},
	}
	// convert the flattened triangle to screenspace
	// draw the triangle on the canvas
	screen_triangle := screen_triangle{
		points: [3]screen_point{
			{x: int(flat_t[0].x), y: int(flat_t[0].y)},
			{x: int(flat_t[1].x), y: int(flat_t[1].y)},
			{x: int(flat_t[2].x), y: int(flat_t[2].y)},
		},
		color: t.color,
	}
	// get top left bound as a screen point (min x, min y)
	top_left := screen_point{
		x: int(math.Min(math.Min(float64(screen_triangle.points[0].x), float64(screen_triangle.points[1].x)), float64(screen_triangle.points[2].x))),
		y: int(math.Min(math.Min(float64(screen_triangle.points[0].y), float64(screen_triangle.points[1].y)), float64(screen_triangle.points[2].y))),
	}

	// get bottom right bound as a screen point (max x, max y)
	bottom_right := screen_point{
		x: int(math.Max(math.Max(float64(screen_triangle.points[0].x), float64(screen_triangle.points[1].x)), float64(screen_triangle.points[2].x))),
		y: int(math.Max(math.Max(float64(screen_triangle.points[0].y), float64(screen_triangle.points[1].y)), float64(screen_triangle.points[2].y))),
	}

	// draw these points to the canvas using the color and the bounding box
	for x := top_left.x; x <= bottom_right.x; x++ {
		for y := top_left.y; y <= bottom_right.y; y++ {
			p := screen_point{x: x, y: y}
			if screen_triangle.is_inside(p) {
				canvasdata.pixels[p.x][p.y] = screen_triangle.color
			}
		}
	}
}

func (t *triangle) rotate(angle float64, pivot *point3d) {
	// rotate each point of the triangle around the pivot
	for i := 0; i < 3; i++ {
		t.points[i].rotate(angle, pivot)
	}
}

func (t *triangle) get_center() point3d {
	// get the center of the triangle
	center := point3d{
		x: (t.points[0].x + t.points[1].x + t.points[2].x) / 3,
		y: (t.points[0].y + t.points[1].y + t.points[2].y) / 3,
	}
	return center
}

type screen_triangle struct {
	points [3]screen_point
	color  pixel
}

// checks if a given point falls inside the screen_triangle
func (t *screen_triangle) is_inside(p screen_point) bool {
	big_triangle_area := t.area()
	// find the area of all 3 triangles formed by the point and the triangle vertices
	t1 := screen_triangle{points: [3]screen_point{t.points[0], t.points[1], p}, color: t.color}
	t2 := screen_triangle{points: [3]screen_point{t.points[1], t.points[2], p}, color: t.color}
	t3 := screen_triangle{points: [3]screen_point{t.points[2], t.points[0], p}, color: t.color}
	a1 := t1.area()
	a2 := t2.area()
	a3 := t3.area()
	// check to see if a1 + a2 + a3 == big_triangle_area
	return a1+a2+a3-big_triangle_area <= 0.001
}

func (t *screen_triangle) area() float64 {
	// using the shoelace formula
	d1 := t.points[0].x*t.points[1].y - t.points[1].x*t.points[0].y
	d2 := t.points[1].x*t.points[2].y - t.points[2].x*t.points[1].y
	d3 := t.points[2].x*t.points[0].y - t.points[0].x*t.points[2].y
	return math.Abs(float64((d1 + d2 + d3) / 2))
}
