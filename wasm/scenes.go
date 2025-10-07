//go:build js && wasm
// +build js,wasm

package main

import (
	"fmt"
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
	translation := point3d{x: 0, y: 0, z: -10}
	// runtime_translation := point3d{x: 0, y: 0, z: 0.01}
	cam_transform_matrix.set_position(&translation)
	cam := camera{}
	cam.set_transform(cam_transform_matrix)
	cam.set_fov(6000)
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
	// var i float64 = 600
	for {
		canvasdata.set_background(pixel{red: 255, green: 255, blue: 200})
		mesh_data.draw(&canvasdata, &cam)
		// cent := mesh_data.get_center()
		// mesh_data.transform(0.30, 10, 0.60, &cent)
		// cam.set_fov(i)
		// i += 10
		// cam_transform_matrix.translate(&runtime_translation)
		// cam.set_transform(cam_transform_matrix)
		// fmt.Println(cam.transform)
		// triangle.draw(&canvasdata, &cam)
		// Rotate the mesh around its center
		// center := mesh_data.get_center()
		// fmt.Printf("Center of mesh: %v\n", center)
		// mesh_data.transform(0.01, 0.01, 0.01, &center)
		// Render the canvas
		canvasdata.render()
		time.Sleep(100 * time.Millisecond)
	}
}
