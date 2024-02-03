package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			color := superSampling(x, y, z)
			img.Set(px, py, color)
		}
	}
	png.Encode(os.Stdout, img)
}

func mandelbrot(z complex128) color.RGBA {
	const iterations = 200
	var v complex128
	var R, G, B, A uint8
	A = 255
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			if n < 2 {
				R = 255
				B = 255
				G = 255
			} else if n < 5 {
				R = 255
				G = 255
			} else if n < 10 {
				B = 255
				G = 255
			} else if n < 20 {
				R = 255
				B = 255
			} else if n < 40 {
				R = 255
			} else if n < 80 {
				G = 255
			} else {
				B = 255
			}
			return color.RGBA{R, G, B, A}
		}
	}
	return color.RGBA{R, G, B, A}
}

func superSampling(x, y float64, z complex128) color.RGBA {
	colorPixel := mandelbrot(z)
	colorPixel1 := mandelbrot(complex((x + 1/256), (y + 1/256)))
	colorPixel2 := mandelbrot(complex((x + 1/256), (y - 1/256)))
	colorPixel3 := mandelbrot(complex((x - 1/256), (y + 1/256)))
	colorPixel4 := mandelbrot(complex((x - 1/256), (y - 1/256)))
	Rsample := uint8((uint16(colorPixel1.R) + uint16(colorPixel2.R) + uint16(colorPixel3.R) + uint16(colorPixel4.R)) / 4)
	Gsample := uint8((uint16(colorPixel1.G) + uint16(colorPixel2.G) + uint16(colorPixel3.G) + uint16(colorPixel4.G)) / 4)
	Bsample := uint8((uint16(colorPixel1.B) + uint16(colorPixel2.B) + uint16(colorPixel3.B) + uint16(colorPixel4.B)) / 4)
	var color color.RGBA
	color.R = uint8((uint16(colorPixel.R) + uint16(Rsample)) / 2)
	color.G = uint8((uint16(colorPixel.G) + uint16(Gsample)) / 2)
	color.B = uint8((uint16(colorPixel.B) + uint16(Bsample)) / 2)
	color.A = 255
	return color
}
