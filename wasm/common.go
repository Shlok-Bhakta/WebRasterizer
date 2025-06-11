//go:build js && wasm
// +build js,wasm

package main

// dot product of two points
func dot(a point, b point) float64 {
	return a.x*b.x + a.y*b.y
}
