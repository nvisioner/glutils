package primitives

import (
	"image"
)

//GetSquare ...
func GetSquareTriangles(hTiles int, vTiles int, tileLengths float32) (vertices []float32, tCoords []float32, indices []uint32) {
	vOfset := (float32(vTiles) * tileLengths) / 2
	hOfset := (float32(hTiles) * tileLengths) / 2
	for v := 0; v <= vTiles; v++ {
		for h := 0; h <= hTiles; h++ {
			vertices = append(vertices, []float32{(float32(h) * tileLengths) - hOfset, 0, (float32(v) * tileLengths) - vOfset}...)
			tCoords = append(tCoords, []float32{(float32(h) * (1 / float32(hTiles))), (float32(v) * (1 / float32(vTiles)))}...)
			if h < hTiles && v < vTiles {
				indices = append(indices, []uint32{
					uint32(h + (hTiles+1)*v),
					uint32(h + (hTiles+1)*v + (hTiles + 1)),
					uint32(h + (hTiles+1)*v + 1),

					uint32(h + (hTiles+1)*v + (hTiles + 1) + 1),
					uint32(h + (hTiles+1)*v + 1),
					uint32(h + (hTiles+1)*v + (hTiles + 1)),
				}...)
			}
		}
	}
	return
}

func GetSquareStrip(hTiles int, vTiles int, tileLength float32) ([]float32, []float32, []uint32) {
	vertices := make([]float32, 0, (hTiles+1)*(vTiles+1)*3)
	tCoords := make([]float32, 0, (hTiles+1)*(vTiles+1)*3)
	indices := make([]uint32, 0, (hTiles+1)*vTiles*2)
	vOfset := (float32(vTiles) * tileLength) / 2
	hOfset := (float32(hTiles) * tileLength) / 2
	for v := 0; v <= vTiles; v++ {
		for h := 0; h <= hTiles; h++ {
			newH := h
			if v%2 != 0 {
				newH = hTiles - h
			}
			vertices = append(vertices, []float32{(float32(h) * tileLength) - hOfset, 0, (float32(v) * tileLength) - vOfset}...)
			tCoords = append(tCoords, []float32{(float32(h) * (1 / float32(hTiles))), (float32(v) * (1 / float32(vTiles)))}...)
			if v < vTiles {
				indices = append(indices, []uint32{
					uint32(newH + (hTiles+1)*v),
					uint32(newH + (hTiles+1)*(v+1)),
				}...)
			}
		}
	}
	return vertices, tCoords, indices
}

func GetSquareStripDisplaced(hTiles int, vTiles int, tileLength float32, img image.Image, magnitud float32) ([]float32, []float32, []uint32) {
	vertices := make([]float32, 0, (hTiles+1)*(vTiles+1)*3)
	tCoords := make([]float32, 0, (hTiles+1)*(vTiles+1)*3)
	indices := make([]uint32, 0, (hTiles+1)*vTiles*2)
	vOfset := (float32(vTiles) * tileLength) / 2
	hOfset := (float32(hTiles) * tileLength) / 2
	for v := 0; v <= vTiles; v++ {
		for h := 0; h <= hTiles; h++ {
			newH := h
			if v%2 != 0 {
				newH = hTiles - h
			}
			tCoordH, tCoordV := (float32(h) * (1 / float32(hTiles))), (float32(v) * (1 / float32(vTiles)))
			vertices = append(vertices, []float32{(float32(h) * tileLength) - hOfset, getDisplacemente(img, magnitud, tCoordH, tCoordV), (float32(v) * tileLength) - vOfset}...)
			tCoords = append(tCoords, []float32{tCoordH, tCoordV}...)
			if v < vTiles {
				indices = append(indices, []uint32{
					uint32(newH + (hTiles+1)*v),
					uint32(newH + (hTiles+1)*(v+1)),
				}...)
			}
		}
	}
	return vertices, tCoords, indices
}

func getDisplacemente(img image.Image, magnitud, h, v float32) float32 {
	imgBound := img.Bounds()
	imgWidth := float32(imgBound.Max.X)
	imgHeight := float32(imgBound.Max.Y)
	pos := img.At(int(imgWidth*h), int(imgHeight*v))
	r, g, b, _ := pos.RGBA()
	lum := (19595*r + 38470*g + 7471*b + 1<<15) >> 24
	return (magnitud * (float32(lum) / 255))
}
