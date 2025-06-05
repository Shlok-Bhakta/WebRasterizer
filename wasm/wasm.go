//go:build js && wasm
// +build js,wasm

package main

func main() {
	canvasdata := canvas{}
	canvasdata.init()
	// set all pixels to red
	for i := 0; i < canvasdata.height; i++ {
		for j := 0; j < canvasdata.width; j++ {
			canvasdata.pixels[i][j] = pixel{red: 255, green: 255, blue: 200, alpha: 255}
		}
	}

	// draw a square
	for i := canvasdata.mapHeight(0); i < canvasdata.mapHeight(1); i++ {
		for j := canvasdata.mapWidth(0); j < canvasdata.mapWidth(1); j++ {
			canvasdata.pixels[i][j] = pixel{red: 0, green: 255, blue: 0, alpha: 255}
		}
	}
	canvasdata.render()
	select {} // Keep the program running
}
