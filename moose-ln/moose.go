package main

import "os"
import "fmt"
import "github.com/fogleman/ln/ln"

func main() {
	// add moose stl
	moose, err := ln.LoadBinarySTL("moose.stl")
	if err != nil {
		panic(err)
	}

	// moose.UnitCube()
	// zoom in a bit

	_ = os.Mkdir("out", 0775)

	for i := 0; i < 360; i += 2 {
		m := ln.Rotate(ln.Vector{0, 0, -2}, ln.Radians(float64(i)))

		shape := ln.NewTransformedShape(moose, m)

		render(shape, fmt.Sprintf("out/out%03d", i))
	}
}

func render(shape ln.Shape, rootname string) {
	// create a scene
	scene := ln.Scene{}

	scene.Add(shape)

	// define camera parameters
	// eye := ln.Vector{6, 6, 6}    // camera position
	eye := ln.Vector{0, -6, 2}    // camera position
	center := ln.Vector{0, 0, 0} // camera looks at
	up := ln.Vector{0, 0, 1}     // up direction

	// define rendering parameters
	width := 512.0   // rendered width
	height := 512.0  // rendered height
	fovy := 20.0     // vertical field of view, degrees
	znear := 0.1     // near z plane
	zfar := 20.0     // far z plane
	step := 0.01     // how finely to chop the paths for visibility testing

	// compute 2D paths that depict the 3D scene
	paths := scene.Render(eye, center, up, width, height, fovy, znear, zfar, step)

	// render the paths in an image
	paths.WriteToPNG(rootname + ".png", width, height)

	// save the paths as an svg
	paths.WriteToSVG(rootname + ".svg", width, height)
}
