# ch03/ex02

## saddle point

```go
package main

const (
	xyrange = 30.0 // range (-xyrange..+xyrange)
)

func saddlePoint(x float64, y float64) float64 {
	r := x*x - y*y
	return r / (xyrange * xyrange)
}
```

![saddle_point.svg](./saddle_point/saddle_point.svg)
