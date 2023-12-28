package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
	"image/gif"
	"math/rand"
	"io"
	"image"
	"image/color"
	"math"
)


var palette = []color.Color{color.Black, color.RGBA{0x00, 0xff, 0x00, 0xff}}

const (
	blackIndex = 0
	greenIndex = 1
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	cycles, err := strconv.Atoi(r.URL.Query().Get("cycles"))
	if err != nil {
		fmt.Println(err)
	}
	lissajous(w, cycles)
}

func lissajous(out io.Writer, cycles int) {
	const (
		res = 0.001
		size = 100
		nframes = 64
		delay = 8
	)
	source := rand.NewSource(time.Now().UTC().UnixNano())
	sRand := rand.New(source)
	freq := sRand.Float64() * 3.0
	anim :=gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i:=0 ; i <nframes; i++ {
		rect := image.Rect(0, 0, 2*size + 1, 2*size + 1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size + int (x*size + 0.5), size + int(y*size+0.5), greenIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}	
	gif.EncodeAll(out, &anim)
}