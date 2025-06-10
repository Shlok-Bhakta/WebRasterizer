//go:build js && wasm
// +build js,wasm

package main

import (
	// "fmt"
	"math/rand"
	"time"
)

func main() {
	canvasdata := canvas{}
	canvasdata.init()
	r := rand.New(rand.NewSource(50))
	// gonna move the triangle by this value in the direction
	vector := point{x: canvasdata.mapWidth(r.Float64() / 8), y: canvasdata.mapHeight(r.Float64() / 8)}
	// draw a triangle
	triangle := triangle{
		points: [3]point{
			{x: 5, y: 10},
			{x: 30, y: 40},
			{x: 5, y: 80},
		},
		color: pixel{red: 255, green: 0, blue: 0},
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
		for i := 0; i < canvasdata.height; i++ {
			for j := 0; j < canvasdata.width; j++ {
				if triangle.is_inside(point{x: j, y: i}) {
					canvasdata.pixels[i][j] = triangle.color
				}
			}
		}
		canvasdata.render()

		// update the triangle position
		for i := 0; i < 3; i++ {
			triangle.points[i].x += vector.x
			triangle.points[i].y += vector.y
			// if the triangle goes out of bounds, reverse the vector
			if triangle.points[i].x < 0 || triangle.points[i].x >= canvasdata.width ||
				triangle.points[i].y < 0 || triangle.points[i].y >= canvasdata.height {
				// reverse the vector
				vector.x = -1 * vector.x
				vector.y = -1 * vector.y
			}
		}
		// just sleep for a bit
		time.Sleep(time.Millisecond * 100)
	}
	// select {} // Keep the program running
}
