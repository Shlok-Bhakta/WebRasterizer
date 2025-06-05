//go:build js && wasm
// +build js,wasm

package main

import (
	"syscall/js"
)

type pixel struct {
	red   uint8
	green uint8
	blue  uint8
}

type canvas struct {
	width     int
	height    int
	pixels    [][]pixel
	imagedata js.Value
	ctx       js.Value
}

func (c *canvas) init() {
	c.ctx = js.Global().Get("document").Call("getElementById", "canvas").Call("getContext", "2d")
	c.width = c.ctx.Get("canvas").Get("width").Int()
	c.height = c.ctx.Get("canvas").Get("height").Int()
	c.imagedata = c.ctx.Call("createImageData", js.ValueOf(c.width), js.ValueOf(c.height))
	c.pixels = make([][]pixel, c.height)
	for i := range c.pixels {
		c.pixels[i] = make([]pixel, c.width)
	}
}

func (c *canvas) setSize(w int, h int) {
	c.ctx.Set("canvas", js.ValueOf(map[string]interface{}{
		"width":  w,
		"height": h,
	}))
}

func (c *canvas) render() {
	// loop over width and height pixels and add them to the imagedata array
	for i := 0; i < c.height; i++ {
		for j := 0; j < c.width; j++ {

		}
	}
}
