package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/big"
	"math/cmplx"
	"net/http"
	"os"
	"strconv"

	"golang.org/x/text/width"
)

func main() {
	http.HandleFunc("/mandelbrot", draw)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type 

func draw(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	scaleS := r.FormValue(scale)
	scale, _ := strconv.Atoi(scaleS)
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
	)
	width, height := scale, scale
	accur := r.FormValue(accuracy)
	var funcM* func 
	switch accur {
	case "1":
	case "2":
	case "3":
	case "4":
	}
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

func mandelbrot1(z complex64) color.Color {
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

func mandelbrot2(z complex128) color.Color {

}

func mandelbrot3(z big.Float) color.Color {

}

func mandelbrot4(z big.Rat) color.color {

}
