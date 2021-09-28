# ch03/ex01

```go
func f(x float64, y float64) float64 {
    r := math.Hypot(x, y)
    return math.Sin(r) / r
}
```

この関数は、 `x = 0.0` かつ `y = 0.0` (原点) の場合、 `sin(0.0) / 0.0 = 0.0 / 0.0` となり、 `NaN` を返却する。
そのため、利用者が NaN 判定をする必要がある。

NaN を含みうることは、 Google Chrome の Developer mode で発見した。
