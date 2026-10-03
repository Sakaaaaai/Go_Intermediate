# 問30｜総合演習（3）詳細取得と削除（GET / DELETE）

## 目標

URLのパスからIDを取り出し、特定の書籍の「取得（GET）」と「削除（DELETE）」を行う処理を実装して、中級コースの集大成とします。

## 今回の演習で学ぶこと

この問題では、URLに含まれるIDを解析する「URLパラメータの抽出」と、完全なRESTfulエンドポイントの構築を学びます。

---

### 「1件操作」の考え方

Q29では `/books`（コレクション全体）を扱いました。Q30では `/books/123` のように**特定の1件**を IDで指定して操作します。

```
/books         → 本のコレクション全体
/books/abc-123 → ID が "abc-123" の本1冊
```

同じ `/books/<ID>` というURLでも：

- `GET /books/abc-123` → その本の詳細を取得
- `DELETE /books/abc-123` → その本を削除

---

**1. URLからIDを取り出す**

標準の `net/http` にはパスパラメータ（`:id`のような記法）の仕組みがないため、文字列操作でIDを取り出します。

```go
// /books/abc-123 → "abc-123" を取り出す
id := strings.TrimPrefix(r.URL.Path, "/books/")

// strings.TrimPrefix の動作:
// TrimPrefix("/books/abc-123", "/books/") → "abc-123"
// TrimPrefix("/books/", "/books/")        → "" （空文字 = IDが指定されていない）
// TrimPrefix("/users/123", "/books/")     → "/users/123" （プレフィックスが違う場合そのまま）
```

---

**2. 1件取得の実装（GET）**

```go
case http.MethodGet:
    book, err := h.repo.GetByID(id)
    if err != nil {
        http.Error(w, "Not Found", http.StatusNotFound) // 見つからなければ 404
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(book) // 見つかればJSONで返す
```

---

**3. 削除の実装（DELETE）**

```go
case http.MethodDelete:
    err := h.repo.Delete(id)
    if err != nil {
        http.Error(w, "Not Found", http.StatusNotFound) // 既に存在しない
        return
    }
    // 削除成功時は「204 No Content」を返す
    // 成功したけどレスポンスボディがないことを示す
    w.WriteHeader(http.StatusNoContent)
```

---

**4. HTTPステータスコードまとめ（すべての演習の総復習）**

Q28〜Q30を通じて使うステータスコードをまとめます：

| コード | 定数                             | 名前                  | 使いどころ                               |
| ------ | -------------------------------- | --------------------- | ---------------------------------------- |
| 200    | `http.StatusOK`                  | OK                    | 成功・データあり（デフォルト）           |
| 201    | `http.StatusCreated`             | Created               | リソースを新規作成した                   |
| 204    | `http.StatusNoContent`           | No Content            | 成功だがレスポンスボディなし（DELETE等） |
| 400    | `http.StatusBadRequest`          | Bad Request           | リクエストが不正（JSONが壊れているなど） |
| 404    | `http.StatusNotFound`            | Not Found             | 指定IDのリソースが存在しない             |
| 405    | `http.StatusMethodNotAllowed`    | Method Not Allowed    | 対応していないHTTPメソッド               |
| 500    | `http.StatusInternalServerError` | Internal Server Error | サーバー側の予期しないエラー             |

---

**5. 図書管理API 全体の設計まとめ**

```
GET    /books         → 全件取得（一覧）         [Q29]
POST   /books         → 新規作成                [Q29]
GET    /books/<id>    → 特定の1件を取得          [Q30]
DELETE /books/<id>    → 特定の1件を削除          [Q30]
```

---

**6. APIの動作確認（curl コマンドの例）**

サーバーを起動してから、コマンドラインで動作確認できます（参考）：

```bash
# 一覧取得
curl http://localhost:8080/books

# 新規作成
curl -X POST http://localhost:8080/books \
  -H "Content-Type: application/json" \
  -d '{"id":"1","title":"Go入門","author":"山田太郎"}'

# 1件取得
curl http://localhost:8080/books/1

# 削除
curl -X DELETE http://localhost:8080/books/1
```

---

## 問題文

`main.go` の `___` の部分を埋めてプログラムを完成させてください。

1. `strings.TrimPrefix(r.URL.Path, "/books/")` でURLからIDを取り出してください。
2. `GET` の場合：`h.repo.GetByID(id)` で取得し、見つかればJSON、なければ404を返してください。
3. `DELETE` の場合：`h.repo.Delete(id)` を呼び、成功したら204を返してください。

### 実行方法

```bash
cd /project/workspace/Week4/Q30
go run main.go
```

### テストでの確認

```bash
cd /project/workspace/Week4/Q30
go test main.go main_test.go
```

## ヒント

- `strings.TrimPrefix("/books/123", "/books/")` → `"123"`
- 見つからなかったとき：`http.Error(w, "Not Found", http.StatusNotFound)`
- 削除成功時：`w.WriteHeader(http.StatusNoContent)` （204 = ボディなしで成功）

## 達成基準

- [ ] URLパスからIDを文字列操作で取り出せる
- [ ] GET（詳細）とDELETEのAPIが実装できる
- [ ] 適切なHTTPステータスコード（200/204/404/405）を返せる

---

## 🎉 中級コース完了！

おめでとうございます！ここまで完了した方は、Go言語を使った本格的なREST APIサーバーを一から構築できる実力が身についています。

学んだことの振り返り：

- **Week1**: インターフェース・型システム・エラーハンドリング
- **Week2**: ゴルーチン・チャネル・Mutex・Context による並行処理
- **Week3**: ファイルI/O・JSON・日時・HTTPクライアント・テスト手法
- **Week4**: HTTPサーバー・ルーティング・JSON API・ミドルウェア・CRUD API
