# 問5｜標準インターフェース（fmt.Stringer）

## 目標
Go標準ライブラリが定義するインターフェース `fmt.Stringer` を実装して、独自の構造体をきれいに文字列として表示する方法を学びます。

## 今回の演習で学ぶこと
この問題では、Goが「標準で用意しているインターフェース」を活用することで、自作の型が標準ライブラリの出力関数と自然に連携できることを学びます。

---

### なぜ `Stringer` が必要か？

自分で作った構造体を `fmt.Println` で表示するとどうなるでしょうか？

```go
type Book struct {
    Title  string
    Author string
}

b := Book{Title: "Go入門", Author: "山田太郎"}
fmt.Println(b)  // → {Go入門 山田太郎}
```

`{Go入門 山田太郎}` と出力され、少し分かりにくいですよね。本としてちゃんと表示したいなら `『Go入門』(著: 山田太郎)` のような形式にしたいはずです。

`fmt.Stringer` インターフェースを実装すると、**`fmt.Println` が自動的にその形式で表示**してくれるようになります。

---

**1. fmt.Stringer とは**

`fmt.Stringer` は `fmt` パッケージに定義されているインターフェースで、以下の1つのメソッドを持ちます。

```go
// fmt パッケージ内での定義（参考）
type Stringer interface {
    String() string
}
```

Goの `fmt.Println` や `fmt.Sprintf("%v", ...)` などは、渡した値が `Stringer` を満たしていれば、自動的に `String()` メソッドを呼び出して表示してくれます。

---

**2. 実装の例**

```go
type Point struct {
    X, Y int
}

// ← String() メソッドを実装するだけで Stringer を満たす
func (p Point) String() string {
    return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

p := Point{X: 3, Y: 5}
fmt.Println(p)               // → (3, 5)  と表示される（明示的に String() を呼ばなくていい）
fmt.Println(p.String())      // → (3, 5)  （明示的に呼んでも同じ結果）
s := fmt.Sprintf("%v", p)    // → "(3, 5)"（%v でも String() が使われる）
```

---

**3. なぜ便利か**

`String()` を定義せずに `fmt.Println(p)` すると `{3 5}` と内部表現がそのまま出てしまいます。
`Stringer` を実装することで、以下のすべての場所で一貫した表示ができます：

| 表示方法 | String()未実装 | String()実装済み |
|----------|--------------|----------------|
| `fmt.Println(p)` | `{3 5}` | `(3, 5)` |
| `fmt.Sprintf("%v", p)` | `{3 5}` | `(3, 5)` |
| `fmt.Sprintf("%s", p)` | エラー | `(3, 5)` |
| ログ出力 | `{3 5}` | `(3, 5)` |

---

**4. 他の重要な標準インターフェースの例**

Goには `Stringer` 以外にも多くの標準インターフェースがあります：

| インターフェース | メソッド | 用途 |
|----------------|---------|------|
| `fmt.Stringer` | `String() string` | カスタム文字列表示 |
| `error` | `Error() string` | エラー型として扱う |
| `io.Reader` | `Read(p []byte) (n int, err error)` | データ読み取り |
| `io.Writer` | `Write(p []byte) (n int, err error)` | データ書き込み |

これらも「このメソッドを実装するだけ」で標準ライブラリと連携できます。

---

## 問題文
`main.go` の `___` の部分を埋めてプログラムを完成させてください。

1. `Title` と `Author` の2つの `string` フィールドを持つ `Book` 構造体を定義してください。
2. `Book` 構造体に `String() string` メソッドを実装し、`"『<Title>』(著: <Author>)"` という形式の文字列を返してください。
3. `fmt.Println(Book{...})` を実行したとき、上記の形式が表示されることを確認します。

### 実行方法
```bash
cd /project/workspace/Week1/Q5
go run main.go
```

### テストでの確認
```bash
cd /project/workspace/Week1/Q5
go test main.go main_test.go
```
## ヒント
- `fmt.Stringer` を「実装した」と言うには `String() string` メソッドを書くだけでいい（`implements` キーワードは不要）
- `fmt.Sprintf("『%s』(著: %s)", b.Title, b.Author)` のように書ける
- `String()` を実装すると `fmt.Println` が自動でそれを呼び出す（明示的に `.String()` を呼ばなくてよい）

## 達成基準
- [ ] `fmt.Stringer` インターフェースが何かを説明できる
- [ ] `String() string` メソッドを実装できる
- [ ] `fmt.Println()` で自動的にカスタム表示されることを確認できる




