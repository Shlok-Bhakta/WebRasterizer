//go:build js && wasm
// +build js,wasm

package main

import (
	"math/rand"
)

type pixel struct {
	red   uint8
	green uint8
	blue  uint8
}

func (p *pixel) random() {
	p.red = uint8(rand.Intn(255))
	p.green = uint8(rand.Intn(255))
	p.blue = uint8(rand.Intn(255))
}

func make_random_pixel() pixel {
	p := pixel{}
	p.random()
	return p
}
