# 練習問題 1.11
[The Moz Top 500 Websites](https://moz.com/top500) から 500 ルートドメインのリスト (`top500Domains.csv`) を取得し、試行した。

```shell
go run fetchall.go $(tail -n +2 top500Domains.csv | cut -f 2 -d ',' | cut -f 2 -d '"' | tr '\n' ' ')
```

## `top500Domains.csv` のファイル仕様
* カンマ `,` 区切りの CSV ファイル
* 要素はすべてダブルクオーテーション `"` で囲まれている
* 第2カラムがルートドメイン

## 試行結果 (1)
ファイルディスクリプタの上限が 256 であったため、 `socket: too many open files` が大量に発生した。
上限を 1024 に変更の上、再度試行した。

```shell
$ ulimit -a
(略)
-n: file descriptors                256
$ ulimit -n 1024
$ ulimit -a
(略)
-n: file descriptors                1024
```

## 試行結果 (2)

取得結果は以下表のようになった。

| 結果                       | 件数 |
|----------------------------|------|
| 成功                       | 469  |
| `no such host`             | 25   |
| `connection reset by peer` | 5    |
| `connection refused`       | 1    |

`connection reset by peer` となったページの影響で、プログラムが終了するまで 606.74 秒かかった。

```shell
$ go run fetchall.go $(tail -n +2 top500Domains.csv | cut -f 2 -d ',' | cut -f 2 -d '"' | tr '\n' ' ')
(各ページの取得結果)
606.74s elapsed
```

各ルーチンの完了の同期が取られることは期待通りだが、 GET リクエストで 600 秒以上待つことは多くの場合期待通りではないため、レスポンスの取得の際に適当なタイムアウト値を設定すべきであると考えた。
