# 問29｜総合演習（2）一覧の取得と作成（GET / POST）

## 目標
Q28で作った `BookRepository` を使って、HTTPリクエストのメソッドに応じて「一覧取得（GET）」と「新規作成（POST）」を同一ハンドラーで処理するRESTful APIを完成させます。

## 今回の演習で学ぶこと
この問題では、同じURLでも「どのHTTPメソッドが来たか」で処理を切り替えるREST APIの基本設計を学びます。

---

### REST APIの設計哲学　〜「URL は名詞、メソッドは動詞」

REST APIの設計では「**リソース（物）**」と「**操作（動作）**」を分けて考えます：

```
URL（名詞）       HTTPメソッド（動詞）  意味
/books           GET                  本の一覧を「読む」
/books           POST                 新しい本を「作る」
/books/123       GET                  ID=123の本を「読む」
/books/123       DELETE               ID=123の本を「消す」
```

URLは「何に対して操作するか（リソース）」、HTTPメソッドは「何をするか（操作）」を表します。
同じ `/books` というURLでもメソッドが違えば全く別の操作です。

---

**1. RESTの設計原則（同じURL、異なるメソッド）**

```
GET  /books  → 一覧を返す（読み取り操作 = データ変更なし）
POST /books  → 新規作成（書き込み操作 = データ変更あり）
```

同じ `/books` というURLに対して、HTTPメソッドで意味を変えるのがRESTの考え方です。

---

**2. HTTPメソッドによる分岐**

```go
func (h *Handler) BooksHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:  // "GET"
        // 一覧を取得してJSONで返す
        books := h.repo.GetAll()
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(books)
        
    case http.MethodPost: // "POST"
        // リクエストボディのJSONをデコードして保存
        var book Book
        if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
            http.Error(w, "Invalid JSON", http.StatusBadRequest) // 400
            return
        }
        h.repo.Save(book)
        w.WriteHeader(http.StatusCreated) // 201 Created
        
    default:
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed) // 405
    }
}
```

---

**3. json.NewDecoder でリクエストのJSONを受け取る**

クライアントが POST で送ってきたデータ（リクエストボディ）を構造体に変換します：

```go
// クライアントが送るJSON（例）:
// {"id": "1", "title": "Go入門", "author": "山田太郎"}

var book Book
err := json.NewDecoder(r.Body).Decode(&book)
// book.ID = "1", book.Title = "Go入門", book.Author = "山田太郎"
```

`json.Unmarshal` との違い：
- `json.Unmarshal`: `[]byte` からデコード（一度全データを読む）
- `json.NewDecoder(r.Body).Decode`: `io.Reader` から直接デコード（ストリーミング・効率的）

---

**4. レスポンスの空スライス対策**

本がまだ1冊も登録されていない場合：

```go
// NG: nil スライスを返すと JSON で null になる
var books []Book // nil スライス
json.Encode(books) // → null

// OK: 空スライスを返すと JSON で [] になる（APIとして親切）
books := []Book{}  // 空スライス
json.Encode(books)  // → []
```

クライアントは `null` より `[]` の方が「配列が空」とわかりやすく扱いやすいです。

---

**5. GETとPOSTの処理フロー**

```
【GET /books リクエスト】
ブラウザ/クライアント
  └─ GET /books ──→ BooksHandler
                         └─ h.repo.GetAll()
                         └─ JSON エンコード ──→ [{...}, {...}]

【POST /books リクエスト】
ブラウザ/クライアント
  └─ POST /books ──→ BooksHandler
  Body: {"id":"1",...}       └─ JSON デコード
                              └─ h.repo.Save(book)
                              └─ 201 Created ──→ 空のレスポンス
```

---

## 問題文
`main.go` の `___` の部分を埋めてプログラムを完成させてください。

1. `r.Method` が `GET` の場合：`h.repo.GetAll()` で全件取得しJSONで返してください。
2. `r.Method` が `POST` の場合：リクエストボディをデコードして `h.repo.Save(book)` し、201を返してください。
3. それ以外の場合：405 Method Not Allowed を返してください。

### 実行方法
```bash
cd /project/workspace/Week4/Q29
go run main.go
```

### テストでの確認
```bash
cd /project/workspace/Week4/Q29
go test main.go main_test.go
```
## ヒント
- `switch r.Method { case http.MethodGet: ... case http.MethodPost: ... default: ... }` の形
- `json.NewDecoder(r.Body).Decode(&book)` でリクエストのJSONを読む
- `w.WriteHeader(http.StatusCreated)` で201を返す

## 達成基準
- [ ] `r.Method` でHTTPメソッドを確認できる
- [ ] GETとPOSTで異なるロジックを実行できる
- [ ] `json.NewDecoder(r.Body).Decode` でリクエストJSONを構造体に変換できる




