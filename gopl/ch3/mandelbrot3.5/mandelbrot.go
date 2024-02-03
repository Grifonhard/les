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
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img)
}

func mandelbrot(z complex128) color.RGBA {
	const iterations = 200
	//const contrast = 15
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
