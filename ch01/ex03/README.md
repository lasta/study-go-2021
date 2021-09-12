# 練習問題 1.3

## 実装内容

### 非効率と思われるもの

```go
package main

func joinInefficiently(values []string) string {
	var joined string
	var separator string
	for i := 0; i < len(values); i++ {
		joined += separator + values[i]
		separator = " "
	}
	return joined
}
```

要素1個を join するたびに新しい strings のオブジェクトを生成しているため、非効率と思われる。

### `strings.Join` を使ったもの

```go
package main

import "strings"

func joinEfficiently(values []string) string {
	return strings.Join(values, " ")
}
```

## 実行時間の差の計測

### 前提条件

1. 要素数 1,000 の string の配列を用意
2. 1.の配列の join を 10,000 回実行するのにかかった時間を、非効率版、 `strings.Join` 版のそれぞれで計測
3. 2.を10回実行

### 結果

| 試行回数 | 非効率版 [s] | `strings.Join` 版 [s] | 実行時間の差 [s] |
|:--------:|-------------:|----------------------:|-----------------:|
|     1    |         1.93 |                  0.07 |             1.86 |
|     2    |         1.89 |                  0.07 |             1.82 |
|     3    |          1.9 |                  0.08 |             1.82 |
|     4    |         1.92 |                  0.08 |             1.84 |
|     5    |         1.91 |                  0.08 |             1.83 |
|     6    |         1.92 |                  0.08 |             1.84 |
|     7    |         1.93 |                  0.07 |             1.85 |
|     8    |         1.92 |                  0.07 |             1.85 |
|     9    |         1.94 |                  0.08 |             1.86 |
|    10    |         2.16 |                  0.09 |             2.07 |
|   平均    |        1.942 |                 0.077 |            1.864 |

非効率と思われていたものは、 `strings.Join` の約25倍の実行時間がかかることがわかった。
