package primitives

import (
	"git.maze.io/go/math32"
	"github.com/go-gl/mathgl/mgl32"
)

// DRAW WITH: gl.DrawElements(gl.TRIANGLES, xSegments*ySegments*6, gl.UNSIGNED_INT, unsafe.Pointer(nil))
func Sphere(ySegments, xSegments int) (vertices, normals, tCoords []float32, indices []uint32) {

	for y := 0; y <= ySegments; y++ {
		for x := 0; x <= xSegments; x++ {
			xSegment := float32(x) / float32(xSegments)
			ySegment := float32(y) / float32(ySegments)

			xPos := float32(math32.Cos(xSegment*math32.Pi*2.0) * math32.Sin(ySegment*math32.Pi))
			yPos := float32(math32.Cos(ySegment * math32.Pi))
			zPos := float32(math32.Sin(xSegment*math32.Pi*2.0) * math32.Sin(ySegment*math32.Pi))

			vertices = append(vertices, xPos, yPos, zPos)
			xPos, yPos, zPos = mgl32.Vec3{xPos, yPos, zPos}.Normalize().Elem()
			normals = append(normals, xPos, yPos, zPos)
			tCoords = append(tCoords, xSegment, ySegment)
		}
	}
	for i := 0; i < ySegments; i++ {
		for j := 0; j < xSegments; j++ {
			a1 := uint32(i*(xSegments+1) + j)
			a2 := uint32((i+1)*(xSegments+1) + j)
			a3 := uint32((i+1)*(xSegments+1) + j + 1)
			b1 := uint32(i*(xSegments+1) + j)
			b2 := uint32((i+1)*(xSegments+1) + j + 1)
			b3 := uint32(i*(xSegments+1) + j + 1)
			indices = append(indices, a1, a2, a3, b1, b2, b3)
		}
	}
	return
}

// DRAW WITH: gl.DrawElements(gl.TRIANGLES, 3*xSegments* (2*zSegments-1), gl.UNSIGNED_INT, unsafe.Pointer(nil))
func Circle(xSegments, zSegments int) (vertices, normals, tCoords []float32, indices []uint32) {
	vertices = append(vertices, 0, 0, 0)
	normals = append(normals, 0, 1, 0)
	tCoords = append(tCoords, 0, 0)

	for x := 1; x <= xSegments; x++ {
		indices = append(indices, 0, uint32(x), uint32(x+1))
	}

	for z := 1; z <= zSegments; z++ {
		for x := 0; x <= xSegments; x++ {
			xSegment := float32(x) / float32(xSegments)
			zSegment := float32(z) / float32(zSegments)

			xPos := float32(math32.Cos(xSegment*math32.Pi*2.0) * zSegment)
			yPos := float32(0)
			zPos := float32(math32.Sin(xSegment*math32.Pi*2.0) * zSegment)

			vertices = append(vertices, xPos, yPos, zPos)
			normals = append(normals, 0, 1, 0)
			tCoords = append(tCoords, xSegment, zSegment)

			if x < xSegments && z < zSegments {
				indices = append(indices, uint32(
					x+1+((z-1)*(xSegments+1))),
					uint32(x+1+((z)*(xSegments+1))),
					uint32(x+2+((z-1)*(xSegments+1))),

					uint32(x+1+((z)*(xSegments+1))+1),
					uint32(x+2+((z-1)*(xSegments+1))),
					uint32(x+1+((z)*(xSegments+1))),
				)
			}
		}
	}
	return
}

// DRAW WITH: gl.DrawElements(gl.TRIANGLES, 6 * xSegments * (ySegments + 2*zSegments - 1), gl.UNSIGNED_INT, unsafe.Pointer(nil))
func Cylinder(ySegments, xSegments, zSegments int) (vertices, normals, tCoords []float32, indices []uint32) {

	verticesB, normalsB, tCoordsB, indicesB := Circle(xSegments, zSegments)
	vertices, normals, indices = verticesB, normalsB, indicesB
	indexOfset := uint32(len(vertices) / 3)
	tCoordsOfset := (float32(zSegments) / float32(ySegments+2*zSegments))

	for i, v := range tCoordsB {
		if i%2 == 1 {
			tCoords = append(tCoords, 1-tCoordsOfset*v)
		} else {
			tCoords = append(tCoords, v)
		}
	}

	for y := 0; y <= ySegments; y++ {
		for x := 0; x <= xSegments; x++ {
			xSegment := float32(x) / float32(xSegments)
			ySegment := float32(y) / float32(ySegments)

			xPos := float32(math32.Cos(xSegment * math32.Pi * 2.0))
			yPos := ySegment
			zPos := float32(math32.Sin(xSegment * math32.Pi * 2.0))

			vertices = append(vertices, xPos, yPos, zPos)
			xPos, yPos, zPos = mgl32.Vec3{xPos, 0, zPos}.Normalize().Elem()
			normals = append(normals, xPos, yPos, zPos)
			ySegment = (float32(y) + float32(zSegments)) / float32(ySegments+2*zSegments)
			tCoords = append(tCoords, xSegment, 1-ySegment)
		}
	}
	for i := 0; i < ySegments; i++ {
		for j := 0; j < xSegments; j++ {
			a1 := uint32(i*(xSegments+1)+j) + indexOfset
			a2 := uint32((i+1)*(xSegments+1)+j) + indexOfset
			a3 := uint32((i+1)*(xSegments+1)+j+1) + indexOfset
			b1 := uint32(i*(xSegments+1)+j) + indexOfset
			b2 := uint32((i+1)*(xSegments+1)+j+1) + indexOfset
			b3 := uint32(i*(xSegments+1)+j+1) + indexOfset
			indices = append(indices, a1, a2, a3, b1, b2, b3)
		}
	}

	indexOfset = uint32(len(vertices) / 3)
	for i := len(indicesB) - 1; i >= 0; i-- {
		indices = append(indices, indicesB[i]+indexOfset)
	}
	for i, v := range verticesB {
		if i%3 == 1 {
			vertices = append(vertices, 1)
		} else {
			vertices = append(vertices, v)
		}
	}
	for i, v := range tCoordsB {
		if i%2 == 1 {
			tCoords = append(tCoords, (float32(zSegments)/float32(ySegments))*v)
		} else {
			tCoords = append(tCoords, v)
		}
	}
	normals = append(normals, normalsB...)
	return
}

func Cone(ySegments, xSegments, zSegments int) (vertices, normals, tCoords []float32, indices []uint32) {

	verticesB, normalsB, tCoordsB, indicesB := Circle(xSegments, zSegments)
	vertices, normals, indices = verticesB, normalsB, indicesB
	indexOfset := uint32(len(vertices) / 3)
	tCoordsOfset := (float32(zSegments) / float32(ySegments+zSegments))

	for i, v := range tCoordsB {
		if i%2 == 1 {
			tCoords = append(tCoords, 1-tCoordsOfset*v)
		} else {
			tCoords = append(tCoords, v)
		}
	}

	for y := 0; y <= ySegments; y++ {
		for x := 0; x <= xSegments; x++ {
			xSegment := float32(x) / float32(xSegments)
			ySegment := float32(y) / float32(ySegments)

			xPos := float32(math32.Cos(xSegment*math32.Pi*2.0) * (1 - ySegment))
			yPos := ySegment
			zPos := float32(math32.Sin(xSegment*math32.Pi*2.0) * (1 - ySegment))

			vertices = append(vertices, xPos, yPos, zPos)
			normals = append(normals, xPos, 1, zPos)
			ySegment = (float32(y) + float32(zSegments)) / float32(ySegments+zSegments)
			tCoords = append(tCoords, xSegment, 1-ySegment)
		}
	}
	for i := 0; i < ySegments; i++ {
		for j := 0; j < xSegments; j++ {
			a1 := uint32(i*(xSegments+1)+j) + indexOfset
			a2 := uint32((i+1)*(xSegments+1)+j) + indexOfset
			a3 := uint32((i+1)*(xSegments+1)+j+1) + indexOfset
			b1 := uint32(i*(xSegments+1)+j) + indexOfset
			b2 := uint32((i+1)*(xSegments+1)+j+1) + indexOfset
			b3 := uint32(i*(xSegments+1)+j+1) + indexOfset
			indices = append(indices, a1, a2, a3, b1, b2, b3)
		}
	}
	return
}

// DRAW WITH: gl.DrawElements(gl.TRIANGLES, xSegments*ySegments*6, gl.UNSIGNED_INT, unsafe.Pointer(nil))
func Square(ySegments, xSegments int, segmentLength float32) (vertices, normals, tCoords []float32, indices []uint32) {
	yOfset := (float32(ySegments) * segmentLength) / 2
	xOfset := (float32(xSegments) * segmentLength) / 2
	for v := 0; v <= ySegments; v++ {
		for h := 0; h <= xSegments; h++ {
			vertices = append(vertices, (float32(h)*segmentLength)-xOfset, 0, (float32(v)*segmentLength)-yOfset)
			normals = append(normals, 0, 1, 0)
			tCoords = append(tCoords, (float32(h) * (1 / float32(ySegments))), (float32(v) * (1 / float32(xSegments))))
			if h < ySegments && v < xSegments {
				indices = append(indices, []uint32{
					uint32(h + (ySegments+1)*v),
					uint32(h + (ySegments+1)*v + (ySegments + 1)),
					uint32(h + (ySegments+1)*v + 1),

					uint32(h + (ySegments+1)*v + (ySegments + 1) + 1),
					uint32(h + (ySegments+1)*v + 1),
					uint32(h + (ySegments+1)*v + (ySegments + 1)),
				}...)
			}
		}
	}
	return
}

// DRAW WITH: gl.DrawElements(gl.TRIANGLES, 6*6, gl.UNSIGNED_INT, unsafe.Pointer(nil))
func Cube(X, Y, Z float32) (vertices, normals, tCoords []float32, indices []uint32) {
	vertices = []float32{
		//Front
		-X / 2, -Y / 2, Z / 2,
		X / 2, -Y / 2, Z / 2,
		-X / 2, Y / 2, Z / 2,
		X / 2, Y / 2, Z / 2,

		//Back
		-X / 2, -Y / 2, -Z / 2,
		X / 2, -Y / 2, -Z / 2,
		-X / 2, Y / 2, -Z / 2,
		X / 2, Y / 2, -Z / 2,

		//Right
		X / 2, -Y / 2, -Z / 2,
		X / 2, -Y / 2, Z / 2,
		X / 2, Y / 2, -Z / 2,
		X / 2, Y / 2, Z / 2,

		//Left
		-X / 2, -Y / 2, -Z / 2,
		-X / 2, -Y / 2, Z / 2,
		-X / 2, Y / 2, -Z / 2,
		-X / 2, Y / 2, Z / 2,

		//Top
		-X / 2, Y / 2, Z / 2,
		X / 2, Y / 2, Z / 2,
		-X / 2, Y / 2, -Z / 2,
		X / 2, Y / 2, -Z / 2,

		//Bottom
		-X / 2, -Y / 2, Z / 2,
		X / 2, -Y / 2, Z / 2,
		-X / 2, -Y / 2, -Z / 2,
		X / 2, -Y / 2, -Z / 2,
	}

	normalsTemp := []float32{
		0, 0, 1,
		0, 0, -1,
		1, 0, 0,
		-1, 0, 0,
		0, 1, 0,
		0, -1, 0,
	}

	for i := 0; i < 6; i++ {
		idx := uint32(3 * i)
		normals = append(normals, normalsTemp[idx], normalsTemp[idx+1], normalsTemp[idx+2], normalsTemp[idx], normalsTemp[idx+1], normalsTemp[idx+2], normalsTemp[idx], normalsTemp[idx+1], normalsTemp[idx+2], normalsTemp[idx], normalsTemp[idx+1], normalsTemp[idx+2])
		tCoords = append(tCoords, 0, 1, 1, 1, 0, 0, 1, 0)
		idx = uint32(4 * i)
		indices = append(indices, idx, idx+1, idx+2, idx+3, idx+2, idx+1)
	}
	return
}
