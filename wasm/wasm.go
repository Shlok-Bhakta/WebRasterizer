//go:build js && wasm
// +build js,wasm

package main

func main() {
	InitBrowser()
	select {} // Keep the program running
}
