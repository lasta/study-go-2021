---
title: "Day1"
date: 2021-09-17T09:33:37+09:00
draft: true
---

## 1日目
### 1章
#### 本文
##### P.6-5
```go
package main

func main() {
	var s, sep = string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
    }
	fmt.Println(s)
}
```

> なぜ2次のオーダーなのか

* `string` とはどういうデータ構造か?
  * `ptr` と `len` をもつ
  * Java と同様に string の値は不変
    * 文字列結合はあらたな string が作られる
  * コピーコストの合計は `(a + n) * n / 2 * x = n^2`
  * 文字列結合の頻度が高い場合は `+=` をつかわない
 
##### P.6-6
> Go ではインクリメントやデクリメントは式ではなく文 (`j = i++` はできない) のはどのようなモチベーションか

* `j = i++ + i++` について、他の言語では処理系等に依存して紛らわしい
* `i++` が文でも困らない

この2つからと思われる。

##### P.8-2
* ブランク識別子 `_` は変数ではないので値の参照はできない

##### P.8-4
> 省略変数宣言をパッケージレベルで使えないのはなぜか

構文解析を簡単化したかったからかも?

##### P.20-1
> チャネルから受信するメッセージの数が不定の場合

* サブルーチンが終了時に channel を閉じる
* WaitGroup (起動したルーチン) を使う
  * 8章 P.274 に出てくる
    * `sync.WaitGroup`, `wg.Add(1)`, `defer wg.Done()`, `wg.Wait()`

##### P.20.1
> 2つの Go-routine 内で `fmt.Printf()` などをしたら、出力が混ざることはあるのか

* Go も含めた多くの言語では、標準出力は排他処理なので、混ざらない。
* 裏を返すと、 Print 文は内部で lock を取ってしまうため、たとえデバッグプリントのつもりでも、並列処理上では挙動が変わり得る

##### P.23-2
> 読み込みの際は lock 不要では?

見かけ上は変数 `counter` は1つだが、実際はそうではない

```
            更新 (読み込み/書き込み)
var count <======================> func handler
       |
	   +-------------------------> func counter
	        読み込み
※ 実際はこうではない
```

* `http.HandlerFunc()` はそれぞれ goroutine を起動する
* 今どきは、 CPU のコアは1つではない
  * goroutine が同じ CPU コア上で動く保証はない
  * (特に Intel processor は) CPU と物理メモリは別もの
  * `Mutex.Lock()` を取らないと、物理メモリに書かれるタイミングは CPU 依存になる
    * L1/2 cache や register にのこったまま、書き込みが後回しにされる場合がある
    * 読み込み時も Lock を取らないと、 `counter()` はキャッシュ上の `count` を見に行ってしまうため、最新の値を取得できないかもしれない

結果、 `Mutex.Lock()` は単に排他かけているだけではなく、 CPU コア間の Line cache の同期を (結果として) 取ってくれている

*この仕様 (Memory model) は言語ごとに仕様が異なる*

* Go は仕様書が2つ用意されている
  * [Language Specification](https://golang.org/ref/spec)
  * [The Go Memory Model](https://golang.org/ref/mem)
    * 特に Happens Before
    * 8章で詳しく解説されている

#### 練習問題
##### ch01/ex12
> 定数で書いていたときに比べて、変数にしたらキャストを細かくいれる必要がある
> 定数と変数で、同じ数値なのに違いがあるのはなぜか

* Go には型の格上げがない
* Go において、定数は型なし整数
  * 変数の型の定義域におさまる場合、その変数の型にキャストされる

```go
// 右辺は浮動小数点表記だが、 int64 の範囲におさまるので許される
var x int64 = 1E10
```

* Go では、定数は 256 bit の精度が保証されている
  * 1 ZiB や 1 YiB の bit 数を保持する変数は作成できない
* 定数同士の演算は compile 時に計算される
  * Go が言語仕様でもつ基本データ型の精度を超える定数も演算できる

##### P.20
* goroutine にて channel への書き込み順序は保証されない
  * 書き込まれたものの取り出し順序は保証される
* goroutine は thread ではない

* OS Thread 数は `GOMAXPROCS` で指定できる
  * デフォルト値は論理コア数
* goroutine ごとに stack を持っている
* Go 側で適宜スケジューラに渡される

##### ch01/ex03
* benchmark は `testing.B` を用いる
  * `go test -bench .`
* コンパイラによる想定以上の最適化に注意
  * リテラルを用いた計算はコンパイル時に計算されてしまう可能性がある
  * 戻り値を使用していないと関数呼び出しすらされない可能性がある

##### ch01/ex06
* switch 文で書いたほうが可読性があがる

##### ch01/ex07 〜 ex09
* `net/httptest` を用いればテストできる

##### ch01/ex11
* production 向けに開発する場合は、 goroutine の同時起動数を制御する

### 2章
#### 本文
##### p.35-1
> アドレスを持たない値とは?

* リテラルと const くらい

##### P.35-7
> 関数がアドレスを返すのが推奨されるのはどのような場合?

* 戻り値が大きいオブジェクトの場合
* 使い分け自体は、明確には存在しない
  * 設計者次第


##### P.35-7
> Go では GC のパラメータチューニングはあるか

* どこまでヒープが大きくなったら (割合) GC するか、くらいしかない
  * `go doc debug.SetGCPercent`
  * これをチューニングすることは殆どない
    * メモリリークをつぶすのが先

##### P.48-2
* 同じディレクトリ内には同じパッケージ名しか使えない
  * `${package_name}_test` のみ ok
* module 名とディレクトリ名は無関係

* import のパス指定
  * `module名/パッケージ名`

* [利用金計算でも使える浮動小数点ライブラリ](https://engineering.mercari.com/blog/entry/20201203-basis-point/)

#### 練習問題
##### ch02/ex03 〜 ex05
* ベンチマークについて
  * 即値がコンパイル時に計算されたり、結果が捨てられる場合に処理が実行されないことに要注意

https://go-talks.appspot.com/github.com/tooru/gopl-popcount/popcount.slide#14

* Hacker's Delight, Figure 5-2 に、 PopCount のめっちゃ速いアルゴリズムが書いてある
  * og であれば `math.OnesCount46` 、Java であれば `Integer.bitCount` にもある

## 次回について
* 練習問題 3.6
  * supersampling
    * 2倍x2倍 で 4 pixel ずつ集める
* 練習問題 3.8
  * `bit.Rat` を用いると計算が終わらない
    * `iteration = 200` を 16 くらいにする
  * 「視覚的に」
    * 一部を拡大しないとわからない
* 練習問題 4.3 〜 4.7
  * 「スライス内での技法」
    * もらった slice の中身を入れ替えるのが期待値
    * 新しい slice を返すのではない
* 練習問題 4.13
  * Open Movie Database はアクセスキーの発行が必要

## その他 - go vendor
* `go get` した場合、 `$GOPATH` 配下に置かれるが、全て read only になる
```sh
go mod vendor
```
をすることで、 `vendor` ディレクトリが生成され、実装のコピーが配置される
* `vendor` ディレクトリがあれば、それを優先して見る
  * debug 時に便利
  * いらなくなったら消せば良い

