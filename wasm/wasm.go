//go:build js && wasm
// +build js,wasm

package main

import (
	// "fmt"
	"math/rand"
	"time"
)

func check_triangle(triangle *triangle, vector *screen_point, canvasdata *canvas) {
	// add a random rotation to the triangle
	t_center := triangle.get_center()
	triangle.rotate(rand.Float64()*0.1, &t_center)

	// update the triangle position
	for i := 0; i < 3; i++ {
		triangle.points[i].x += float64(vector.x)
		triangle.points[i].y += float64(vector.y)
	}

	for i := 0; i < 3; i++ {
		screen_point := triangle.points[i].toScreen(canvasdata)
		// if the triangle goes out of bounds, reverse the vector
		if screen_point.x < 0 || screen_point.x >= canvasdata.width {
			// reverse the vector
			vector.x = -1 * vector.x
		}
		if screen_point.y < 0 || screen_point.y >= canvasdata.height {
			// reverse the vector
			vector.y = -1 * vector.y
		}
	}
}

func main() {
	canvasdata := canvas{}
	canvasdata.init()
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// gonna move the triangle by this value in the direction
	// draw a triangle
	triangles := make([]triangle, 0)
	vectors := make([]screen_point, 0)
	for i := 0; i < 100; i++ {
		// create a random triangle
		t := triangle{
			points: [3]point{
				{x: 5 + r.Float64()*10, y: 5 + r.Float64()*10},
				{x: 30 + r.Float64()*10, y: 5 + r.Float64()*10},
				{x: 5 + r.Float64()*10, y: 30 + r.Float64()*10},
			},
			color: make_random_pixel(),
		}
		triangles = append(triangles, t)
		// create a random vector
		v := screen_point{
			x: canvasdata.mapWidth(r.Float64() / 20),
			y: canvasdata.mapHeight(r.Float64() / 20),
		}
		vectors = append(vectors, v)
	}
	// render loop
	for {
		// set all pixels to red
		for i := 0; i < canvasdata.height; i++ {
			for j := 0; j < canvasdata.width; j++ {
				canvasdata.pixels[i][j] = pixel{red: 255, green: 255, blue: 200}
			}
		}

		// draw the triangles
		for k := 0; k < len(triangles); k++ {
			for i := 0; i < canvasdata.height; i++ {
				for j := 0; j < canvasdata.width; j++ {
					if triangles[k].is_inside(point{x: float64(j), y: float64(i)}) {
						canvasdata.pixels[i][j] = triangles[k].color
					}
				}
			}
		}
		canvasdata.render()

		for k := 0; k < len(triangles); k++ {
			check_triangle(&triangles[k], &vectors[k], &canvasdata)
		}
		// just sleep for a bit
		time.Sleep(time.Millisecond * 100)
	}
	// select {} // Keep the program running
}
