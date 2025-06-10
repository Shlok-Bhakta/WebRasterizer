//go:build js && wasm
// +build js,wasm

package main

import (
	// "fmt"
	"time"
)

func main() {
	canvasdata := canvas{}
	canvasdata.init()
	xval := 5
	// render loop
	for {
		// set all pixels to red
		for i := 0; i < canvasdata.height; i++ {
			for j := 0; j < canvasdata.width; j++ {
				canvasdata.pixels[i][j] = pixel{red: 255, green: 255, blue: 200}
			}
		}

		// draw a triangle
		triangle := triangle{
			points: [3]point{
				{x: 5, y: 10},
				{x: 30, y: 40},
				{x: xval, y: 80},
			},
			color: pixel{red: 255, green: 0, blue: 0},
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
		xval += 1
		// just sleep for a bit
		time.Sleep(time.Millisecond * 100)
	}
	// select {} // Keep the program running
}
