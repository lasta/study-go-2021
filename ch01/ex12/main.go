package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// fills palette with 1st color in this slice.
var palette = []color.Color{
	// background color := black
	color.Black,
	color.RGBA{
		R: 0xff,
		G: 0x00,
		B: 0x00,
		A: 0xff,
	},
	color.RGBA{
		R: 0x00,
		G: 0xff,
		B: 0x00,
		A: 0xff,
	},
	color.RGBA{
		R: 0x00,
		G: 0x00,
		B: 0xff,
		A: 0xff,
	},
	color.RGBA{
		R: 0xff,
		G: 0xff,
		B: 0x00,
		A: 0xff,
	},
}

const (
	blackIndex  = 0
	redIndex    = 1
	greenIndex  = 2
	blueIndex   = 3
	yellowIndex = 4
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			fmt.Fprint(w, err)
			return
		}

		lissajous(w, r.Form)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissajous(out io.Writer, params url.Values) {
	const (
		defaultCycles  = 5     // laps of oscillator x
		resolution     = 0.001 // 回転の分解能
		size           = 100   // キャンパスのの大きさ; [-size..+size]
		defaultNframes = 64    // animation frames
		defaultDelay   = 8     // 10ms 単位でのフレーム間の遅延
	)

	cycles := defaultCycles
	if len(params["cycles"]) > 0 {
		var err error
		if cycles, err = strconv.Atoi(params["cycles"][0]); err != nil {
			fmt.Fprintf(out, "failed to parse cycle: %v", params["cycles"][0])
			return
		}
	}

	nframes := defaultNframes
	if len(params["nframes"]) > 0 {
		var err error
		if nframes, err = strconv.Atoi(params["nframes"][0]); err != nil {
			fmt.Fprintf(out, "failed to parse cycle: %v", params["nframes"][0])
			return
		}
	}

	delay := defaultDelay
	if len(params["delay"]) > 0 {
		var err error
		if delay, err = strconv.Atoi(params["delay"][0]); err != nil {
			fmt.Fprintf(out, "failed to parse cycle: %v", params["delay"][0])
			return
		}
	}

	freq := rand.Float64() * 3.0 // 発振器 y の相対周波数
	animation := gif.GIF{LoopCount: nframes}
	phase := 0.0 // 位相差

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < float64(cycles*2)*math.Pi; t += resolution {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)

			if x >= 0 && y >= 0 {
				img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), redIndex)
				continue
			}
			if x >= 0 {
				img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), greenIndex)
				continue
			}
			if y >= 0 {
				img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blueIndex)
				continue
			}
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), yellowIndex)
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
