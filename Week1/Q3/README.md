# 問3｜型スイッチ（Type Switch）

## 目標
`any` 型（空インターフェース）で渡ってきた値を、実際の型ごとに処理を分岐させる「型スイッチ」の書き方をマスターします。

## 今回の演習で学ぶこと
この問題では、型スイッチを使って**型によって動作が変わる汎用的な関数**を書く方法を学びます。

---

### 型スイッチが必要な場面

Q2では `%v` を使えばどんな型でも文字列化できることを学びました。しかし、**型によって表示形式を変えたい**場合はどうすればよいでしょうか？

```
- int が来たら → "整数です: 42"
- string が来たら → "文字列です: hello"
- bool が来たら → "真偽値です: true"
```

これを型アサーションで書こうとすると：

```go
if val, ok := v.(int); ok {
    return fmt.Sprintf("整数です: %d", val)
}
if val, ok := v.(string); ok {
    return fmt.Sprintf("文字列です: %s", val)
}
// ... こんなにifが増えていく
```

読みにくいですね。**型スイッチを使うと、一気にすっきり書けます。**

---

**1. 型スイッチの基本書式**

型スイッチは通常の `switch` 文と似ていますが、`v.(type)` という特殊な書き方で型によって分岐します。

```go
func Describe(v any) string {
    switch val := v.(type) {
    case int:
        return fmt.Sprintf("整数です: %d", val)
    case string:
        return fmt.Sprintf("文字列です: %s", val)
    case bool:
        return fmt.Sprintf("真偽値です: %v", val)
    default:
        return "未知の型です"
    }
}
```

> **ポイント**: `switch val := v.(type)` の `val` には、各ケースで**その型として変換された値**が入ります。
> - `case int:` の中では `val` は `int` 型
> - `case string:` の中では `val` は `string` 型
> - `default:` の中では `val` の型は元の `any` のまま

---

**2. 通常の型アサーションとの違い**

| | 型アサーション | 型スイッチ |
|--|--------------|------------|
| 書き方 | `v.(型名)` | `switch v.(type)` |
| 対応型数 | 1つ | 複数 |
| 失敗時 | パニックの可能性あり | 安全（`default` で対応） |
| 使いどころ | 型が決まっているとき | 複数の型を安全に扱うとき |

```go
// 型アサーション（1つの型に決め打ち）
s := v.(string) // v が string でなければパニック

// 型スイッチ（複数の型に安全に対応）
switch v.(type) {
case string:
    ...
case int:
    ...
}
```

---

**3. 型スイッチの実用的な使いどころ**

実際の開発でよく使われる場面を覚えておきましょう：

- **JSONをパース**して得られた `interface{}` 型のデータの種類を調べるとき
  - JSONの配列・オブジェクト・数値・文字列を型スイッチで識別
- **ライブラリのエラー型**により処理を分けるとき
  - ネットワークエラー・タイムアウトエラー・認証エラーを型で区別
- **汎用的なユーティリティ関数**を作るとき
  - ログ出力関数やシリアライゼーション処理

---

**4. 複数の型を同じケースにまとめる書き方**

同じ処理をしたい型が複数ある場合、`,` で区切ってまとめられます：

```go
switch val := v.(type) {
case int, int64, float64:  // 数値系は全部まとめて
    return fmt.Sprintf("数値: %v", val)
case string:
    return fmt.Sprintf("文字列: %s", val)
}
```

---

## 問題文
`main.go` の `___` の部分を埋めてプログラムを完成させてください。

`Describe(v any) string` 関数を実装してください。引数の型に応じて以下の文字列を返します：
- `int` 型なら `"整数です: <数値>"`
- `string` 型なら `"文字列です: <文字列>"`
- `bool` 型なら `"真偽値です: <true/false>"`
- それ以外なら `"未知の型です"`

### 実行方法
```bash
cd /project/workspace/Week1/Q3
go run main.go
```

### テストでの確認
```bash
cd /project/workspace/Week1/Q3
go test main.go main_test.go
```
## ヒント
- `switch val := v.(type)` の `val` には各ケースの型に変換された値が入る
- `fmt.Sprintf("整数です: %d", val)` でケースごとにフォーマットを変える
- `default:` は必ず最後に書く（どのケースにも一致しなかった場合の処理）

## 達成基準
- [ ] 型スイッチが書ける
- [ ] `default` ケースで未知の型に対応できる
- [ ] `val` を使って型ごとに適切な表示ができる




