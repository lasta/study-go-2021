package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

// Lattice pixels around E
// a b c
// d e f
// g h i
type Lattice struct {
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

	writeImage("original.png", img)
	writeImage("antialiasing.png", antialias(img))
}

func writeImage(fileName string, img image.Image) {
	f, err := os.Create(fileName)
	defer f.Close()

	if err != nil {
		fmt.Printf("failed to create '%s'\n", fileName)
		return
	}
	err = png.Encode(f, img)
	if err != nil {
		fmt.Printf("failed to write into '%s'\n", fileName)
		return
	}
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			rgba := int(n) << 48 / iterations
			return color.YCbCr{
				Y:  uint8(rgba >> 32),
				Cr: uint8(rgba >> 16),
				Cb: uint8(rgba),
			}
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
func crop(origin *image.Image, x int, y int) (subImage Lattice) {
	subImage = Lattice{
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
