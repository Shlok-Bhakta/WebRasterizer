//go:build js && wasm
// +build js,wasm

package main

type point struct {
	x int
	y int
}

func (p *point) distance(other point) float64 {
	dx := p.x - other.x
	dy := p.y - other.y
	return float64(dx*dx + dy*dy)
}
