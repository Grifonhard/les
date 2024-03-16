package main

import (
	"image"
	"image/color"
	"image/png"
	"math/big"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 8000, 8000
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := big.NewFloat(float64(py)/float64(height)*float64(ymax-ymin) + float64(ymin))
		for px := 0; px < width; px++ {
			x := big.NewFloat(float64(px)/float64(width)*float64(xmax-xmin) + float64(xmin))
			img.Set(px, py, mandelbrot(x, y))
		}
	}
	png.Encode(os.Stdout, img)
}

func mandelbrot(x, y *big.Float) color.RGBA {
	const iterations = 200
	//const contrast = 15
	var R, G, B, A uint8
	A = 255

	a, _ := x.Float64()
	b, _ := y.Float64()
	var v complex128 = complex(a, b)
	var vx *big.Float = x
	var vy *big.Float = y
	//вычисление в несколько шагов с помощью требует несколько переменных
	var vx1, vx2, vx3, vy1, vy2, vy3 *big.Float
	var vxf, vyf float64

	for n := uint8(0); n < iterations; n++ {
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
		//вычисление действительной части
		vx1.Mul(vx, vx)
		vx2.Mul(vy, vy)
		vx2.Neg(vx2)
		vx3.Add(vx1, vx2)
		vx = vx3.Add(vx3, x)
		//вычисление мнимой части
		vy1.Mul(vx, vy)
		vy2.Mul(big.NewFloat(2.0), vy1)
		vy = vy3.Add(vy2, y)
		//приведение к комплексному числу
		vxf, _ = vx.Float64()
		vyf, _ = vy.Float64()
		v = complex(vxf, vyf)
	}
	return color.RGBA{R, G, B, A}
}
