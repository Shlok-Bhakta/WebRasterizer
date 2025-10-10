//go:build js && wasm
// +build js,wasm

package main

import (
	"fmt"
	"math"
)

type triangle struct {
	points [3]point3d
	color  pixel
}

func (t *triangle) draw(canvasdata *canvas, cam *camera) {
	// draw the triangle on the canvas
	project_1, err := cam.project_point(t.points[0], canvasdata)
	if err != nil {
		fmt.Printf("Error projecting triangle point 1: %v\n", err)
		// return
	}
	project_2, err := cam.project_point(t.points[1], canvasdata)
	if err != nil {
		fmt.Printf("Error projecting triangle point 2: %v\n", err)
		// return
	}
	project_3, err := cam.project_point(t.points[2], canvasdata)
	if err != nil {
		fmt.Printf("Error projecting triangle point 3: %v\n", err)
		// return
	}

	screen_triangle := screen_triangle{
		points: [3]screen_point{
			project_1,
			project_2,
			project_3,
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
			// check if point is outsede the screen, if so do nothing
			if x < 0 || x >= canvasdata.width || y < 0 || y >= canvasdata.height {
				continue
			}
			p := screen_point{x: x, y: y}
			z_depth := screen_triangle.is_inside(p)
			if z_depth > 0 && z_depth < canvasdata.zbuffer[p.y][p.x] {
				canvasdata.pixels[p.y][p.x] = screen_triangle.color
				canvasdata.zbuffer[p.y][p.x] = z_depth
			}
		}
	}
}

// roll rotates the triangle around a pivot point
func (t *triangle) transform(roll float64, pitch float64, yaw float64, pivot *point3d) {
	// construct a transformation matrix for the roll, pitch, and yaw
	roll_matrix := matrix4x4{
		{1, 0, 0, 0},
		{0, math.Cos(roll), -math.Sin(roll), 0},
		{0, math.Sin(roll), math.Cos(roll), 0},
		{0, 0, 0, 1},
	}
	pitch_matrix := matrix4x4{
		{math.Cos(pitch), 0, -math.Sin(pitch), 0},
		{0, 1, 0, 0},
		{math.Sin(pitch), 0, math.Cos(pitch), 0},
		{0, 0, 0, 1},
	}
	yaw_matrix := matrix4x4{
		{math.Cos(yaw), -math.Sin(yaw), 0, 0},
		{math.Sin(yaw), math.Cos(yaw), 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}

	// multiply the transformation matrices together
	matrix := yaw_matrix.multiply(&pitch_matrix)
	matrix = matrix.multiply(&roll_matrix)

	// apply the transformation matrix to the triangle points with pivot
	for i := 0; i < 3; i++ {
		// Translate to origin (subtract pivot)
		translated := point3d{
			x: t.points[i].x - pivot.x,
			y: t.points[i].y - pivot.y,
			z: t.points[i].z - pivot.z,
		}

		// Apply rotation
		rotated := translated.transform(&matrix)

		// Translate back (add pivot)
		t.points[i] = point3d{
			x: rotated.x + pivot.x,
			y: rotated.y + pivot.y,
			z: rotated.z + pivot.z,
		}
	}
}

func (t *triangle) get_center() point3d {
	// get the center of the triangle
	center := point3d{
		x: (t.points[0].x + t.points[1].x + t.points[2].x) / 3,
		y: (t.points[0].y + t.points[1].y + t.points[2].y) / 3,
		z: (t.points[0].z + t.points[1].z + t.points[2].z) / 3,
	}
	return center
}

type screen_triangle struct {
	points [3]screen_point
	color  pixel
}

// checks if a given point falls inside the screen_triangle and also returns the z depth
func (t *screen_triangle) is_inside(p screen_point) float64 {
	big_triangle_area := t.area()
	// find the area of all 3 triangles formed by the point and the triangle vertices
	t1 := screen_triangle{points: [3]screen_point{t.points[0], t.points[1], p}, color: t.color}
	t2 := screen_triangle{points: [3]screen_point{t.points[1], t.points[2], p}, color: t.color}
	t3 := screen_triangle{points: [3]screen_point{t.points[2], t.points[0], p}, color: t.color}
	a1 := t1.area()
	a2 := t2.area()
	a3 := t3.area()
	// check to see if a1 + a2 + a3 == big_triangle_area
	if a1+a2+a3-big_triangle_area > 0.1 {
		return -1
	} else {
		w1 := a2 / big_triangle_area
		w2 := a3 / big_triangle_area
		w3 := a1 / big_triangle_area
		return t.points[0].z*w1 + t.points[1].z*w2 + t.points[2].z*w3
	}
}

func (t *screen_triangle) area() float64 {
	// using the shoelace formula
	d1 := t.points[0].x*t.points[1].y - t.points[1].x*t.points[0].y
	d2 := t.points[1].x*t.points[2].y - t.points[2].x*t.points[1].y
	d3 := t.points[2].x*t.points[0].y - t.points[0].x*t.points[2].y
	return math.Abs(float64((d1 + d2 + d3) / 2))
}
