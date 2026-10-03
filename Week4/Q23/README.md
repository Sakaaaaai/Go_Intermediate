# 問23｜リクエストとレスポンス

## 目標
HTTPリクエストのクエリパラメータを取得し、その内容に応じて異なるレスポンスを返す方法を学びます。

## 今回の演習で学ぶこと
この問題では、「URLにパラメータを付けてサーバーに情報を渡す」という基本的なWeb通信の仕組みを学びます。

---

### クエリパラメータとは何か

URLにパラメータを付けてサーバーに情報を渡すのはWebの基本です。
Googleの検索も、クエリパラメータを使っています：

```
https://www.google.com/search?q=golang&hl=ja
                               ↑ これがクエリパラメータ
```

`?` 以降の部分がクエリパラメータで、`キー=値` の形式です。複数ある場合は `&` でつなぎます。

---

**1. クエリパラメータとは**

URLの `?` 以降に `キー=値` の形式で付加される情報です。

```
http://localhost:8080/?name=Alice&age=30
                       ↑              ↑
                  name=Alice     age=30（2つのパラメータ）
```

---

**2. クエリパラメータの取得**

```go
func myHandler(w http.ResponseWriter, r *http.Request) {
    // r.URL.Query() でクエリパラメータ全体のマップを取得
    // .Get("キー名") で値を取得（存在しなければ空文字列を返す）
    name := r.URL.Query().Get("name")
    age := r.URL.Query().Get("age")  // 文字列として返ってくる
    
    if name == "" {
        fmt.Fprint(w, "Hello, Guest!")
    } else {
        fmt.Fprintf(w, "Hello, %s! (Age: %s)", name, age)
    }
}
```

---

**3. r.URL.Query() の詳細**

`r.URL.Query()` は `url.Values` 型を返します。これは `map[string][]string` です（1つのキーに複数の値を持てる）。

```go
// URLが /search?tag=go&tag=web の場合
tags := r.URL.Query()["tag"]  // ["go", "web"] という文字列スライス
tag := r.URL.Query().Get("tag") // "go"（最初の値だけ）
```

---

**4. リクエストから取得できる情報まとめ**

```go
r.Method          // "GET", "POST", "DELETE" などのHTTPメソッド
r.URL.Path        // "/users/123" などのパス部分
r.URL.Query()     // クエリパラメータ（マップ）
r.URL.RawQuery    // "name=Alice&age=30"（生の文字列）

r.Header.Get("Authorization")  // ヘッダーの値
r.Header.Get("Content-Type")   // リクエストのContent-Type

r.Body            // リクエストボディ（POSTなどのデータ）
r.ContentLength   // ボディのバイト数
r.RemoteAddr      // クライアントのIPアドレスとポート
```

---

**5. ステータスコードの設定**

```go
// デフォルトは200 OK（`w.WriteHeader` を呼ばなくてもOK）

// エラーレスポンスを返す
http.Error(w, "Not Found", http.StatusNotFound) // 404 とメッセージを同時に
http.Error(w, "Bad Request", http.StatusBadRequest) // 400

// ステータスコードだけ設定（ボディは別途）
w.WriteHeader(http.StatusCreated) // 201
fmt.Fprint(w, "作成しました")
```

> **重要**: `w.WriteHeader()` は一度呼ぶと後から変更できません。ヘッダー（`w.Header().Set(...)`）の設定は必ず `WriteHeader` より前に行ってください。

---

**6. クエリパラメータの典型的な使い方**

```
検索のフィルター: GET /products?category=electronics&price_max=10000
ページネーション: GET /users?page=2&per_page=20
ソート順の指定: GET /items?sort=name&order=desc
```

---

## 問題文
`main.go` の `___` の部分を埋めてプログラムを完成させてください。

1. `"/"` にアクセスしたときのハンドラーを登録してください。
2. `r.URL.Query().Get("name")` でパラメータ `name` を取得してください。
3. `name` が空ならば `"Hello, Guest!"` を、値があれば `"Hello, <name>!"` を返してください。

### 実行方法
```bash
cd /project/workspace/Week4/Q23
go run main.go
```

### テストでの確認
```bash
cd /project/workspace/Week4/Q23
go test main.go main_test.go
```
## ヒント
- `r.URL.Query().Get("key")` は値がなければ空文字列 `""` を返す
- `fmt.Fprintf(w, "Hello, %s!", name)` で変数を埋め込んで書き込める
- 空チェックは `if name == ""` でシンプルに書ける

## 達成基準
- [ ] クエリパラメータを取得できる
- [ ] パラメータの有無で分岐した応答が返せる
- [ ] `fmt.Fprintf` でフォーマット付きレスポンスが書けた




