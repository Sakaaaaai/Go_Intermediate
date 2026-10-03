# 問25｜JSON APIの作成

## 目標
GoのWebサーバーから、HTML文字列ではなく「JSON形式のデータ」をレスポンスとして返し、フロントエンドアプリなどから利用できるWeb APIを作る方法を学びます。

## 今回の演習で学ぶこと
この問題では、Week3で学んだJSON処理とHTTPサーバーを組み合わせて、実際のAPIに近い実装を行います。

---

### JSON APIとHTMLの違い

Webの世界には2種類のサーバーがあります：

**HTMLサーバー（従来型）**:
```
クライアント → GET /page
サーバー → <html><body>こんにちは</body></html>
ブラウザがHTMLをそのまま表示
```

**JSON API（現代型）**:
```
クライアント → GET /api/users
サーバー → {"users": [{"name": "Alice"}, {"name": "Bob"}]}
クライアントが受け取ってJavaScriptが動的に表示 or モバイルアプリが使う
```

JSON APIのメリット：
- **ブラウザ以外でも使える**: モバイルアプリ、他のサービスなど
- **フロントエンドと独立**: React/Vue などのSPAとの連携が容易
- **再利用可能**: 同じAPIを複数の画面で使える

---

**1. JSON APIの基本**

REST APIでは、HTMLを返すのではなくデータをJSON形式で返すのが標準です。
クライアント（ブラウザやアプリ）はそのJSONを受け取って自由に使います。

---

**2. Content-Type ヘッダーの設定**

JSONを返す際は、必ず `Content-Type: application/json` を設定します。これでクライアントは「JSONが来る」とわかり、適切に処理できます。

```go
w.Header().Set("Content-Type", "application/json")
```

> **重要**: `w.Header().Set(...)` は、`w.Write(...)` や `w.WriteHeader(...)` を呼ぶ**前に**設定しなければなりません。一度書き込み開始するとヘッダーは変更できません。

設定しなかった場合：ブラウザがJSONを「普通のテキスト」として扱い、アプリが正しく処理できないことがあります。

---

**3. json.NewEncoder で直接レスポンスに書き込む**

```go
type Response struct {
    Status  string `json:"status"`
    Message string `json:"message"`
}

// 方法1: json.Marshal + w.Write（2ステップ）
jsonBytes, _ := json.Marshal(Response{Status: "ok", Message: "Hello"})
w.Write(jsonBytes)

// 方法2: json.NewEncoder（1ステップ・効率的）
w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(Response{
    Status:  "success",
    Message: "Hello from Go API!",
})
```

`json.NewEncoder(w).Encode(...)` は `json.Marshal` + `w.Write` + 改行の追加を1回でやってくれます。ストリーミング処理に適しており、大きなデータでもメモリ効率が良いです。

---

**4. 典型的なJSON APIのレスポンス形式**

```json
// 成功レスポンス
{"status": "success", "data": {"name": "Alice", "age": 30}}

// エラーレスポンス
{"status": "error", "message": "User not found"}

// 一覧レスポンス
{"status": "success", "data": [...], "total": 100}
```

---

**5. エラーレスポンス**

```go
if err != nil {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusInternalServerError) // 500
    json.NewEncoder(w).Encode(map[string]string{
        "error": "Internal Server Error",
    })
    return
}
```

---

## 問題文
`main.go` の `___` の部分を埋めてプログラムを完成させてください。

1. `w.Header().Set(...)` で `Content-Type` を `"application/json"` に設定してください。
2. `json.NewEncoder(w).Encode(...)` で `Response` 構造体をJSONとして書き込んでください。

### 実行方法
```bash
cd /project/workspace/Week4/Q25
go run main.go
```

### テストでの確認
```bash
cd /project/workspace/Week4/Q25
go test main.go main_test.go
```
## ヒント
- ヘッダーの設定は `w.Header().Set("Content-Type", "application/json")`
- `json.NewEncoder(w).Encode(構造体)` は変換と書き込みを同時に行う
- ヘッダー設定は書き込みより**前**に行う（順番が重要）

## 達成基準
- [ ] `Content-Type: application/json` を設定してJSON APIレスポンスを返せる
- [ ] `json.NewEncoder(w).Encode` で構造体をJSONに変換して書き込める
- [ ] JSON APIとHTML返却の違いを説明できる




