# 問7｜高度なエラー判定（errors.Is / errors.As）

## 目標
ラップされた複数のエラー層を貫通して、特定のエラーかどうかを判定する `errors.Is` と、特定の型のエラーを取り出す `errors.As` を習得します。

## 今回の演習で学ぶこと
この問題では、問6で学んだエラーラップを前提として、呼び出し側でのスマートなエラー判定を学びます。

---

### なぜ `errors.Is` が必要か？

Q6でエラーをラップして返す方法を学びました。では、呼び出し側でそのエラーを判定する場合はどうすればよいでしょうか？

```go
var ErrNotFound = errors.New("not found")

func FindUser(id string) error {
    // ... 処理 ...
    return fmt.Errorf("FindUser: %w", ErrNotFound) // ErrNotFound をラップして返す
}

// 呼び出し側でエラーを判定したい
err := FindUser("123")
```

---

**1. = での比較ではダメな理由**

問6でエラーをラップした場合、`err == ErrNotFound` では判定できません。ラップされた外側のエラーと内側のエラーは別物だからです。

```go
wrappedErr := fmt.Errorf("operation failed: %w", ErrNotFound)

fmt.Println(wrappedErr == ErrNotFound)  // false！（ラップされているので別のエラー）
fmt.Println(wrappedErr)                 // "operation failed: not found"（メッセージは見える）
```

`==` での比較は「同じオブジェクトかどうか」を調べますが、ラップすると別のオブジェクトになるので、同じ意味のエラーでも `false` になってしまいます。

---

**2. errors.Is：「このエラーが含まれているか」を判定**

`errors.Is` は何重にラップされていても、中に指定したエラーが含まれているかを調べてくれます。内側のエラーを次々と `Unwrap` しながら探していきます。

```go
var ErrNotFound = errors.New("not found")

// 2重にラップされたエラー
wrappedErr := fmt.Errorf("outer: %w", fmt.Errorf("inner: %w", ErrNotFound))

fmt.Println(errors.Is(wrappedErr, ErrNotFound)) // true！（何重でも貫通して探す）
fmt.Println(wrappedErr == ErrNotFound)           // false（== は貫通しない）
```

実用的な使い方：
```go
err := FindUser("123")
if errors.Is(err, ErrNotFound) {
    // ユーザーが見つからなかった場合の処理
    fmt.Println("ユーザーが存在しません")
} else if err != nil {
    // その他のエラー
    fmt.Println("予期しないエラー:", err)
}
```

---

**3. errors.As：「特定の型のエラーとして取り出す」**

`errors.As` は、エラーの連鎖の中から特定の型のエラーを取り出します。独自エラー型を定義している場合に活用します。

```go
// 独自のエラー型（フィールドを持てる）
type ValidationError struct {
    Field string
    Msg   string
}
func (e *ValidationError) Error() string {
    return fmt.Sprintf("%s: %s", e.Field, e.Msg)
}

// 使い方
var valErr *ValidationError
if errors.As(err, &valErr) {
    // valErr にエラーの詳細情報が格納されている
    fmt.Println("バリデーションエラー:", valErr.Field, valErr.Msg)
}
```

`errors.Is` と異なり、エラーの **値** ではなく **型** を調べて取り出します。

---

**4. is/As の使い分け**

| 用途 | 使うもの | 例 |
|------|----------|----|
| 特定のエラー変数が原因かどうか | `errors.Is` | `errors.Is(err, ErrNotFound)` |
| 特定のエラー型として取り出したい | `errors.As` | `errors.As(err, &valErr)` |

---

**5. `errors.New` でグローバルなセンチネルエラーを定義する**

アプリケーション全体で使う「既知のエラー」は、グローバル変数として定義するのが慣例です：

```go
// パッケージレベル（関数の外）で定義
var (
    ErrNotFound       = errors.New("not found")
    ErrPermissionDenied = errors.New("permission denied")
    ErrTimeout        = errors.New("timeout")
)
```

これらを「**センチネルエラー**（見張り番エラー）」と呼びます。`errors.Is` で比較対象として使われます。

---

## 問題文
`main.go` の `___` の部分を埋めてプログラムを完成させてください。

1. グローバルなエラー変数 `ErrPermissionDenied` を `errors.New` で定義してください。
2. 関数 `DoSomething()` が `ErrPermissionDenied` を `%w` でラップして返すようにしてください。
3. `main` 関数で `errors.Is` を使い、返ってきたエラーが `ErrPermissionDenied` を含むか判定してください。

### 実行方法
```bash
cd /project/workspace/Week1/Q7
go run main.go
```

### テストでの確認
```bash
cd /project/workspace/Week1/Q7
go test main.go main_test.go
```
## ヒント
- `errors.Is(err, ErrTarget)` はネストしたエラーも潜って確認する
- `errors.As(err, &target)` で特定型のエラーを抽出できる（今回は`Is`を使う）
- センチネルエラーは `var ErrXxx = errors.New("説明文")` で定義する

## 達成基準
- [ ] グローバルなエラー変数（センチネルエラー）を定義できる
- [ ] `errors.Is` と `errors.As` の違いを説明できる
- [ ] ラップされたエラーを `errors.Is` で正確に判定できる




