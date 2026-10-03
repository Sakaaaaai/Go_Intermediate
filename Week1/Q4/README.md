# 問4｜構造体の埋め込み（Embedding）

## 目標

Go言語における「継承」に代わるコード再利用の仕組みである「構造体の埋め込み（Embedding）」を学びます。

## 今回の演習で学ぶこと

この問題では、Go言語が持つ「継承なし、合成あり」の設計哲学を実践的に学びます。

---

### Goには「継承」がない　〜 なぜコンポジションで代用するのか

Java や Python を学んだことがある方は「継承」をご存知かもしれません：

```java
// Java での継承の例
class Admin extends User {  // Admin が User を継承（extends）
    ...
}
```

**Goにはこの「継承（extends）」がありません。** 代わりに使うのが「**埋め込み（Embedding）**」です。

**「継承」の問題点**:

- 親クラスの変更が子クラス全体に影響する（密結合）
- 「is-a」関係（AdminはUserである）を強制する

**「コンポジション（合成）」の利点**:

- 「has-a」関係（AdminはUserを持っている）で柔軟
- 必要な機能だけを組み合わせられる

---

**1. Goには継承がない**

JavaやPythonのような `class A extends B` という継承の仕組みはGoにはありません。代わりに、**構造体を別の構造体の中に埋め込む**ことで、フィールドやメソッドを使い回します。これを「コンポジション（合成）」と呼びます。

---

**2. 埋め込みの書き方**

フィールド名を書かずに型名だけを書くと「埋め込み」になります。

```go
type User struct {
    Name string
    Age  int
}

// 通常のフィールドとして持つ場合（埋め込みではない）
type Admin1 struct {
    User User   // フィールド名 "User" が必要
    Role string
}

// 埋め込みの場合（型名のみ）
type Admin struct {
    User        // フィールド名なし → 埋め込み！
    Role string
}
```

---

**3. 埋め込んだ型のフィールド・メソッドが使える**

`Admin` のインスタンスから、埋め込んだ `User` のフィールドやメソッドに**直接**アクセスできます。まるで `Admin` 自身のフィールドのように使えます。

```go
admin := Admin{
    User: User{Name: "田中", Age: 30},
    Role: "管理者",
}

// どちらの書き方でも同じ
fmt.Println(admin.Name)       // admin.User.Name の省略形
fmt.Println(admin.User.Name)  // こちらでも同じ

// User に定義したメソッドも呼べる
fmt.Println(admin.Profile())  // admin.User.Profile() と同じ
```

> **重要**: `admin.Name` と書けるのは、Goが自動的に `admin.User.Name` として補完してくれるからです。これを「**フィールドのプロモーション**」と呼びます。

---

**4. メソッドのオーバーライド（上書き）**

埋め込んだ型と同じ名前のメソッドを定義すると、そちらが優先されます（上書き）。

```go
// User の Profile メソッド
func (u User) Profile() string {
    return fmt.Sprintf("%sさんは%d歳です", u.Name, u.Age)
}

// Admin 独自の Profile（User のものを上書き）
func (a Admin) Profile() string {
    // User の Profile を呼びつつ、Role 情報を追加
    return fmt.Sprintf("%s（役割: %s）", a.User.Profile(), a.Role)
}
```

`Admin` で呼ぶと Admin 版の `Profile`、`User` を取り出して呼ぶと User 版の `Profile` が実行されます。

---

**5. 埋め込みの図解**

```
Admin 構造体
├── User（埋め込み）
│   ├── Name: string  ← admin.Name でアクセス可能
│   ├── Age: int      ← admin.Age でアクセス可能
│   └── Profile()     ← admin.Profile() でアクセス可能
└── Role: string      ← admin.Role でアクセス可能
```

---

## 問題文

`main.go` の `___` の部分を埋めてプログラムを完成させてください。

1. `string` 型の `Name` と `int` 型の `Age` を持つ `User` 構造体を定義してください。
2. `User` に `Profile() string` メソッドを実装し、`"<Name>さんは<Age>歳です"` を返してください。
3. `User` を埋め込んだ `Admin` 構造体を定義し、`string` 型の `Role` フィールドを追加してください。
4. `Admin` インスタンスから `Profile()` メソッドを呼べることを確認してください。

### 実行方法

```bash
cd /project/workspace/Week1/Q4
go run main.go
```

### テストでの確認

```bash
cd /project/workspace/Week1/Q4
go test main.go main_test.go
```

## ヒント

- 型名だけを書くと埋め込みになる：`type Admin struct { User; Role string }`
- `admin.Name` は `admin.User.Name` の短縮形（フィールドのプロモーション）
- どちらの書き方でもアクセスできる
- 構造体の初期化：`Admin{User: User{Name: "田中", Age: 30}, Role: "管理者"}`

## 達成基準

- [ ] 構造体への埋め込みが書ける
- [ ] 埋め込んだ構造体のメソッドを呼び出せる
- [ ] フィールドへの直接アクセス（プロモーション）ができる
