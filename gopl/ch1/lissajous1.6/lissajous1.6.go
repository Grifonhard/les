package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{color.Black, color.RGBA{0x00, 0xff, 0x00, 0xff}, color.RGBA{0xff, 0x00, 0x00, 0xff}, color.RGBA{0x00, 0x00, 0xff, 0xff} }

const (
	blackIndex = 0
	greenIndex = 1
	redIndex = 2
	blueIndex = 3
)

func main() {
	f, err := os.Create("a.gif")
	if err != nil {
    	panic(err)
	}
	defer f.Close()
	lissajous(f)
}

func lissajous(out io.Writer) {
	const (
		cycles = 5
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
		
		changeIndex := 0
		index := greenIndex 
		
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)			
			img.SetColorIndex(size + int (x*size + 0.5), size + int(y*size+0.5), uint8(index))
			
			
			if changeIndex > 10 {
				index = redIndex
			}
			if changeIndex > 20 {
				index = blueIndex
			}
			if changeIndex >30 {
				changeIndex = 0
				index = greenIndex
			}
			changeIndex++

			
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}	
	gif.EncodeAll(out, &anim)
}