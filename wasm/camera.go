//go:build js && wasm
// +build js,wasm

package main

import "errors"

type camera struct {
	transform matrix4x4 // Camera location in world
	fov       float64   // Field of view in radians
}

func (c *camera) project_point(p point3d, canvas *canvas) (screen_point, error) {
	// Simple perspective projection
	view_matrix := c.get_view_matrix()
	camera_space := p.transform(&view_matrix)

	if camera_space.z <= 0.1 {
		// If the point is too close to the camera, we can't project it
		return screen_point{x: -1, y: -1}, errors.New("Point too close to Camera") // Invalid point
	}
	screen_x := (camera_space.x/camera_space.z)*c.fov + float64(canvas.width)/2
	screen_y := (camera_space.y/camera_space.z)*c.fov + float64(canvas.height)/2

	return screen_point{x: int(screen_x), y: int(screen_y)}, nil
}

func (c *camera) get_view_matrix() matrix4x4 {
	return c.transform.inverse()
}
