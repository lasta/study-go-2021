# ch04.ex02 実行結果
```
# default: SHA256
$ echo 'a' | go run main.go
ca978112ca1bbdcafac231b39a23dc4da786eff8147c4e72b9807785afee48bb

$ echo 'a' | go run main.go -hashfunc=SHA256
ca978112ca1bbdcafac231b39a23dc4da786eff8147c4e72b9807785afee48bb

$ echo 'a' | go run main.go -hashfunc=SHA384
54a59b9f22b0b80880d8427e548b7c23abd873486e1f035dce9cd697e85175033caa88e6d57bc35efae0b5afd3145f31

$ echo 'a' | go run main.go -hashfunc=SHA512
1f40fc92da241694750979ee6cf582f2d5d7d28e18335de05abc54d0560e0f5302860c652bf08d560252aa5e74210546f369fbbbce8c12cfc7957b2652fe9a75

# unknown hash function
$ echo 'a' | go run main.go -hashfunc=hoge
unkown hash function: hoge
exit status 1
```
