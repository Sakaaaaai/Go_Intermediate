# 問28｜総合演習（1）設計と事前準備

## 目標
Week4の集大成となる「図書管理REST API」を3ステップで作ります。このQ28では、API全体で使う「型の定義」と「インメモリDB」を準備します。

## 今回の演習で学ぶこと
この問題では、本格的なREST APIを構築するための土台を、Q27で学んだリポジトリパターンを使って実装します。

---

### これまでの学習の集大成

Week4の総合演習（Q28〜Q30）では、これまで学んだすべての知識を組み合わせます：

| 問題 | 内容 | 使うスキル |
|------|------|-----------|
| Q28 | 型定義・リポジトリ実装 | 構造体、インターフェース、JSON タグ |
| Q29 | 一覧取得・新規作成（GET/POST） | HTTP ハンドラー、JSON エンコード/デコード |
| Q30 | 詳細取得・削除（GET/DELETE） | URL パラメータ、ステータスコード |

---

**1. REST API の設計（CRUD）**

REST APIでは、以下の4つの基本操作をHTTPメソッドとURLで表現します（これを **CRUD** と言います）：

| 操作 | HTTPメソッド | URL | 説明 |
|------|------------|-----|------|
| **C**reate（作成） | `POST` | `/books` | 新しい本を登録 |
| **R**ead（読取）　 | `GET` | `/books` | 全体の一覧を取得 |
| **R**ead（読取）　 | `GET` | `/books/:id` | 特定の本を取得 |
| **U**pdate（更新） | `PUT` | `/books/:id` | 特定の本を更新 |
| **D**elete（削除） | `DELETE` | `/books/:id` | 特定の本を削除 |

---

**2. Book 型の定義**

JSONタグ込みで API のデータ型を定義します。

```go
type Book struct {
    ID     string `json:"id"`     // URLパスで使う識別子
    Title  string `json:"title"`  // 本のタイトル
    Author string `json:"author"` // 著者名
}
```

---

**3. BookRepository インターフェース**

Q27の設計を拡張して、`GetAll`・`GetByID`・`Save`・`Delete` の4操作を定義します。

```go
type BookRepository interface {
    GetAll() []Book                     // 全件取得（一覧）
    GetByID(id string) (Book, error)    // 1件取得（存在しない場合はerror）
    Save(book Book) error               // 保存（新規・更新）
    Delete(id string) error             // 削除
}
```

---

**4. MemoryRepo（インメモリ実装）**

```go
type MemoryRepo struct {
    data map[string]Book  // ID をキーとした本の Map
}

func (r *MemoryRepo) GetAll() []Book {
    books := []Book{}  // nil ではなく空スライスを返す（JSON で [] になる）
    for _, b := range r.data {
        books = append(books, b)
    }
    return books
}

func (r *MemoryRepo) GetByID(id string) (Book, error) {
    book, ok := r.data[id]
    if !ok {
        return Book{}, errors.New("not found") // 見つからなければエラー
    }
    return book, nil
}

func (r *MemoryRepo) Save(book Book) error {
    r.data[book.ID] = book // 同じIDなら上書き
    return nil
}

func (r *MemoryRepo) Delete(id string) error {
    if _, ok := r.data[id]; !ok {
        return errors.New("not found")
    }
    delete(r.data, id) // マップからキーを削除
    return nil
}
```

---

**5. nil と空スライスの違い（重要）**

```go
var a []Book    // nil スライス → JSON では `null`
b := []Book{}  // 空スライス → JSON では `[]`

json.Marshal(a) // → null
json.Marshal(b) // → []
```

APIでは `null` より `[]` を返す方がクライアントにとって扱いやすいため、`books := []Book{}` で初期化します。

---

## 問題文
`main.go` の `___` の部分を埋めてプログラムを完成させてください。

1. `ID`, `Title`, `Author` の3フィールドを持ち、JSONタグ（`"id"`, `"title"`, `"author"`）がついた `Book` 構造体を定義してください。
2. `GetAll`, `GetByID`, `Save`, `Delete` の4つのメソッドを持つ `BookRepository` インターフェースを定義してください。
3. `MemoryRepo`（`map[string]Book`を持つ構造体）に各メソッドを実装してください。

### 実行方法
```bash
cd /project/workspace/Week4/Q28
go run main.go
```

### テストでの確認
```bash
cd /project/workspace/Week4/Q28
go test main.go main_test.go
```
## ヒント
- `GetAll` はすべての `Book` をスライスとして返す（`[]Book{}` で初期化）
- `GetByID` は見つからなかったとき `errors.New("not found")` などを返す
- `Save` はID重複を気にせず `r.data[book.ID] = book` で上書き
- `Delete` はマップからの削除に `delete(map, key)` を使う

## 達成基準
- [ ] REST APIのCRUDとHTTPメソッドの対応を説明できる
- [ ] JSONタグ付きの構造体を定義できる
- [ ] `BookRepository` インターフェースを実装した `MemoryRepo` が書ける




