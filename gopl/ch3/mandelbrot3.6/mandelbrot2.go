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
		for px := 0; px < width; px++ {
			y := float64(py)/(height)*(ymax-ymin) + ymin
			x := float64(px)/(width)*(xmax-xmin) + xmin
			x1 := x - (xmax-xmin)/(width*2)
			x2 := x + (xmax-xmin)/(width*2)
			y1 := y - (ymax-ymin)/(height*2)
			y2 := y + (ymax-ymin)/(height*2)
			z1 := complex(x1, y1)
			z2 := complex(x1, y2)
			z3 := complex(x2, y1)
			z4 := complex(x2, y2)
			color1 := mandelbrot(z1)
			color2 := mandelbrot(z2)
			color3 := mandelbrot(z3)
			color4 := mandelbrot(z4)
			var color color.RGBA
			color.R = uint8((uint16(color1.R) + uint16(color2.R) + uint16(color3.R) + uint16(color4.R)) / 4)
			color.G = uint8((uint16(color1.G) + uint16(color2.G) + uint16(color3.G) + uint16(color4.G)) / 4)
			color.B = uint8((uint16(color1.B) + uint16(color2.B) + uint16(color3.B) + uint16(color4.B)) / 4)
			color.A = 255
			img.Set(px, py, color)
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
