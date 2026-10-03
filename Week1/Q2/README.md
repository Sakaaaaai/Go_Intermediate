# 問2｜空インターフェースと型アサーション

## 目標
あらゆる型の値を保持できる `any`（空インターフェース）と、格納された値の実際の型を取り出す「型アサーション」を学びます。

## 今回の演習で学ぶこと
この問題では、Goの柔軟な型システムを活用した**動的な値の扱い方**を学びます。

---

### なぜ「空インターフェース」が必要か？

Q1では「`Speak()` を実装している型」を `Speaker` インターフェースで表しました。では、「**どんな型でも受け取れる**」関数を作りたい場合はどうすればよいでしょうか？

例えば、「何でも受け取って、それを文字列にして返す」ような汎用の関数です。これに答えるのが `any`（空インターフェース）です。

---

**1. 空インターフェース（any）とは**

`any` は「すべての型を受け入れられる型」です。Go 1.18 以降では `any` と書けますが、本質は `interface{}` と同じです。
「メソッドを一切持たないインターフェース」なので、あらゆる型が自動的に満たします。

```go
var v any = 42        // int を格納
v = "Hello"           // string に変更もできる
v = []int{1, 2, 3}   // スライスも格納できる
v = true              // bool も格納できる
```

> **注意**: `any` に一度値を入れると、取り出して使う前に「中身は何の型か」を確認する必要があります。

---

**2. 型アサーション：中に入っている型を確認して取り出す**

`any` 型の変数から実際の値を取り出すには「**型アサーション**」を使います。

```go
var v any = "Go言語"

// v の中身が string だとわかっているなら
s := v.(string)  // string 型として取り出す
fmt.Println(s)   // "Go言語"
```

ただし、もし中身が `string` でなければ**パニック（強制終了）**になります。
安全な書き方は `ok` を使います：

```go
s, ok := v.(string)  // ok が true なら成功、false なら失敗（パニックしない）
if ok {
    fmt.Println("文字列:", s)
} else {
    fmt.Println("文字列ではありませんでした")
}
```

> **使い分け**:
> - 型が**確実**にわかっているとき → `s := v.(string)`（シンプルだが失敗するとパニック）
> - 型が**不確実**なとき → `s, ok := v.(string)`（安全）

---

**3. 型スイッチ：型によって処理を分岐する**

複数の型に対応したい場合は `switch v.(type)` が便利です。型アサーションを複数の `if` で書くより格段に読みやすくなります。

```go
func describe(v any) string {
    switch val := v.(type) {
    case int:
        return fmt.Sprintf("整数: %d", val)
    case string:
        return fmt.Sprintf("文字列: %s", val)
    case bool:
        return fmt.Sprintf("真偽値: %v", val)
    default:
        return fmt.Sprintf("未知の型: %T", val) // %T で型名を表示
    }
}
```

---

**4. `fmt.Sprintf` の `%v` で汎用フォーマット**

実は型を特定しなくても、`%v` というフォーマット動詞を使えば、**あらゆる型の値を文字列として表示**できます。

```go
fmt.Sprintf("Value: %v", 42)        // "Value: 42"
fmt.Sprintf("Value: %v", "hello")   // "Value: hello"
fmt.Sprintf("Value: %v", true)      // "Value: true"
fmt.Sprintf("Value: %v", []int{1,2}) // "Value: [1 2]"
```

これが今回の演習で使う核心的なポイントです！

---

## 問題文
`main.go` の `___` の部分を埋めてプログラムを完成させてください。

`FormatValue(v any) string` 関数を実装してください。どんな型の引数でも `"Value: <値>"` という形式の文字列を返します。

- `int` 型なら `"Value: <数値>"`
- `string` 型なら `"Value: <文字列>"`
- その他の型でも `"Value: <値>"` を返す（`fmt.Sprintf` で統一）

### 実行方法
```bash
cd /project/workspace/Week1/Q2
go run main.go
```

### テストでの確認
```bash
cd /project/workspace/Week1/Q2
go test main.go main_test.go
```
## ヒント
- `fmt.Sprintf("Value: %v", v)` の `%v` はあらゆる型に対応した汎用フォーマット
- 型アサーションを使わなくても `%v` で解決できる場合がある（型スイッチを使わずに実装可能！）
- `%T` を使うと値の型名を確認できる（デバッグに便利）

## 達成基準
- [ ] `any` 型の変数にさまざまな型の値を代入できる
- [ ] `fmt.Sprintf` と `%v` を使って値を文字列化できる
- [ ] 型アサーションの安全な書き方（`v, ok := x.(型)`）を理解している




