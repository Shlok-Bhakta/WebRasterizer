//go:build js && wasm
// +build js,wasm

package main

import (
	// "fmt"
	"math/rand"
	"time"
)

func check_triange(triangle *triangle, vector *point, canvasdata *canvas) {
	// update the triangle position
	for i := 0; i < 3; i++ {
		triangle.points[i].x += vector.x
		triangle.points[i].y += vector.y
	}
	for i := 0; i < 3; i++ {
		// if the triangle goes out of bounds, reverse the vector
		if triangle.points[i].x < 0 || triangle.points[i].x >= canvasdata.width {
			// reverse the vector
			vector.x = -1 * vector.x
		}
		if triangle.points[i].y < 0 || triangle.points[i].y >= canvasdata.height {
			// reverse the vector
			vector.y = -1 * vector.y
		}
	}
}

func main() {
	canvasdata := canvas{}
	canvasdata.init()
	r := rand.New(rand.NewSource(50))
	// gonna move the triangle by this value in the direction
	// draw a triangle
	triangles := make([]triangle, 0)
	vectors := make([]point, 0)
	for i := 0; i < 10; i++ {
		// create a random triangle
		t := triangle{
			points: [3]point{
				{x: 5, y: 10},
				{x: 30, y: 40},
				{x: 5, y: 80},
			},
			color: make_random_pixel(),
		}
		triangles = append(triangles, t)
		// create a random vector
		v := point{
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

		// print the triangle points
		// fmt.Printf("Triangle points: %+v\n", triangle)
		// draw the triangles
		for k := 0; k < len(triangles); k++ {
			for i := 0; i < canvasdata.height; i++ {
				for j := 0; j < canvasdata.width; j++ {
					if triangles[k].is_inside(point{x: j, y: i}) {
						canvasdata.pixels[i][j] = triangles[k].color
					}
				}
			}
		}
		canvasdata.render()

		for k := 0; k < len(triangles); k++ {
			check_triange(&triangles[k], &vectors[k], &canvasdata)
		}
		// just sleep for a bit
		time.Sleep(time.Millisecond * 100)
	}
	// select {} // Keep the program running
}
