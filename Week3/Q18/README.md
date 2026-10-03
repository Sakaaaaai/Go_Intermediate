# 問18｜HTTPクライアント

## 目標
Goの標準パッケージ `net/http` を使い、外部のWebサーバーやAPIにGETリクエストを送り、レスポンスのデータを取得する方法を学びます。

## 今回の演習で学ぶこと
この問題では、別のサービスのAPIからデータを取得するという、バックエンド開発でも頻出の「HTTPクライアント実装」を学びます。

---

### HTTPクライアントとは

ブラウザでウェブページを開くとき、ブラウザはWebサーバーに「HTTPリクエスト」を送り、HTMLを受け取ります。Go のプログラムでも同様に、外部サービスの API に HTTP リクエストを送ってデータを取得できます。

よくある使い場面：
- **外部API連携**: 天気予報API、地図API、支払いAPIなどのデータ取得
- **マイクロサービス間通信**: 別のGoサービスのAPIを呼び出す
- **Webスクレイピング**: ウェブページの内容を取得して解析
- **ヘルスチェック**: 外部サービスの死活監視

---

**1. 最も基本的な GET リクエスト**

```go
import "net/http"

// http.Get は指定URLにGETリクエストを送る
resp, err := http.Get("https://api.example.com/users")
if err != nil {
    log.Fatal(err) // ネットワークエラーなど
}
defer resp.Body.Close() // 必ずボディを閉じる！
```

`http.Get` の戻り値：
- `resp`: `*http.Response`（レスポンス情報全体）
- `err`: ネットワークエラー（サーバーに繋がらなかった場合）

---

**2. resp.Body を必ず閉じる（重要）**

`http.Get` が返す `resp.Body` は `io.ReadCloser` という「読み取り後に閉じる必要があるもの」です。`defer resp.Body.Close()` を忘れると、TCP/IPコネクションがリリースされず、**リソースリーク（接続が使われ続ける）**が発生します。

```go
resp, err := http.Get(url)
if err != nil { log.Fatal(err) }
defer resp.Body.Close()  // ← この行が重要！
```

> **なぜ `defer` を使うのか？** `defer` を使うと、関数終了時に確実に `Close` が呼ばれます。途中で `return` してもエラーが起きても必ず閉じられます。

---

**3. ボディの読み込み**

```go
// io.ReadAll でレスポンスボディ全体を []byte として読み込む
data, err := io.ReadAll(resp.Body)
if err != nil {
    log.Fatal(err)
}
content := string(data) // []byte → string に変換
fmt.Println(content)
```

大きなレスポンスの場合は `bufio.Scanner` などを使って少しずつ読むこともできますが、通常は `io.ReadAll` で十分です。

---

**4. ステータスコードの確認**

サーバーが返すステータスコードを確認して、成功か失敗かを判断します：

```go
if resp.StatusCode != http.StatusOK { // 200 OK でなければ
    fmt.Printf("エラー: %d %s\n", resp.StatusCode, resp.Status)
    return
}
```

よく使うHTTPステータスコード：

| コード | 定数 | 意味 |
|--------|------|------|
| 200 | `http.StatusOK` | 成功 |
| 201 | `http.StatusCreated` | 作成成功 |
| 400 | `http.StatusBadRequest` | リクエストが不正 |
| 401 | `http.StatusUnauthorized` | 認証が必要 |
| 404 | `http.StatusNotFound` | 見つからない |
| 500 | `http.StatusInternalServerError` | サーバーエラー |

---

**5. POSTリクエストの場合（参考）**

```go
// JSON データを POST する
body := strings.NewReader(`{"name":"Alice","age":30}`)
resp, err := http.Post("https://api.example.com/users", "application/json", body)
```

より柔軟な設定（ヘッダー追加など）には `http.NewRequest` + `http.Client.Do` を使います。

---

**6. リクエストの全体的な流れ**

```
1. http.Get(url) でリクエスト送信
2. エラーチェック（ネットワークエラー）
3. defer resp.Body.Close() でリソース確保
4. resp.StatusCode でHTTPエラーチェック
5. io.ReadAll(resp.Body) でボディ読み込み
6. string(data) で文字列に変換
```

---

## 問題文
`main.go` の `___` の部分を埋めてプログラムを完成させてください。

`FetchWebsite(url string) string` 関数を作成してください：
1. `http.Get(url)` でGETリクエストを送信してください。
2. `defer resp.Body.Close()` でボディのクローズを予約してください。
3. `io.ReadAll` でボディの全データを読み込み、文字列にして返してください。

### 実行方法
```bash
cd /project/workspace/Week3/Q18
go run main.go
```

### テストでの確認
```bash
cd /project/workspace/Week3/Q18
go test main.go main_test.go
```
## ヒント
- `resp, err := http.Get(url)` → `data, err := io.ReadAll(resp.Body)` → `string(data)` の流れ
- `defer resp.Body.Close()` は `http.Get` が成功した直後に書く習慣をつけよう
- `io` パッケージのインポートが必要：`import "io"`

## 達成基準
- [ ] `http.Get` でGETリクエストを送れる
- [ ] `io.ReadAll` でレスポンスボディを読み込める
- [ ] `defer Body.Close()` の意義を説明できる




