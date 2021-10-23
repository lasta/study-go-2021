package main

import "fmt"

type Point struct {
	X int
	Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	var w1 Wheel
	w1.X = 8
	w1.Y = 8
	w1.Radius = 5
	w1.Spokes = 20
	w2 := Wheel{
		Circle: Circle{
			Point:  Point{X: 8, Y: 8},
			Radius: 5,
		},
		Spokes: 20,
	}

	fmt.Printf("%#v\n", w1)
	fmt.Printf("%#v\n", w2)
}
