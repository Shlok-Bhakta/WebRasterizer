//go:build js && wasm
// +build js,wasm

package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"syscall/js"
)

const WORLD_SCALE = 2 // 1 unit in the world is 50 pixels on the screen

type mesh struct {
	triangles []triangle
}

// man the JS bridge feels so jank
func load_obj_from_browser() []string {
	object_array := js.Global().Get("allObjects")
	objects := make([]string, object_array.Length())
	for i := 0; i < object_array.Length(); i++ {
		objects[i] = object_array.Index(i).String()
	}
	return objects
}

// Extracts the point x y and z positions and stores them in a point3d array. (1 based indexing remember)
func extract_positions(obj_data string) []point3d {
	// split by newlines
	lines := strings.Split(obj_data, "\n")
	// loop through all lines and check to see if the first letter is v
	positions := []point3d{}
	for _, line := range lines {
		if strings.HasPrefix(line, "v ") {
			// split by spaces
			parts := strings.Fields(line)
			if len(parts) < 4 {
				continue // skip if not enough parts
			}
			// parse x, y, z as float64
			x, err := strconv.ParseFloat(parts[1], 64)
			if err != nil {
				fmt.Printf("Error parsing x line %s: %v\n", line, err)
				continue
			}
			y, err := strconv.ParseFloat(parts[2], 64)
			if err != nil {
				fmt.Printf("Error parsing y line %s: %v\n", line, err)
				continue
			}
			z, err := strconv.ParseFloat(parts[3], 64)
			if err != nil {
				fmt.Printf("Error parsing z line %s: %v\n", line, err)
				continue
			}
			positions = append(positions, point3d{x: x * WORLD_SCALE, y: y * WORLD_SCALE, z: z * WORLD_SCALE})
		}
	}
	return positions
}

func extract_point_indexes(obj_data string) [][]int {
	// split by newlines
	lines := strings.Split(obj_data, "\n")
	// loop through all lines and check to see if the first letter is f
	point_indexes := [][]int{}
	for _, line := range lines {
		if strings.HasPrefix(line, "f ") {
			// split by spaces
			parts := strings.Fields(line)
			// go through all parts in format p1/junk/junk ... and add them to the point indexes array
			indexes := []int{}
			for _, part := range parts[1:] { // skip the first part which is "f"
				// split by slashes
				subparts := strings.Split(part, "/")
				// just take the first one and add it to indexes so that the boi will be happy
				index, err := strconv.ParseInt(subparts[0], 10, 0)
				if err != nil {
					fmt.Printf("Error parsing index line %s: %v\n", line, err)
					continue
				}
				// convert to int and subtract 1 for 0-based indexing
				index = index - 1
				indexes = append(indexes, int(index))
			}
			point_indexes = append(point_indexes, indexes)
		}
	}
	return point_indexes
}

func enqueue(point_queue []point3d, point point3d) []point3d {
	if len(point_queue) == 3 {
		point_queue[1] = point
		return point_queue
	}
	point_queue = append(point_queue, point)
	fmt.Println(point_queue)
	fmt.Println(point)
	return point_queue
}

// Should extract all the data from the OBJ while also triangulating it with a basic algorithm
func parse_obj() mesh {
	objects := load_obj_from_browser()
	mesh_data := mesh{
		triangles: make([]triangle, 0),
	}
	// make a queue of points and we can push points in and make triangles from there
	point_queue := make([]point3d, 0, 3)
	// Support multiple objects :tada:
	for _, object := range objects {
		positions := extract_positions(object)
		for i := range positions {
			positions[i].z *= WORLD_SCALE
			positions[i].x *= WORLD_SCALE
			positions[i].y *= WORLD_SCALE
		}
		fmt.Printf("Loaded %d positions: %v\n", len(positions), positions)
		point_indexes := extract_point_indexes(object)
		fmt.Printf("Loaded %d point indexes: %v\n", len(point_indexes), point_indexes)
		// For each of the point_indexes arrays I think we can run the triangulation on them
		for _, face := range point_indexes {
			for _, point := range face {
				// queue up the next point
				point_queue = enqueue(point_queue, positions[point])
				// check if deque has 3 points
				if len(point_queue) == 3 {
					fmt.Println("yay ok so now we make tri!")
					// make a triangle with the points
					triangle := triangle{
						points: [3]point3d{
							point_queue[0],
							point_queue[1],
							point_queue[2],
						},
						color: make_random_pixel(),
					}
					fmt.Println(triangle)
					mesh_data.triangles = append(mesh_data.triangles, triangle)
				}

				// fmt.Printf("Enqueuing point: %v\n", positions[point])
			}
			// flush the queue so it is empty after the face has been triangulated
			point_queue = point_queue[:0]
		}
	}
	return mesh_data
}

func (m *mesh) draw(canvasdata *canvas, cam *camera) {
	for _, triangle := range m.triangles {
		triangle.draw(canvasdata, cam)
	}
}

func (m *mesh) get_center() point3d {
	if len(m.triangles) == 0 {
		return point3d{x: 0, y: 0, z: 0}
	}
	// Get the center of the first triangle
	center := m.triangles[0].get_center()
	for _, triangle := range m.triangles[1:] {
		center.x += triangle.get_center().x
		center.y += triangle.get_center().y
		center.z += triangle.get_center().z
	}
	center.x /= float64(len(m.triangles))
	center.y /= float64(len(m.triangles))
	center.z /= float64(len(m.triangles))
	return center
}

func (m *mesh) transform(roll float64, pitch float64, yaw float64, pivot *point3d) {
	// construct a transformation matrix for the roll, pitch, and yaw
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

	// multiply the transformation matrices together
	matrix := yaw_matrix.multiply(&pitch_matrix)
	matrix = matrix.multiply(&roll_matrix)

	// apply the transformation matrix to the triangle points with pivot
	for i := range m.triangles {
		for j := range m.triangles[i].points {
			// Translate to origin (subtract pivot)
			translated := point3d{
				x: m.triangles[i].points[j].x - pivot.x,
				y: m.triangles[i].points[j].y - pivot.y,
				z: m.triangles[i].points[j].z - pivot.z,
			}
			
			
			// Apply rotation
			rotated := translated.transform(&matrix)
			
			// Translate back (add pivot)
			m.triangles[i].points[j] = point3d{
				x: rotated.x + pivot.x,
				y: rotated.y + pivot.y,
				z: rotated.z + pivot.z,
			}
		}
	}
}
