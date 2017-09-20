package main

import "fmt"
import "github.com/fogleman/ln/ln"

func main() {
	// add two tetrahedrons
	a1 := ln.Vector{ 1,  1,  1}
	a2 := ln.Vector{ 1, -1, -1}
	a3 := ln.Vector{-1,  1, -1}
	a4 := ln.Vector{-1, -1,  1}

	b1 := ln.Vector{-1, -1, -1}
	b2 := ln.Vector{-1,  1,  1}
	b3 := ln.Vector{ 1, -1,  1}
	b4 := ln.Vector{ 1,  1, -1}

	t1 := ln.NewTriangle(a1, a2, a3)
	t2 := ln.NewTriangle(a1, a2, a4)
	t3 := ln.NewTriangle(a1, a3, a4)
	t4 := ln.NewTriangle(a2, a3, a4)

	t5 := ln.NewTriangle(b1, b2, b3)
	t6 := ln.NewTriangle(b1, b2, b4)
	t7 := ln.NewTriangle(b1, b3, b4)
	t8 := ln.NewTriangle(b2, b3, b4)

	tetra := ln.NewMesh([]*ln.Triangle {t1, t2, t3, t4, t5, t6, t7, t8})

	for i := 0; i < 90; i += 2 {
		fmt.Println(i)
		m := ln.Rotate(ln.Vector{0, 0, -2}, ln.Radians(float64(i)))

		shape := ln.NewTransformedShape(tetra, m)

		render(shape, fmt.Sprintf("out%03d", i))
	}
}

func render(shape ln.Shape, rootname string) {
	// create a scene
	scene := ln.Scene{}

	scene.Add(shape)

	// define camera parameters
	eye := ln.Vector{6, 6, 6}    // camera position
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
