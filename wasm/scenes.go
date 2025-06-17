//go:build js && wasm
// +build js,wasm

package main

import (
	// "fmt"
	"time"
)

func triangle3d() {
	canvasdata := canvas{}
	canvasdata.init()
	// draw a triangle
	t := triangle{
		points: [3]point3d{
			{x: 5, y: 5, z: 0},
			{x: 30, y: 5, z: 0},
			{x: 5, y: 30, z: 10},
		},
		color: make_random_pixel(),
	}
	for {
		for i := 0; i < canvasdata.height; i++ {
			for j := 0; j < canvasdata.width; j++ {
				canvasdata.pixels[i][j] = pixel{red: 255, green: 255, blue: 200}
			}
		}
		t.draw(&canvasdata)
		canvasdata.render()
		time.Sleep(time.Millisecond * 100)
	}
}
