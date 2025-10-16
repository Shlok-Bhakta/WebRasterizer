//go:build js && wasm
// +build js,wasm

package main

import "math"

type matrix4x4 [4][4]float64

func (m *matrix4x4) multiply(o *matrix4x4) matrix4x4 {
	return matrix4x4{
		{
			m[0][0]*o[0][0] + m[0][1]*o[1][0] + m[0][2]*o[2][0] + m[0][3]*o[3][0],
			m[0][0]*o[0][1] + m[0][1]*o[1][1] + m[0][2]*o[2][1] + m[0][3]*o[3][1],
			m[0][0]*o[0][2] + m[0][1]*o[1][2] + m[0][2]*o[2][2] + m[0][3]*o[3][2],
			m[0][0]*o[0][3] + m[0][1]*o[1][3] + m[0][2]*o[2][3] + m[0][3]*o[3][3],
		},
		{
			m[1][0]*o[0][0] + m[1][1]*o[1][0] + m[1][2]*o[2][0] + m[1][3]*o[3][0],
			m[1][0]*o[0][1] + m[1][1]*o[1][1] + m[1][2]*o[2][1] + m[1][3]*o[3][1],
			m[1][0]*o[0][2] + m[1][1]*o[1][2] + m[1][2]*o[2][2] + m[1][3]*o[3][2],
			m[1][0]*o[0][3] + m[1][1]*o[1][3] + m[1][2]*o[2][3] + m[1][3]*o[3][3],
		},
		{
			m[2][0]*o[0][0] + m[2][1]*o[1][0] + m[2][2]*o[2][0] + m[2][3]*o[3][0],
			m[2][0]*o[0][1] + m[2][1]*o[1][1] + m[2][2]*o[2][1] + m[2][3]*o[3][1],
			m[2][0]*o[0][2] + m[2][1]*o[1][2] + m[2][2]*o[2][2] + m[2][3]*o[3][2],
			m[2][0]*o[0][3] + m[2][1]*o[1][3] + m[2][2]*o[2][3] + m[2][3]*o[3][3],
		},
		{
			m[3][0]*o[0][0] + m[3][1]*o[1][0] + m[3][2]*o[2][0] + m[3][3]*o[3][0],
			m[3][0]*o[0][1] + m[3][1]*o[1][1] + m[3][2]*o[2][1] + m[3][3]*o[3][1],
			m[3][0]*o[0][2] + m[3][1]*o[1][2] + m[3][2]*o[2][2] + m[3][3]*o[3][2],
			m[3][0]*o[0][3] + m[3][1]*o[1][3] + m[3][2]*o[2][3] + m[3][3]*o[3][3],
		},
	}
}

func (m *matrix4x4) add(o *matrix4x4) matrix4x4 {
	return matrix4x4{
		{
			m[0][0] + o[0][0], m[0][1] + o[0][1], m[0][2] + o[0][2], m[0][3] + o[0][3],
		},
		{
			m[1][0] + o[1][0], m[1][1] + o[1][1], m[1][2] + o[1][2], m[1][3] + o[1][3],
		},
		{
			m[2][0] + o[2][0], m[2][1] + o[2][1], m[2][2] + o[2][2], m[2][3] + o[2][3],
		},
		{
			m[3][0] + o[3][0], m[3][1] + o[3][1], m[3][2] + o[3][2], m[3][3] + o[3][3],
		},
	}
}

// Bit of a hack, but it works when scale is uniform
func (m *matrix4x4) inverse() matrix4x4 {
	return matrix4x4{
		{m[0][0], m[1][0], m[2][0], -(m[0][0]*m[0][3] + m[1][0]*m[1][3] + m[2][0]*m[2][3])},
		{m[0][1], m[1][1], m[2][1], -(m[0][1]*m[0][3] + m[1][1]*m[1][3] + m[2][1]*m[2][3])},
		{m[0][2], m[1][2], m[2][2], -(m[0][2]*m[0][3] + m[1][2]*m[1][3] + m[2][2]*m[2][3])},
		{0, 0, 0, 1},
	}
}

// sets the position of the matrix to the given point
func (m *matrix4x4) set_position(p *point3d) {
	m[0][3] = p.x
	m[1][3] = p.y
	m[2][3] = p.z
}

// sets the position of the matrix to the given point
func (m *matrix4x4) translate(p *point3d) {
	m[0][3] += p.x
	m[1][3] += p.y
	m[2][3] += p.z
}

func (m *matrix4x4) rotate(roll float64, pitch float64, yaw float64) {
	rotation := make_rotation_matrix(roll, pitch, yaw)
	position := point3d{x: m[0][3], y: m[1][3], z: m[2][3]}
	*m = rotation.multiply(m)
	m[0][3] = position.x
	m[1][3] = position.y
	m[2][3] = position.z
}

func identity() matrix4x4 {
	return matrix4x4{
		{1, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}
}

func make_rotation_matrix(roll float64, pitch float64, yaw float64) matrix4x4 {
	roll_matrix := matrix4x4{
		{1, 0, 0, 0},
		{0, math.Cos(roll), -math.Sin(roll), 0},
		{0, math.Sin(roll), math.Cos(roll), 0},
		{0, 0, 0, 1},
	}
	pitch_matrix := matrix4x4{
		{math.Cos(pitch), 0, -math.Sin(pitch), 0},
		{0, 1, 0, 0},
		{math.Sin(pitch), 0, math.Cos(pitch), 0},
		{0, 0, 0, 1},
	}
	yaw_matrix := matrix4x4{
		{math.Cos(yaw), -math.Sin(yaw), 0, 0},
		{math.Sin(yaw), math.Cos(yaw), 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}
	ry := yaw_matrix.multiply(&pitch_matrix)
	return ry.multiply(&roll_matrix)
}

