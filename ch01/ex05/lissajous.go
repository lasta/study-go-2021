package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

// fills palette with 1st color in this slice.
var palette = []color.Color{
	// background color := black
	color.Black,
	// lissajous line color := green
	color.RGBA{
		R: 0x00,
		G: 0xff,
		B: 0x00,
		A: 0xff,
	},
}

const (
	blackIndex = 0 // 1st color in `palette`
	greenIndex = 1 // 2nd color in `palette`
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles     = 5     // laps of oscillator x
		resolution = 0.001 // 回転の分解能
		size       = 100   // キャンパスのの大きさ; [-size..+size]
		nframes    = 64    // animation frames
		delay      = 8     // 10ms 単位でのフレーム間の遅延
	)

	freq := rand.Float64() * 3.0 // 発振器 y の相対周波数
	animation := gif.GIF{LoopCount: nframes}
	phase := 0.0 // 位相差

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < cycles*2*math.Pi; t += resolution {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)

			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), greenIndex)
		}

		phase += 0.1
		animation.Delay = append(animation.Delay, delay)
		animation.Image = append(animation.Image, img)
	}

	err := gif.EncodeAll(out, &animation)
	if err != nil {
		fmt.Printf("failed to encode. cause: %v", err)
		return
	}
}
