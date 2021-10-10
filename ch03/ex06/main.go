package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

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

	writeImage("mandelbrot.png", img)
	writeImage("supersampled.png", supersample(img))
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

// cf: assets/fig.png
func supersample(original image.Image) image.Image {
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
