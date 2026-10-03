# 問26｜ミドルウェアの実装

## 目標
HTTPリクエストを受け取る「前」や「後」に共通処理を挟み込む「ミドルウェア」パターンの実装方法を学びます。

## 今回の演習で学ぶこと
この問題では、「すべてのリクエストに共通の処理を適用する」という実務で非常に重要なパターンを学びます。

---

### ミドルウェアが必要な理由

Webアプリには「どのエンドポイントでも必要な処理」があります：

```go
// こんな処理を全エンドポイントに書くとしたら...
func usersHandler(w http.ResponseWriter, r *http.Request) {
    log.Printf("Request: %s %s", r.Method, r.URL.Path)  // ← 全エンドポイントで同じ
    // 認証チェック ...                                   // ← 全エンドポイントで同じ
    // 本来の処理
    json.NewEncoder(w).Encode(users)
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
    log.Printf("Request: %s %s", r.Method, r.URL.Path)  // ← コピペ！
    // 認証チェック ...                                   // ← コピペ！
    json.NewEncoder(w).Encode(products)
}
```

コピペが増えると、変更するときに全エンドポイントを修正しなければなりません。
**ミドルウェア**を使うと共通処理を1箇所にまとめられます。

---

**1. ミドルウェアが必要な理由**

各エンドポイントで「アクセスログを出す」「認証チェックをする」という処理を毎回書くのは非効率です。ミドルウェアを使うと共通処理を1箇所にまとめられます。

```
リクエスト → ミドルウェアA（ログ） → ミドルウェアB（認証） → ハンドラー → レスポンス
```

---

**2. Goのミドルウェアパターン**

「ハンドラーを受け取って、新しいハンドラーを返す関数」がミドルウェアです。

```go
// ミドルウェアの型：HandlerFunc を受け取り、HandlerFunc を返す
func Logger(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // ① 前処理（リクエスト受信時のログ出力など）
        start := time.Now()
        fmt.Printf("→ %s %s\n", r.Method, r.URL.Path)
        
        // ② 次のハンドラーを呼び出す（ここが重要！）
        next(w, r)
        
        // ③ 後処理（処理時間の記録など）
        fmt.Printf("← %s %s %v\n", r.Method, r.URL.Path, time.Since(start))
    }
}
```

---

**3. ミドルウェアの動作の流れ**

```
Logger(apiHandler) の場合:

リクエスト受信
   ↓
Logger の前処理（ログ出力）
   ↓
next(w, r) → apiHandler が実行される
   ↓
Logger の後処理（処理時間の記録）
   ↓
レスポンス送信
```

---

**4. ミドルウェアの適用**

```go
mux.HandleFunc("/api", Logger(apiHandler))  // Logger でラップして登録
```

---

**5. 複数のミドルウェアを重ねる**

```go
// Logger → Auth → apiHandler の順で実行
mux.HandleFunc("/api", Logger(Auth(apiHandler)))

// 可読性を上げるためのラップ関数パターン
type MiddlewareFunc func(http.HandlerFunc) http.HandlerFunc

func Chain(handler http.HandlerFunc, middlewares ...MiddlewareFunc) http.HandlerFunc {
    for i := len(middlewares) - 1; i >= 0; i-- {
        handler = middlewares[i](handler)
    }
    return handler
}

mux.HandleFunc("/api", Chain(apiHandler, Logger, Auth, RateLimit))
```

---

**6. よくあるミドルウェアの種類**

| ミドルウェア | 用途 |
|------------|------|
| ログ記録 | アクセスログをファイルに記録 |
| 認証・認可チェック | JWTトークンの検証など |
| レスポンス時間計測 | パフォーマンス監視 |
| CORS ヘッダー設定 | クロスオリジンリクエストの許可 |
| パニックのリカバリー | `recover()` でサーバークラッシュを防ぐ |
| レートリミット | 特定 IP からのリクエスト数を制限 |

---

## 問題文
`main.go` の `___` の部分を埋めてプログラムを完成させてください。

1. `Logger(next http.HandlerFunc) http.HandlerFunc` を実装してください。
2. ハンドラーの先頭で `"Requested: <URLのパス>"` をログ出力してください。
3. その直後に `next(w, r)` を呼び出してハンドラーを継続させてください。

### 実行方法
```bash
cd /project/workspace/Week4/Q26
go run main.go
```

### テストでの確認
```bash
cd /project/workspace/Week4/Q26
go test main.go main_test.go
```
## ヒント
- `func Logger(next http.HandlerFunc) http.HandlerFunc { return func(w, r) { ... } }` の構造
- `r.URL.Path` でリクエストのURLパスを取得できる
- `next(w, r)` を忘れると、本来のハンドラーが実行されないので注意

## 達成基準
- [ ] ミドルウェアパターンの概念を説明できる
- [ ] ハンドラーをラップするミドルウェア関数を書ける
- [ ] `next(w, r)` を呼んで後続ハンドラーに繋げられる




