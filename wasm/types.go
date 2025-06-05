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
	alpha uint8
}

type canvas struct {
	width     int
	height    int
	pixels    [][]pixel
	imagedata js.Value
	ctx       js.Value
	element   js.Value
}

func (c *canvas) init() {
	c.element = js.Global().Get("document").Call("getElementById", "canvas")
	c.ctx = js.Global().Get("document").Call("getElementById", "canvas").Call("getContext", "2d")
	c.width = c.ctx.Get("canvas").Get("width").Int()
	c.height = c.ctx.Get("canvas").Get("height").Int()
	c.setSizeFromDocument()
	// c.setSize(700, 500)
	// Setup resize listener
	js.Global().Call("addEventListener", "resize", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		c.setSizeFromDocument()
		c.render()
		return nil
	}))
	c.imagedata = c.ctx.Call("createImageData", js.ValueOf(c.width), js.ValueOf(c.height))
	c.pixels = make([][]pixel, c.height)
	for i := range c.pixels {
		c.pixels[i] = make([]pixel, c.width)
	}
}

func (c *canvas) setSize(w int, h int) {
	c.element.Set("width", w)
	c.element.Set("height", h)
	c.width = w
	c.height = h
}

func (c *canvas) setSizeFromDocument() {
	resolution := 4
	c.width = js.Global().Get("window").Get("innerWidth").Int() / resolution
	c.height = js.Global().Get("window").Get("innerHeight").Int() / resolution
	c.setSize(c.width, c.height)
}

func (c *canvas) render() {
	data := make([]uint8, c.width*c.height*4)
	// loop over width and height pixels and add them to the imagedata array
	totalPixels := c.width * c.height
	for i := 0; i < totalPixels; i++ {
		postition := i * 4
		data[postition+0] = c.pixels[i/c.width][i%c.width].red
		data[postition+1] = c.pixels[i/c.width][i%c.width].green
		data[postition+2] = c.pixels[i/c.width][i%c.width].blue
		data[postition+3] = c.pixels[i/c.width][i%c.width].alpha
	}
	js.CopyBytesToJS(c.imagedata.Get("data"), data)
	c.ctx.Call("putImageData", c.imagedata, 0, 0)
}
