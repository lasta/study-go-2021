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

// NinePixel pixels around E
// a b c
// d e f
// g h i
type NinePixel struct {
	a color.Color
	b color.Color
	c color.Color
	d color.Color
	e color.Color
	f color.Color
	g color.Color
	h color.Color
	i color.Color
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
		width, height                  = 1024, 1024
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

func antialias(origin image.Image) image.Image {
	bound := origin.Bounds()
	width := bound.Size().X
	height := bound.Size().Y

	dest := image.NewRGBA(bound)

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			cropped := crop(&origin, x, y)

			// split pixel E into 4 pixels
			// a  | b       | c
			// ---+---------+---
			// d  | E00 E01 | f
			//    | E10 E11 |
			// ---+---------+---
			// g  | h       | i
			rgbaE00 := average(cropped.a, cropped.b, cropped.d, cropped.e)
			rgbaE01 := average(cropped.b, cropped.c, cropped.e, cropped.f)
			rgbaE10 := average(cropped.d, cropped.e, cropped.g, cropped.h)
			rgbaE11 := average(cropped.e, cropped.f, cropped.h, cropped.i)

			dest.Set(x, y, average(rgbaE00, rgbaE01, rgbaE10, rgbaE11))
		}
	}
	return dest
}

// crop image around x, y (point E)
//
// a b c
// d e f
// g h i
func crop(origin *image.Image, x int, y int) (subImage NinePixel) {
	subImage = NinePixel{
		a: (*origin).At(x-1, y-1),
		b: (*origin).At(x, y-1),
		c: (*origin).At(x+1, y-1),
		d: (*origin).At(x-1, y),
		e: (*origin).At(x, y),
		f: (*origin).At(x+1, y),
		g: (*origin).At(x-1, y+1),
		h: (*origin).At(x, y+1),
		i: (*origin).At(x+1, y+1),
	}
	bound := (*origin).Bounds()
	xMax := bound.Max.X - 1
	yMax := bound.Max.Y - 1

	// upper-left corner
	// x: unavailable pixel, o: available pixel
	// xxx
	// xoo
	// xoo
	if x == 0 && y == 0 {
		subImage.a = subImage.e
		subImage.b = subImage.e
		subImage.c = subImage.f
		subImage.d = subImage.e
		subImage.g = subImage.h
		return
	}

	// upper-right corner
	// xxx
	// oox
	// oox
	if x == xMax && y == 0 {
		subImage.a = subImage.d
		subImage.b = subImage.e
		subImage.c = subImage.e
		subImage.f = subImage.e
		subImage.i = subImage.h
		return
	}

	// lower-left corner
	// xoo
	// xoo
	// xxx
	if x == 0 && y == yMax {
		subImage.a = subImage.b
		subImage.d = subImage.e
		subImage.g = subImage.e
		subImage.h = subImage.e
		subImage.i = subImage.f
		return
	}

	// lower-right corner
	// oox
	// oox
	// xxx
	if x == xMax && y == yMax {
		subImage.c = subImage.b
		subImage.f = subImage.e
		subImage.g = subImage.d
		subImage.h = subImage.e
		subImage.i = subImage.e
		return
	}

	// not corner but left-bound
	if x == 0 {
		subImage.a = subImage.b
		subImage.d = subImage.e
		subImage.g = subImage.h
		return
	}

	// not corner but right-bound
	if x == xMax {
		subImage.c = subImage.b
		subImage.f = subImage.e
		subImage.i = subImage.h
	}

	// not corner but upper-bound
	if y == 0 {
		subImage.a = subImage.d
		subImage.b = subImage.e
		subImage.c = subImage.f
		return
	}

	// not corner but lower-bound
	if y == yMax {
		subImage.g = subImage.d
		subImage.h = subImage.e
		subImage.i = subImage.f
		return
	}

	// inner
	return
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
