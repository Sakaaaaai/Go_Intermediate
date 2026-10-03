# 問22｜HTTPサーバーの基礎

## 目標
Goの標準パッケージだけでブラウザからアクセスできるWebサーバーを立ち上げ、レスポンスを返す流れを理解します。

## 今回の演習で学ぶこと
この問題では、GoのWebサーバーの最小構成と、リクエストを受け取ってレスポンスを返す仕組みを学びます。

---

### GoでWebサーバーが作れることを知る

Python では Flask や Django、Node.js では Express のような**フレームワーク**をインストールするのが一般的ですが、**Go は標準ライブラリだけでフル機能のWebサーバーを作れます**。

これはGoの設計思想の一つで：
- 依存パッケージが少ない → セキュリティリスクが低い
- 本番環境でのパフォーマンスが高い
- シンプルでわかりやすい

---

**1. Goの HTTP サーバーの最小構成**

これだけでブラウザからアクセスできるサーバーが動きます：

```go
package main

import (
    "fmt"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello, World!")  // ブラウザに文字を送る
}

func main() {
    http.HandleFunc("/", helloHandler)  // "/" にきたら helloHandler を呼ぶ
    http.ListenAndServe(":8080", nil)   // ポート8080でサーバー起動
}
```

起動後に `http://localhost:8080` をブラウザで開くと `Hello, World!` と表示されます。

---

**2. ハンドラー関数の引数**

```go
func myHandler(w http.ResponseWriter, r *http.Request) {
    // w = ResponseWriter
    //   → クライアント（ブラウザ）へのレスポンスを書き込むインターフェース
    //   → fmt.Fprint(w, "文字列") や json.NewEncoder(w).Encode() で書き込む
    //   → io.Writer を実装しているので、書き込み系関数に渡せる
    
    // r = Request
    //   → クライアントからのリクエスト情報をすべて持っている
    //   → r.Method, r.URL, r.Header などでアクセスする
    //   → r.Body でリクエストボディ（POST データなど）を読める
}
```

---

**3. レスポンスの書き込み方**

```go
// テキストを直接書く
fmt.Fprint(w, "Hello!")
fmt.Fprintf(w, "Hello, %s!", name)  // フォーマット付き

// ステータスコードを設定（デフォルトは200 OK）
w.WriteHeader(http.StatusNotFound)   // 404 を返す
w.WriteHeader(http.StatusCreated)    // 201 を返す

// ヘッダーを設定（WriteHeader や Fprint よりも前に行う）
w.Header().Set("Content-Type", "text/plain")
w.Header().Set("Content-Type", "application/json")

// 注意: WriteHeader を呼んだ後はヘッダーの変更ができない
```

---

**4. サーバーが動く仕組み**

```
ブラウザ: GET / HTTP/1.1
   ↓
ListenAndServe が受信
   ↓
URL パターン "/" にマッチするハンドラーを検索
   ↓
helloHandler(w, r) を呼び出す
   ↓
fmt.Fprint(w, "Hello!") でレスポンス書き込み
   ↓
ブラウザ: "Hello!" が表示される
```

---

**5. ポート番号の意味**

```go
http.ListenAndServe(":8080", nil)
//                   ↑
//              ":" は "全てのネットワークインターフェースで待機"
//              8080 は待機するポート番号
```

ポート番号の慣習：
- `:80` → 通常のHTTP（Webサーバーの標準ポート）
- `:443` → HTTPS（暗号化通信の標準ポート）
- `:8080` → 開発用によく使われる（80番は管理者権限が必要）

---

## 問題文
`main.go` の `___` の部分を埋めてプログラムを完成させてください。

1. `http.HandleFunc` でルートパス `"/"` にハンドラーを登録してください。
2. ハンドラー内で `fmt.Fprint(w, "Welcome to Go Web Server!")` を呼んでください。
3. `http.ListenAndServe(":8080", nil)` でポート8080でサーバーを起動してください。

### 実行方法
```bash
cd /project/workspace/Week4/Q22
go run main.go
```

### テストでの確認
```bash
cd /project/workspace/Week4/Q22
go test main.go main_test.go
```
## ヒント
- `fmt.Fprint(w, "...")` は `fmt.Println` ではなく「書き込み先を指定する版」
- `:8080` のように `:` 始まりで書くとすべてのIPアドレスでリッスンします
- `http.ListenAndServe` はブロッキング関数（永遠に待ち続ける）

## 達成基準
- [ ] `http.HandleFunc` でパスにハンドラーを登録できる
- [ ] `http.ResponseWriter` にレスポンスを書き込める
- [ ] `http.ListenAndServe` でサーバーを起動できる




