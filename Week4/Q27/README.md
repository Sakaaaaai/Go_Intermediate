# 問27｜データベース（モック）の結合

## 目標
「データを保存・取得する」操作を**インターフェースで抽象化**し、テスト用のインメモリ（メモリ上の）データストレージを実装します。これはDI（依存性の注入）の基礎にもなります。

## 今回の演習で学ぶこと
この問題では、「後から本物のDBに差し替えられる設計」を実践します。

---

### なぜインターフェースでデータベースを抽象化するのか

本番のサーバーでは PostgreSQL や MySQL などのデータベースを使いますが、テスト時に実際の DB を使うと：

- **遅い**: DBの起動・接続に時間がかかる
- **不安定**: ネットワーク・DB サーバーの状態に依存
- **後始末が必要**: テストデータの削除が必要
- **環境依存**: 開発環境ごとに DB の設定が違う

**インメモリ実装（メモリ上の疑似DB）をテスト時に使う**と、上記の問題をすべて解決できます。

これを可能にする設計が「**リポジトリパターン**」です。

---

**1. なぜインターフェースで抽象化するか**

DBへの直接依存があると：
- 本物のDB接続なしにテストできない
- 後でDBを変えると（MySQL→PostgreSQL）コードを大幅に書き換えになる

インターフェースで契約だけを定義すれば：
- テスト時はインメモリ実装を使える（高速・安定）
- 本番では本物のDB実装を使える
- 切り替えは初期化部分だけを変えればいい

---

**2. リポジトリパターン**

データアクセス処理をインターフェースにまとめる設計パターンです。

```go
// インターフェース（契約）
// "こういう操作ができる" という定義だけ。どうやって保存するかは書かない
type BookRepository interface {
    Save(id string, book Book)       // 保存（作成または更新）
    Get(id string) (Book, bool)      // 取得
}

// インメモリ実装（テスト用）
type MemoryRepo struct {
    data map[string]Book  // データはメモリ上のマップ
}

func (r *MemoryRepo) Save(id string, book Book) {
    r.data[id] = book
}

func (r *MemoryRepo) Get(id string) (Book, bool) {
    book, ok := r.data[id]
    return book, ok
}
```

---

**3. 使い方（DI：依存性の注入）**

```go
// ハンドラーがインターフェース型の repo を持つ
// 具体的な実装（MemoryRepo or PostgresRepo）を知らなくていい
type Handler struct {
    repo BookRepository // インターフェース型
}

// テスト時は MemoryRepo を注入（高速・独立）
h := &Handler{repo: &MemoryRepo{data: make(map[string]Book)}}

// 本番では DB接続実装を注入
// h := &Handler{repo: &PostgresRepo{db: db}}
```

これを「**DI（Dependency Injection = 依存性の注入）**」と言います。ハンドラーが「どのリポジトリを使うか」を外から注入することで、テストと本番で同じハンドラーコードが使えます。

---

**4. インターフェースで設計できている場合の図**

```
Handler
  │
  └──── BookRepository（インターフェース）
           │
           ├── MemoryRepo   ← テスト時はこちら
           └── PostgresRepo ← 本番はこちら
```

---

## 問題文
`main.go` の `___` の部分を埋めてプログラムを完成させてください。

1. `Title` と `Author` の2フィールドを持つ `Book` 構造体を定義してください。
2. `Save` と `Get` メソッドを持つ `BookRepository` インターフェースを定義してください。
3. `map[string]Book` を持つ `MemoryRepo` を作成し、`BookRepository` を実装してください。

### 実行方法
```bash
cd /project/workspace/Week4/Q27
go run main.go
```

### テストでの確認
```bash
cd /project/workspace/Week4/Q27
go test main.go main_test.go
```
## ヒント
- インターフェースで定義 → 構造体で実装（Week1の内容の復習）
- `MemoryRepo` は `&MemoryRepo{data: make(map[string]Book)}` で初期化
- `Get` メソッドは `book, ok := r.data[id]` で見つかったかも返す

## 達成基準
- [ ] リポジトリパターンの概念を説明できる
- [ ] インターフェースを使ってデータアクセスを抽象化できる
- [ ] 同一インターフェースの複数実装を切り替えられる設計が理解できる




