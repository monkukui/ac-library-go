# ac-library-go

AtCoder Library(ACL) の Go 移植版です。
ACL で提供されている全てのライブラリを移植することを目指します。

ACL については以下をご覧ください。
- [AtCoder Library (ACL) - AtCoder](https://atcoder.jp/posts/517) 
- [AtCoder Library - Codeforces](https://codeforces.com/blog/entry/82400)

パッケージ構成は、暫定で以下のように ACL の 1 ファイルに 1 パッケージを対応させる形とします。

```
.
├── dsu
│   └── dsu.go
├── fenwicktree
│   └── fenwicktree.go
├── go.mod
└── internal
    └── math
        └── math.go
```

## Contribute 方法
- ALC で提供されているアルゴリズム・データ構造を実装して，PR を送ってください。
- 実装したいものがあるときは、該当の issue で宣言してからはじめてもらえると、衝突を防げます。
- その他お気軽に issue や draft PR も歓迎です。
