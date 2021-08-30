---
title: "Chpt 01. Tutor"
date: 2021-08-29T13:33:46+09:00
---

## 基本構文

* Go では全てのインデックスの指定は半開 (half-open)
    * `s[m:n]` : `0 ≤ m ≤ n ≤ len(s)`
* `os.Args[0]` はコマンド自身の名前
    * Python などと同じ

### for 文

```java
public class Echo1 {

    public static void main(String[] args) {
        // Java では String の初期値は null
        String s = "";
        String sep = " ";
        for (int i = 1; i < args.length; i++) {
            s += sep + args[i];
        }
        System.out.println(s);
    }
}
```

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	// Go では string の初期値は ""
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
```

無限ループは条件を指定しなければ ok

```go
package main

func main() {
	for {
		// 無限ループ
	}
}
```


