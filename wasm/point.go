//go:build js && wasm
// +build js,wasm

package main

import "math"

type point struct {
	x float64
	y float64
}

type screen_point struct {
	x int
	y int
}

func (p *point) toScreen(canvas *canvas) screen_point {
	// convert the point to screen coordinates
	screen_p := screen_point{
		x: int(p.x),
		y: int(p.y),
	}
	return screen_p
}

func (p *point) distance(other point) float64 {
	dx := p.x - other.x
	dy := p.y - other.y
	return float64(dx*dx + dy*dy)
}

// rotate rotates the point around a pivot by a given angle in radians.
func (p *point) rotate(angle float64, pivot *point) {
	sin := math.Sin(angle)
	cos := math.Cos(angle)

	// send point to origin
	p.x -= pivot.x
	p.y -= pivot.y

	newx := p.x*cos - p.y*sin
	newy := p.x*sin + p.y*cos

	p.x = newx + pivot.x
	p.y = newy + pivot.y
}
