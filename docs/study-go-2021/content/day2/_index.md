---
title: "Day2"
date: 2021-10-15T09:33:37+09:00
draft: false
---

# 2日目
## 3章
### 補填
* 3.6
  * 本章では `bytes.Buffer` を使っているが、文字列については `strings.Builder` が Go 1.10 から追加されたのでそれを利用する

### 練習問題
#### 3.1
* 「有限ではない」なので、無限大も除外すべき
  * この問題では発生しないが

* label も使える (P.287 で登場) 

```go
package main

func main() {
	for {
	innerfor:
		for {
			continue innerfor
		}
    }
}
```

#### 3.8
* `big.Rat` は iteration　数を調整しないと全然計算が終わらない
* `big.Float`, `big.Rat` はメモリを確保する
  * `-benchmem` でメモリ関連情報を取得できる