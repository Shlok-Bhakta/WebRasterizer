//go:build js && wasm
// +build js,wasm

package main

import (
	"math/rand"
)

func main() {
	canvasdata := canvas{}
	canvasdata.init()
	// set all pixels to red
	for i := 0; i < canvasdata.height; i++ {
		for j := 0; j < canvasdata.width; j++ {
			randred := rand.Intn(256)
			randgreen := rand.Intn(256)
			randblue := rand.Intn(256)
			canvasdata.pixels[i][j] = pixel{red: uint8(randred), green: uint8(randgreen), blue: uint8(randblue), alpha: 255}
		}
	}
	canvasdata.render()
	select {} // Keep the program running
}
