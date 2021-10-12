package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math/cmplx"
	"net/http"
	"net/url"
	"strconv"
)

type DrawingParameter struct {
	xCenter      float64
	yCenter      float64
	zoom         float64
	antialiasing bool
}

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if err := request.ParseForm(); err != nil {
			fmt.Fprint(writer, err)
			return
		}
		drawingParameter := parseRequestParameters(request.Form)
		err := draw(writer, drawingParameter)
		if err != nil {
			fmt.Fprint(writer, err)
			return
		}
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func parseRequestParameters(urlParams url.Values) (params DrawingParameter) {
	params = DrawingParameter{
		xCenter:      0.0,
		yCenter:      0.0,
		zoom:         1.0,
		antialiasing: false,
	}

	if xCenter, err := strconv.ParseFloat(urlParams.Get("x"), 64); err == nil {
		params.xCenter = xCenter
	}
	if yCenter, err := strconv.ParseFloat(urlParams.Get("y"), 64); err == nil {
		params.yCenter = yCenter
	}
	if zoom, err := strconv.ParseFloat(urlParams.Get("zoom"), 64); err == nil {
		params.zoom = zoom
	}
	if antialiasing, err := strconv.ParseBool(urlParams.Get("antialiasing")); err == nil {
		params.antialiasing = antialiasing
	}
	return params
}

func draw(writer io.Writer, params DrawingParameter) error {
	const (
		xmin1x, xmax1x, ymin1x, ymax1x = -2, 2, -2, 2
		width, height = 1024, 1024
	)
	xmin := xmin1x / params.zoom
	xmax := xmax1x / params.zoom
	ymin := ymin1x / params.zoom
	ymax := ymax1x / params.zoom

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for py := 0; py < height; py++ {
		y := (float64(py)-params.yCenter)/height*(ymax-ymin) + ymin

		for px := 0; px < width; px++ {
			x := (float64(px)-params.xCenter)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, fractal(z))
		}
	}

	if !params.antialiasing {
		err := png.Encode(writer, img)
		if err != nil {
			return err
		}
	}
	err := png.Encode(writer, antialias(img))
	if err != nil {
		return err
	}
	return nil
}

// f(z) = z^4 - 1
//
// Newton-Raphson method
// z[n + 1] = z[n] - f(z[n]) / f'(z[n])
//          = z[n] - (z[n]^4 - 1) / (4 * z[n]^3)
//          = (3 / 4) * z[n] + 1 / (4 * z[n]^3)
//          = (3 * z[n] + 1 / z[n]^3) / 4
// z[0] = -1
func fractal(z complex128) color.Color {
	const iterations = 64
	const contrast = 4
	const epsilon = 1e-6

	for iteration := uint8(0); iteration < iterations; iteration++ {
		z = (3*z + 1/(z*z*z)) / 4

		if cmplx.Abs((1+0i)-z) < epsilon {
			return color.RGBA{R: 0xff - contrast*iteration, A: 0xc0}
		}
		if cmplx.Abs((-1+0i)-z) < epsilon {
			return color.RGBA{G: 0xff - contrast*iteration, A: 0xc0}
		}
		if cmplx.Abs((0+1i)-z) < epsilon {
			return color.RGBA{B: 0xff - contrast*iteration, A: 0xc0}
		}
		if cmplx.Abs((0-1i)-z) < epsilon {
			return color.RGBA{R: 0xff - contrast*iteration, G: 0xff - contrast*iteration, A: 0xc0}
		}
	}
	return color.Black
}

func antialias(original image.Image) image.Image {
	bound := original.Bounds()
	width := bound.Size().X
	height := bound.Size().Y

	dest := image.NewRGBA(bound)

	for x := 0; x < width-1; x++ {
		for y := 0; y < height-1; y++ {
			// A B C
			// D E F
			// G H I
			rgbaA := original.At(x, y)
			rgbaB := original.At(x, y)
			rgbaC := original.At(x+1, y)
			rgbaD := original.At(x, y)
			rgbaE := original.At(x, y)
			rgbaF := original.At(x+1, y)
			rgbaG := original.At(x, y+1)
			rgbaH := original.At(x, y+1)
			rgbaI := original.At(x+1, y+1)

			if x > 0 && y > 0 {
				rgbaA = original.At(x-1, y-1)
			}
			if x > 0 {
				rgbaD = original.At(x-1, y)
				rgbaG = original.At(x-1, y+1)
			}
			if y > 0 {
				rgbaB = original.At(x, y-1)
				rgbaC = original.At(x+1, y-1)
			}

			// split pixel E into 4 pixels
			// A  | B       | C
			// ---+---------+---
			// D  | E00 E01 | F
			//    | E10 E11 |
			// ---+---------+---
			// G  | H       | I
			rgbaE00 := average(rgbaA, rgbaB, rgbaD, rgbaE)
			rgbaE01 := average(rgbaB, rgbaC, rgbaE, rgbaF)
			rgbaE10 := average(rgbaD, rgbaE, rgbaG, rgbaH)
			rgbaE11 := average(rgbaE, rgbaF, rgbaH, rgbaI)

			dest.Set(x, y, average(rgbaE00, rgbaE01, rgbaE10, rgbaE11))
		}
	}
	return dest
}

func average(pixels ...color.Color) color.Color {
	var rs []int
	var gs []int
	var bs []int
	var as []int

	for _, pixel := range pixels {
		pixel.RGBA()
		r, g, b, a := pixel.RGBA()
		rs = append(rs, int(uint8(r)))
		gs = append(gs, int(uint8(g)))
		bs = append(bs, int(uint8(b)))
		as = append(as, int(uint8(a)))
	}

	var size = len(pixels)
	return color.RGBA{
		R: uint8(sum(rs) / size),
		G: uint8(sum(gs) / size),
		B: uint8(sum(bs) / size),
		A: uint8(sum(as) / size),
	}
}

func sum(elements []int) (sum int) {
	for _, element := range elements {
		sum += element
	}
	return
}
