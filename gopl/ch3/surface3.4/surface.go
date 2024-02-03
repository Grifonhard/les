package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
)

const (
	cells   = 100  // number of grid cells
	xyrange = 30.0 // x, y axis range (-xyrange..+xyrange)
	angle   = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

//func for coordinate calculation

func corner(i, j, width, height int) (float64, float64) {
	// find point (x,y) at corner of cell (i,j)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y) // compute surface height z

	// project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy)
	xyscale := float64(width) / 2 / xyrange // pixels per x or y unit
	zscale := float64(height) * 0.4         // pixels per z unit
	sx := float64(width)/2 + (x-y)*cos30*xyscale
	sy := float64(height)/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}

//func for post in response

func surface(w io.Writer, color string, width int, height int) {
	firstString := "<svg xmlns='http://www.w3.org/2000/svg' " +
		"style='stroke: grey; fill: " + color + "; stroke-width: 0.7' "
	fmt.Fprintf(w, firstString+"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j, width, height)
			bx, by := corner(i, j, width, height)
			cx, cy := corner(i, j+1, width, height)
			dx, dy := corner(i+1, j+1, width, height)
			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintln(w, "</svg>")
}

func draw(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	color := r.FormValue("color")
	widthS := r.FormValue("width")
	width, _ := strconv.Atoi(widthS)
	heightS := r.FormValue("height")
	height, _ := strconv.Atoi(heightS)
	w.Header().Set("Content-Type", "image/svg+xml")
	surface(w, color, width, height)
}

func main() {
	http.HandleFunc("/draw", draw)
	log.Fatal(http.ListenAndServe("localhost:8010", nil))
}

//http://localhost:8010/draw?color=white&width=600&height=600
