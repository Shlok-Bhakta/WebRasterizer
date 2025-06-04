package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	fmt.Println("WebAssembly loaded!")

	// Keep the program running
	select {}
}

func hello(this js.Value, args []js.Value) interface{} {
	fmt.Println("Hello from WebAssembly!")
	return "Hello from Go WASM!"
}

func init() {
	// Register the function to be callable from JavaScript
	js.Global().Set("goHello", js.FuncOf(hello))
}
