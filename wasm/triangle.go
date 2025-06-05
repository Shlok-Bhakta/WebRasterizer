//go:build js && wasm
// +build js,wasm

package main

type triangle struct {
	points [3]point
	color  pixel
}
