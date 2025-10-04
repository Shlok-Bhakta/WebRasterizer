//go:build js && wasm
// +build js,wasm

package main

import (
	"fmt"
	"syscall/js"
)

type canvas struct {
	width     int
	height    int
	scale     uint8
	pixels    [][]pixel
	imagedata js.Value
	ctx       js.Value
	element   js.Value
}

func (c *canvas) init() {
	c.element = js.Global().Get("document").Call("getElementById", "canvas")
	c.ctx = js.Global().Get("document").Call("getElementById", "canvas").Call("getContext", "2d")
	c.scale = 4 // Set the scale factor for the canvas
	c.setSizeFromDocument()
	// c.setSize(700, 500)
	// Setup resize listener
	js.Global().Call("addEventListener", "resize", js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		c.setSizeFromDocument()
		c.render()
		return nil
	}))
}

func (c *canvas) setSize(w int, h int) {
	c.element.Set("width", w)
	c.element.Set("height", h)
	c.width = w
	c.height = h
}

func (c *canvas) setSizeFromDocument() {
	c.width = js.Global().Get("window").Get("innerWidth").Int() / int(c.scale)
	c.height = js.Global().Get("window").Get("innerHeight").Int() / int(c.scale)
	c.imagedata = c.ctx.Call("createImageData", js.ValueOf(c.width), js.ValueOf(c.height))
	if c.pixels == nil {
		c.pixels = make([][]pixel, c.height)
		for i := range c.pixels {
			c.pixels[i] = make([]pixel, c.width)
		}
	} else {
		newpixels := make([][]pixel, c.height)
		for i := range newpixels {
			newpixels[i] = make([]pixel, c.width)
		}
		for i := 0; i < c.height; i++ {
			for j := 0; j < c.width; j++ {
				newpixels[i][j] = c.pixels[i][j]
			}
		}
		c.pixels = newpixels
	}
	fmt.Printf("Setting canvas size to %d x %d\n", c.width, c.height)
	c.setSize(c.width, c.height)
}

func (c *canvas) setBackground(p pixel) {
	for i := 0; i < c.height; i++ {
		for j := 0; j < c.width; j++ {
			c.pixels[i][j] = p
		}
	}
}

func (c *canvas) render() {
	data := make([]uint8, c.width*c.height*4)
	fmt.Printf("Rendering canvas with width: %d, height: %d, scale: %d\n", c.width, c.height, c.scale)
	// loop over width and height pixels and add them to the imagedata array
	totalPixels := c.width * c.height
	for i := 0; i < totalPixels; i++ {
		postition := i * 4
		data[postition+0] = c.pixels[i/c.width][i%c.width].red
		data[postition+1] = c.pixels[i/c.width][i%c.width].green
		data[postition+2] = c.pixels[i/c.width][i%c.width].blue
		data[postition+3] = 255
	}
	js.CopyBytesToJS(c.imagedata.Get("data"), data)
	c.ctx.Call("putImageData", c.imagedata, 0, 0)
}

// takes a number from 0 to 1 and maps it to the width of the canvas
func (c *canvas) mapWidth(x float64) int {
	if x < 0 {
		x = 0
	}
	if x > 1 {
		x = 1
	}
	portion := int(float64(c.width) * x)
	if portion < 0 {
		portion = 0
	}
	if portion > c.width {
		portion = c.width
	}
	return portion
}

// takes a number from 0 to 1 and maps it to the height of the canvas
func (c *canvas) mapHeight(y float64) int {
	if y < 0 {
		y = 0
	}
	if y > 1 {
		y = 1
	}
	portion := int(float64(c.height) * y)
	if portion < 0 {
		portion = 0
	}
	if portion > c.height {
		portion = c.height
	}
	return portion
}
