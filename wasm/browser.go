//go:build js && wasm
// +build js,wasm

package main

import (
	"syscall/js"
)

func InitBrowser() {
	js.Global().Get("console").Call("log", "Go WASM Canvas loaded!")
	canvas := js.Global().Get("document").Call("getElementById", "canvas")
	// Set the Canvas css to be position absolute and left 0, top 0
	ctx := canvas.Call("getContext", "2d")
	setCanvasSize(canvas)
	render(ctx, canvas)
	// Listen for window resize
	js.Global().Call("addEventListener", "resize", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		setCanvasSize(canvas)
		render(ctx, canvas)
		return nil
	}))
}

func setCanvasSize(canvas js.Value) {
	width := js.Global().Get("window").Get("innerWidth").Int()
	height := js.Global().Get("window").Get("innerHeight").Int()

	canvas.Set("width", width)
	canvas.Set("height", height)

	js.Global().Get("console").Call("log", "Canvas size:", width, "x", height)
}

func render(ctx js.Value, canvas js.Value) {
	width := canvas.Get("width").Int()
	height := canvas.Get("height").Int()

	// Set fill color to red
	ctx.Set("fillStyle", "purple")

	// Fill the entire canvas
	ctx.Call("fillRect", 0, 0, width, height)

	js.Global().Get("console").Call("log", "Red background drawn!")
}
