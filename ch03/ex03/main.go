package main

import (
	"fmt"
	"image/color"
	"math"
)

const (
	width, height = 600, 320            // size of campus (pixel number)
	cells         = 100                 // number of cells
	xyrange       = 30.0                // range (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x unit and y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x axis and y axis (= 30°)
)

var sin30 = math.Sin(angle)
var cos30 = math.Cos(angle)

func main() {
	fmt.Printf("<svg "+
		"xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, _ := corner(i+1, j)
			bx, by, _ := corner(i, j)
			cx, cy, _ := corner(i, j+1)
			dx, dy, c := corner(i+1, j+1)

			if Any([]bool{
				math.IsNaN(ax), math.IsNaN(ay),
				math.IsNaN(bx), math.IsNaN(by),
				math.IsNaN(cx), math.IsNaN(cy),
				math.IsNaN(dx), math.IsNaN(dy),
			}) {
				continue
			}

			fmt.Printf("<polygon points='%g,%g,%g,%g %g,%g,%g,%g' stroke='rgb(%d,%d,%d)'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, c.R, c.G, c.B)
		}
	}
	fmt.Println("</svg>")
}

func corner(i int, j int) (float64, float64, color.RGBA) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z, c := f(x, y)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, c
}

func f(x float64, y float64) (float64, color.RGBA) {
	r := math.Hypot(x, y)
	sinR := math.Sin(r)
	c := color.RGBA{}
	switch {
	case sinR > 0:
		c.R = uint8(sinR * float64(0xff))
	case sinR < 0:
		c.B = uint8(-sinR * float64(0xff))
	}
	return sinR / r, c
}

func Any(values []bool) bool {
	for _, value := range values {
		if value {
			return true
		}
	}
	return false
}
