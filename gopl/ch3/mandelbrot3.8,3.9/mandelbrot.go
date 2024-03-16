// аккураси 3 и 4 не работает, случается паника. runtime error: invalid memory address or nil pointer dereference
package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/big"
	"math/cmplx"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/mandelbrot", draw)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func draw(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	w.Header().Set("Content-Type", "image/png")
	var scale int = 1024
	var accur string = "2"
	scaleS := r.FormValue("scale")
	scale, _ = strconv.Atoi(scaleS)
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
	)
	width, height := scale, scale
	accur = r.FormValue("accuracy")
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		switch accur {
		case "1":
			y := float32(py)/float32(height)*float32(ymax-ymin) + float32(ymin)
			for px := 0; px < width; px++ {
				x := float32(px)/float32(width)*float32(xmax-xmin) + float32(xmin)
				var z complex64 = complex(x, y)
				img.Set(px, py, mandelbrot1(z))
			}
		case "2":
			y := float64(py)/float64(height)*float64(ymax-ymin) + float64(ymin)
			for px := 0; px < width; px++ {
				x := float64(px)/float64(width)*float64(xmax-xmin) + float64(xmin)
				var z complex128 = complex(x, y)
				img.Set(px, py, mandelbrot2(z))
			}
		case "3":
			y := big.NewFloat(float64(py)/float64(height)*float64(ymax-ymin) + float64(ymin))
			for px := 0; px < width; px++ {
				x := big.NewFloat(float64(px)/float64(width)*float64(xmax-xmin) + float64(xmin))
				img.Set(px, py, mandelbrot3(x, y))
			}
		case "4":
			//ya, yb, xa, xb - числитель и знаменатель у и х
			ya := int64(py*(ymax-ymin) + ymin*height)
			yb := int64(height)
			y := big.NewRat(ya, yb)
			for px := 0; px < width; px++ {
				xa := int64(px*(xmax-xmin) + xmin*width)
				xb := int64(width)
				x := big.NewRat(xa, xb)
				img.Set(px, py, mandelbrot4(x, y))
			}
		}
	}
	png.Encode(w, img)
}

func mandelbrot1(z complex64) color.Color {
	const iterations = 200
	const contrast = 15
	var v complex64
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(complex128(v)) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func mandelbrot2(z complex128) color.Color {
	const iterations = 200
	const contrast = 15
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func mandelbrot3(x, y *big.Float) color.Color {
	const iterations = 200
	const contrast = 15
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
			return color.Gray{255 - contrast*n}
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
	return color.Black
}

func mandelbrot4(x, y *big.Rat) color.Color {
	const iterations = 200
	const contrast = 15
	a, _ := x.Float64()
	b, _ := y.Float64()
	var v complex128 = complex(a, b)
	var vx *big.Rat = x
	var vy *big.Rat = y
	//вычисление в несколько шагов с помощью требует несколько переменных
	var vx1, vx2, vx3, vy1, vy2, vy3 *big.Rat
	var vxf, vyf float64
	for n := uint8(0); n < iterations; n++ {
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
		//вычисление действительной части
		vx1.Mul(vx, vx)
		vx2.Mul(vy, vy)
		vx2.Neg(vx2)
		vx3.Add(vx1, vx2)
		vx = vx3.Add(vx3, x)
		//вычисление мнимой части
		vy1.Mul(vx, vy)
		vy2.Mul(big.NewRat(2, 1), vy1)
		vy = vy3.Add(vy2, y)
		//приведение к комплексному числу
		vxf, _ = vx.Float64()
		vyf, _ = vy.Float64()
		v = complex(vxf, vyf)
	}
	return color.Black
}
