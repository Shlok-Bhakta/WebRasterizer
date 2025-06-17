//go:build js && wasm
// +build js,wasm

package main

import "math"

type point3d struct {
	x, y, z float64
}

// type point2d struct {
// 	x, y float64
// }

type screen_point struct {
	x, y int
}

func (p *point3d) distance(other point3d) float64 {
	dx := p.x - other.x
	dy := p.y - other.y
	return float64(dx*dx + dy*dy)
}

// rotate rotates the point around a pivot by a given angle in radians.
func (p *point3d) rotate(angle float64, pivot *point3d) {
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

func (p *point3d) transform(matrix *matrix4x4) point3d {
	return point3d{
		x: matrix[0][0]*p.x + matrix[0][1]*p.y + matrix[0][2]*p.z + matrix[0][3],
		y: matrix[1][0]*p.x + matrix[1][1]*p.y + matrix[1][2]*p.z + matrix[1][3],
		z: matrix[2][0]*p.x + matrix[2][1]*p.y + matrix[2][2]*p.z + matrix[2][3],
	}
}
