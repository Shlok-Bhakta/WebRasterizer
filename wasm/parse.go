//go:build js && wasm
// +build js,wasm

package main

import (
	"fmt"
	"strconv"
	"strings"
	"syscall/js"
)

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
			positions = append(positions, point3d{x: x, y: y, z: z})
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

// Should extract all the data from the OBJ while also triangulating it with a basic algorithm
func parse_obj() mesh {
	objects := load_obj_from_browser()
	mesh_data := mesh{
		triangles: make([]triangle, len(objects)),
	}
	for _, object := range objects {
		positions := extract_positions(object)
		fmt.Printf("Loaded %d positions: %v\n", len(positions), positions)
		point_indexes := extract_point_indexes(object)
		fmt.Printf("Loaded %d point indexes: %v\n", len(point_indexes), point_indexes)

		// triangle := triangle{
		// 	points: [3]point3d{
		// 		{x: 0, y: 0, z: 0},
		// 		{x: 0, y: 0, z: 0},
		// 		{x: 0, y: 0, z: 0},
		// 	},
		// 	color: make_random_pixel(),
		// }
		// append(mesh_data.triangles, triangle)
	}
	return mesh_data
}
