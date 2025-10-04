//go:build js && wasm
// +build js,wasm

package main

import (
	"fmt"
	"math"
	"time"
)

// "fmt"

// func triangle3d() {
// 	canvasdata := canvas{}
// 	canvasdata.init()
// 	// draw a triangle
// 	t := triangle{
// 		points: [3]point3d{
// 			{x: 50 * 3, y: 5 * 3, z: 0},
// 			{x: 30 * 3, y: 5 * 3, z: 0},
// 			{x: 5 * 3, y: 30 * 3, z: 10},
// 		},
// 		color: make_random_pixel(),
// 	}
// 	for {
// 		for i := 0; i < canvasdata.height; i++ {
// 			for j := 0; j < canvasdata.width; j++ {
// 				canvasdata.pixels[i][j] = pixel{red: 255, green: 255, blue: 200}
// 			}
// 		}
// 		t.draw(&canvasdata)
// 		t_center := t.get_center()
// 		t.transform(0.1, 0.00, 0.00, &t_center)
// 		canvasdata.render()
// 		time.Sleep(time.Millisecond * 100)
// 	}
// }

func cube() {
	canvasdata := canvas{}
	canvasdata.init()
	cam_transform_matrix := identity()
	translation := point3d{x: 0, y: 0, z: -3}
	cam_transform_matrix.translate(&translation)
	cam := camera{
		transform: cam_transform_matrix,
		fov:       600 * (math.Pi / 180.0),
	}

	// Load the cube from OBJ file
	mesh_data := parse_obj()
	fmt.Printf("Loaded mesh with %d triangles & mesh data %v\n", len(mesh_data.triangles), mesh_data)
	// render each triangle
	// triangle := triangle{
	// 	points: [3]point3d{
	// 		{x: 0, y: 0, z: 0},
	// 		{x: 1, y: 0, z: 0},
	// 		{x: 0, y: 1, z: 0},
	// 	},
	// 	color: make_random_pixel(),
	// }
	for {
		canvasdata.setBackground(pixel{red: 255, green: 255, blue: 200})
		mesh_data.draw(&canvasdata, &cam)
		// Rotate the mesh around its center
		center := mesh_data.get_center()
		// fmt.Printf("Center of mesh: %v\n", center)
		mesh_data.transform(0.01, 0.01, 0.01, &center)
		// Render the canvas
		canvasdata.render()
		time.Sleep(100 * time.Millisecond)
	}
}
