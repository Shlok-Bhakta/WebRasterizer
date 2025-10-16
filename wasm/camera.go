//go:build js && wasm
// +build js,wasm

package main

import (
	"errors"
	"math"
	"syscall/js"
	"fmt"
)
const SPEED float64 = 0.2
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
	screen_z := camera_space.z

	return screen_point{x: int(screen_x), y: int(screen_y), z: screen_z}, nil
}

func (c *camera) set_fov(f float64) {
	c.fov = (math.Pi / 180.0) * f
}

func (c *camera) set_transform(m matrix4x4) {
	c.transform = m
}

func (c *camera) get_view_matrix() matrix4x4 {
	return c.transform.inverse()
}

type InputState struct {
	W int
	A int
	S int
	D int
	SPACE int
	SHIFT int
	mouseX int
	mouseY int
	mouseDeltaX int
	mouseDeltaY int
}

func bool_to_int(b bool) int {
	if b {
		return 1
	}
	return 0
}

func (c *camera) js_transform()  {
	jsinputstate := js.Global().Get("window").Get("inputState")
	inputState := InputState{
		W: bool_to_int(jsinputstate.Get("W").Bool()),
		A: bool_to_int(jsinputstate.Get("A").Bool()),
		S: bool_to_int(jsinputstate.Get("S").Bool()),
		D: bool_to_int(jsinputstate.Get("D").Bool()),
		SPACE: bool_to_int(jsinputstate.Get("SPACE").Bool()),
		SHIFT: bool_to_int(jsinputstate.Get("SHIFT").Bool()),
		mouseX: jsinputstate.Get("mouseX").Int(),
		mouseY: jsinputstate.Get("mouseY").Int(),
		mouseDeltaX: jsinputstate.Get("mouseDeltaX").Int(),
		mouseDeltaY: jsinputstate.Get("mouseDeltaY").Int(),
	}
	p := point3d{
		x: float64(inputState.D - inputState.A) * SPEED , 
		y: float64(inputState.SHIFT - inputState.SPACE) * SPEED , 
		z: float64(inputState.W - inputState.S) * SPEED}
	c.transform.translate(&p)

	fmt.Println(inputState)
}