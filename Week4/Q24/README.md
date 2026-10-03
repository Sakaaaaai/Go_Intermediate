# 問24｜マルチプレクサ（ルーティング）

## 目標
`http.ServeMux` を使い、URLごとに異なるハンドラーを割り当てるルーティングの仕組みを学びます。

## 今回の演習で学ぶこと
この問題では、アプリの「入口」を整理するルーティングの仕組みを学びます。

---

### なぜルーティングが必要か

Q22では1つのハンドラーで全リクエストを処理しました。しかし実際のWebアプリには複数のページや機能があります：

```
/ → トップページを表示
/users → ユーザー一覧を返す
/products → 商品一覧を返す
/api/health → ヘルスチェックを返す
```

これをルーティングなしで実現しようとすると：

```go
func handler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path == "/" {
        // トップページ
    } else if r.URL.Path == "/users" {
        // ユーザー一覧
    } else if r.URL.Path == "/products" {
        // 商品一覧
    }
    // ... どんどん長くなる
}
```

`ServeMux` を使うとこれを分かりやすく整理できます。

---

**1. なぜルーティングが必要か**

`http.HandleFunc("/", ...)` だけでは、`/users` でも `/products` でも全部同じハンドラーで受け取ってしまいます。URLによって異なる処理を行うには「ルーティング」が必要です。

---

**2. http.ServeMux（マルチプレクサ）**

`ServeMux` は「あのURLにきたらこのハンドラーを使う」という対応表を管理するルーターです。

```go
mux := http.NewServeMux() // 独自のルーターを作成

mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "pong")
})

mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "ユーザー一覧")
})

mux.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "商品一覧")
})

// 第2引数に mux を渡してサーバー起動
http.ListenAndServe(":8080", mux)
```

---

**3. デフォルトServeMuxとの違い**

Q22で使った `http.HandleFunc` はグローバルな `DefaultServeMux` に登録するため、テストや大規模アプリでは問題が起こりやすいです。

| | デフォルトServeMux | 独自 ServeMux |
|--|-----------------|--------------|
| 作成 | 不要（グローバル） | `http.NewServeMux()` |
| 登録 | `http.HandleFunc(...)` | `mux.HandleFunc(...)` |
| テスト | 状態が残って汚染されやすい | **独立して使える** |
| 大規模アプリ | パターンの管理が難しい | **モジュール分割しやすい** |

`http.NewServeMux()` で独自インスタンスを作ることで：
- テスト時に独立したルーターを使える
- グローバル状態汚染を避けられる

---

**4. パターンマッチングのルール**

どのURLがどのハンドラーにマッチするかのルールです：

```
/ping     → /ping にのみ完全一致
/api/     → /api/ で始まるすべてのパスにマッチ（末尾スラッシュ = 前方一致）
/         → どのパスにもマッチ（デフォルトハンドラー）
```

複数のパターンが一致する場合は**より具体的なパターン**が優先されます：

```
URL: /api/users
  /api/ にもマッチ
  / にもマッチ
  → /api/ が優先（より具体的）
```

---

**5. 404（Not Found）の返し方**

`/` にしかハンドラーがなければ全部のURLが `/` にマッチしてしまいます。
特定のパスにのみ応答したい場合は、パスチェックを追加します：

```go
mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/ping" { // 完全一致チェック
        http.NotFound(w, r) // 404 を返す
        return
    }
    fmt.Fprint(w, "pong")
})
```

---

## 問題文
`main.go` の `___` の部分を埋めてプログラムを完成させてください。

1. `http.NewServeMux()` で新しいマルチプレクサ `mux` を作成してください。
2. `/ping` に `"pong"` を返すハンドラーを登録してください。
3. `/info` に `"Go Server Setup"` を返すハンドラーを登録してください。
4. `http.ListenAndServe` の第2引数に `mux` を渡してください。

### 実行方法
```bash
cd /project/workspace/Week4/Q24
go run main.go
```

### テストでの確認
```bash
cd /project/workspace/Week4/Q24
go test main.go main_test.go
```
## ヒント
- `mux.HandleFunc("/path", func(w, r) { ... })` でルートを追加する
- 第2引数を `nil` にするとデフォルトのグローバルルーターが使われる
- `mux.HandleFunc` と `http.HandleFunc` の違いを意識しよう

## 達成基準
- [ ] `http.NewServeMux` でカスタムルーターを作れる
- [ ] 複数のURLに対して異なるハンドラーを割り当てられる
- [ ] `ListenAndServe` の第2引数の意味を説明できる




