# 問14｜Context入門

## 目標
並行処理の「キャンセル」や「タイムアウト」を安全に伝播させるための標準的な仕組み `context.Context` の基本的な使い方を学びます。

## 今回の演習で学ぶこと
この問題では、実務で必須となるキャンセル可能な非同期処理のパターンを学びます。

---

### `context.Context` が必要な理由

Q12では `select` + `time.After` でタイムアウトを実装しました。実際のWebアプリケーションでは、1つのHTTPリクエストに対して複数のゴルーチンが動きます：

```
HTTPリクエスト受信
├── ゴルーチンA: データベース検索
├── ゴルーチンB: 外部API呼び出し
└── ゴルーチンC: キャッシュ確認
```

ユーザーがリクエストをキャンセルしたとき（ブラウザを閉じるなど）、**これらを一斉に止める**必要があります。
`context.Context` がその「キャンセル伝播」の役割を担います。

---

**1. Context が必要な理由**

Webサーバーでは1つのリクエストに対して複数の処理（DB検索、外部API呼び出しなど）が走ります。ユーザーがリクエストをキャンセルしたとき、**それらを一斉に止める仕組み**が必要です。`context.Context` がその役割を担います。

---

**2. Context の種類**

```go
// ルートとなるコンテキスト（キャンセルなし・終了なし）
ctx := context.Background()  // プログラムのルートで使う
ctx := context.TODO()        // まだ Context が決まっていない場所の仮置き

// タイムアウト付きContext（指定時間後に自動キャンセル）
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel() // 必ずリソースを解放する

// 手動キャンセル可能Context
ctx, cancel := context.WithCancel(context.Background())
cancel() // 任意のタイミングでキャンセルできる

// 期限付きContext（特定の日時でキャンセル）
deadline := time.Now().Add(5 * time.Second)
ctx, cancel := context.WithDeadline(context.Background(), deadline)
```

---

**3. キャンセルの検知**

`ctx.Done()` はキャンセルされると値が送られてくるチャネルです。`select` と組み合わせます。

```go
func doWork(ctx context.Context, done chan struct{}) {
    // 重い処理をシミュレート
    time.Sleep(5 * time.Second)
    done <- struct{}{}
}

ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()

done := make(chan struct{})
go doWork(ctx, done)

select {
case <-done:
    fmt.Println("処理完了")
case <-ctx.Done():
    fmt.Println("キャンセルまたはタイムアウト:", ctx.Err())
    // ctx.Err() は:
    // - context.DeadlineExceeded（タイムアウト）
    // - context.Canceled（手動キャンセル）
}
```

---

**4. Context の伝播（連鎖）**

Context は「親から子」に伝播します。親がキャンセルされると、子も自動的にキャンセルされます：

```
Background()
└── WithTimeout(5s) ← ここがキャンセルされると...
    ├── WithCancel() ← 自動的にキャンセル
    └── WithCancel() ← 自動的にキャンセル
```

これによって、ルートのコンテキストをキャンセルするだけで、全ての関連処理を一斉に止められます。

---

**5. Context の受け渡しのルール**

実務でよく使われる慣習です：

```go
// ✅ 良い例：Context を引数の第1引数として渡す
func DoWork(ctx context.Context, arg string) error { ... }

// ❌ 悪い例：Context をフィールドに持たせる
type Worker struct {
    ctx context.Context  // フィールドに持たせない！
}
```

- Context は引数の**最初**に渡す慣例：`func DoWork(ctx context.Context, arg string) error`
- Context をフィールドに持たせない（引数で渡す）
- `context.TODO()` はまだContextを決めていない場合の仮置き
- `defer cancel()` は必ず書く（リソースリーク防止）

---

## 問題文
`main.go` の `___` の部分を埋めてプログラムを完成させてください。

1. `context.Background()` をベースに `context.WithTimeout` で2秒のタイムアウトContextを作ってください。
2. 重い処理ゴルーチン（`slowProcess`）にContextを渡し、`select` で完了チャネル `done` と `ctx.Done()` を待ち受けてください。
3. 処理が5秒かかるのに対し、タイムアウトが2秒なのでキャンセルされる動作を確認してください。

### 実行方法
```bash
cd /project/workspace/Week2/Q14
go run main.go
```

### テストでの確認
```bash
cd /project/workspace/Week2/Q14
go test main.go main_test.go
```
## ヒント
- `ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)`
- `defer cancel()` を必ず書く（リソースリーク防止）
- `ctx.Err()` でキャンセルの理由を確認できる（`context.DeadlineExceeded` or `context.Canceled`）

## 達成基準
- [ ] `context.WithTimeout` でタイムアウト付きContextを作れる
- [ ] `ctx.Done()` でキャンセルを検知できる
- [ ] `defer cancel()` の重要性を説明できる




