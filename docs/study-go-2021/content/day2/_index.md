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

#### 3.12
* `string(rune)` は成功するが、 `string(int)` は期待通りの挙動をしない

```go
package main

import "fmt"

func main() {
  fmt.Println('a')          // 97
  fmt.Println(string('a'))       // a
  fmt.Println(string(rune('a'))) // a
  fmt.Println(string(97))        // a
}
```

## 4章
### 補填
* 4.1
  * 実は `copy(dst []byte, src string) int` ができる
* slice はもともと pointer なので、 `&slice` のようにわたす旨味はほぼない

* map の繰り返しの順序は定義されていない
  * バグを見つけやすくするよう、意図的に順序がランダムになるようになっている
    * 実際には偏りがあるので注意
  * `fmt` パッケージに渡した場合はソートされるので固定

* map を set として使うこともある

### 練習問題
* 4.3 〜 4.7 : もとのスライスを書き換えるやつ
* もとのスライスとの判定
  * 長さを判定した上で、もとのスライスを subslice すれば良い

#### 4.4
* 最大公約数 (GCD) を用いると swap が固定サイズ、かつスッキリ実装できる

#### 4.6
* byte slice だが UTF-8 でエンコードされているものなので、 1 文字 1 byte とは限らない
  * `utf.Decode` などを用いて何 byte のものか判定しないとならない

#### 4.11
* sub-command ライブラリ: [cobra](https://github.com/spf13/cobra)
* CLI : [gocui](https://github.com/jroimartin/gocui)

#### 4.12
* 2000以上あるので並列実行で取得
* info API がある
  * https://xkcd.com/info.0.json

#### 4.13
* `ioutil.ReadAll` せずに `io.Copy` でファイルに書く