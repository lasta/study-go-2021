package main

import (
	"bytes"
	"fmt"
	"image/color"
	"io"
	"log"
	"math"
	"net/http"
	"net/url"
	"strconv"
)

const (
	defaultWidth  = 600
	defaultHeight = 320
	cells         = 100         // number of cells
	xyrange       = 30.0        // range (-xyrange..+xyrange)
	angle         = math.Pi / 6 // angle of x axis and y axis (= 30°)
)

type DrawingParameter struct {
	height      int
	width       int
	peakColor   color.RGBA
	valleyColor color.RGBA
}

var sin30 = math.Sin(angle)
var cos30 = math.Cos(angle)

var red = color.RGBA{R: 0xff}
var green = color.RGBA{G: 0xff}
var blue = color.RGBA{B: 0xff}
var black = color.RGBA{}

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if err := request.ParseForm(); err != nil {
			fmt.Fprint(writer, err)
			return
		}
		writer.Header().Set("Content-Type", "image/svg+xml")
		drawingParameter := parseRequestParameters(request.Form)
		draw(writer, drawingParameter)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func parseRequestParameters(urlParams url.Values) (params DrawingParameter) {
	params = DrawingParameter{
		height:      defaultHeight,
		width:       defaultWidth,
		peakColor:   color.RGBA{},
		valleyColor: color.RGBA{},
	}

	if width, err := strconv.Atoi(urlParams.Get("width")); err == nil {
		params.width = width
	}

	if height, err := strconv.Atoi(urlParams.Get("height")); err == nil {
		params.height = height
	}

	params.peakColor = selectColor(urlParams.Get("peak-color"))
	params.valleyColor = selectColor(urlParams.Get("valley-color"))

	return params
}

func selectColor(colorName string) color.RGBA {
	switch colorName {
	case "red":
		return red
	case "blue":
		return blue
	case "green":
		return green
	}
	return black
}

func draw(writer io.Writer, parameter DrawingParameter) {
	width := parameter.width
	height := parameter.height
	peak := parameter.peakColor
	valley := parameter.valleyColor

	xyscale := float64(width) / 2.0 / xyrange
	zscale := float64(height) * 0.4

	var buffer bytes.Buffer
	buffer.WriteString(
		fmt.Sprintf("<svg "+
			"xmlns='http://www.w3.org/2000/svg' "+
			"style='stroke: grey; fill: white; stroke-width: 0.7' "+
			"width='%d' height='%d'>", width, height),
	)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, _ := corner(i+1, j, width, height, xyscale, zscale, peak, valley)
			bx, by, _ := corner(i, j, width, height, xyscale, zscale , peak, valley)
			cx, cy, _ := corner(i, j+1, width, height, xyscale, zscale, peak, valley)
			dx, dy, c := corner(i+1, j+1, width, height, xyscale, zscale, peak, valley)

			if Any([]bool{
				math.IsNaN(ax), math.IsNaN(ay),
				math.IsNaN(bx), math.IsNaN(by),
				math.IsNaN(cx), math.IsNaN(cy),
				math.IsNaN(dx), math.IsNaN(dy),
			}) {
				continue
			}

			buffer.WriteString(
				fmt.Sprintf("<polygon points='%g,%g,%g,%g %g,%g,%g,%g' stroke='rgb(%d,%d,%d)'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy, c.R, c.G, c.B),
			)
		}
	}

	buffer.WriteString(fmt.Sprintln("</svg>"))
	buffer.WriteTo(writer)
}

func corner(
	i int,
	j int,
	width int,
	height int,
	xyscale float64,
	zscale float64,
	peakColor color.RGBA,
	valleyColor color.RGBA,
) (sx float64, sy float64, c color.RGBA) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z, c := f(x, y, peakColor, valleyColor)

	sx = float64(width)/2 + (x-y)*cos30*xyscale
	sy = float64(height)/2 + (x+y)*sin30*xyscale - z*zscale
	return
}

func f(x float64, y float64, peakColor color.RGBA, valleyColor color.RGBA) (float64, color.RGBA) {
	r := math.Hypot(x, y)
	sinR := math.Sin(r)
	c := color.RGBA{}
	switch {
	case sinR > 0:
		c = adjustColor(sinR, peakColor)
	case sinR < 0:
		c = adjustColor(-sinR, valleyColor)
	}
	return sinR / r, c
}

func adjustColor(v float64, rgb color.RGBA) (c color.RGBA) {
	c = color.RGBA{}
	if v > 1.0 {
		return
	}
	if v < -1.0 {
		return
	}
	if v == 0.0 {
		return
	}

	colorElemValue := uint8(v * float64(0xff))
	switch rgb {
	case red:
		c.R = colorElemValue
		return
	case green:
		c.G = colorElemValue
		return
	case blue:
		c.B = colorElemValue
		return
	}
	return
}

func Any(values []bool) bool {
	for _, value := range values {
		if value {
			return true
		}
	}
	return false
}
