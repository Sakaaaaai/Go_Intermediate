# 問20｜テーブルドリブンテスト

## 目標
複数のテストケースをスライスで定義し、ループで一括検証する「テーブルドリブンテスト（Table Driven Test）」というGoの標準的なテスト手法を学びます。

## 今回の演習で学ぶこと
この問題では、Goコミュニティで最も推奨されているテストの書き方を習得します。

---

### なぜテーブルドリブンテストが必要か

Q19の基本テストでは、1つのケースに対して1つのテスト関数を書きました。しかし実際には、同じ関数に対して複数のパターンをテストしたいです：

```go
// 非効率な書き方（繰り返しが多い）
func TestIsEven(t *testing.T) {
    if got := IsEven(2); got != true {
        t.Errorf(...)
    }
    if got := IsEven(3); got != false {
        t.Errorf(...)
    }
    if got := IsEven(0); got != true {
        t.Errorf(...)
    }
    // パターンを追加するたびに同じコードが増える...
}
```

テーブルドリブンテストを使うと、**テストデータ（テーブル）とテストロジックを分離**できます。

---

**1. テーブルドリブンテストとは**

1つの関数に対して「いろいろなパターンの入力と、それに対応する期待値」を一覧（テーブル）として定義し、ループでまとめてテストする手法です。

```go
func TestIsEven(t *testing.T) {
    // テーブル（テストケースの一覧）を定義
    tests := []struct {
        input    int
        expected bool
    }{
        {2, true},   // 正の偶数
        {3, false},  // 奇数
        {0, true},   // 0は偶数
        {-4, true},  // 負の偶数
        {-3, false}, // 負の奇数
    }

    // テーブルをループして全ケースを検証
    for _, tt := range tests {
        got := IsEven(tt.input)
        if got != tt.expected {
            t.Errorf("IsEven(%d) = %v; want %v", tt.input, got, tt.expected)
        }
    }
}
```

新しいケースを追加したいときは、テーブルに1行追加するだけです。ロジック部分は変更不要。

---

**2. name フィールドでテストケースを識別**

テストが失敗したとき「どのケースが失敗したか」をわかりやすくするため、`name` フィールドを追加することが多いです：

```go
tests := []struct {
    name     string
    input    int
    expected bool
}{
    {"正の偶数", 2, true},
    {"奇数", 3, false},
    {"ゼロ", 0, true},
}

for _, tt := range tests {
    got := IsEven(tt.input)
    if got != tt.expected {
        t.Errorf("[%s] IsEven(%d) = %v; want %v", tt.name, tt.input, got, tt.expected)
        // → "[奇数] IsEven(3) = false; want true" のように失敗理由がはっきりわかる
    }
}
```

---

**3. t.Run でサブテストに分ける**

各ケースを `t.Run` でサブテストにすると、どのケースが失敗したかが特定しやすくなります。また `-run` フラグで特定のサブテストだけ実行できます。

```go
for _, tt := range tests {
    tt := tt // ループ変数のコピーを作る（並列実行時に重要）
    t.Run(tt.name, func(t *testing.T) {
        got := IsEven(tt.input)
        if got != tt.expected {
            t.Errorf("got %v; want %v", got, tt.expected)
        }
    })
}
```

実行結果が明確になります：
```
--- FAIL: TestIsEven/奇数 (0.00s)
    main_test.go:30: got true; want false
```

---

**4. なぜテーブルドリブンテストが良いか**

| 項目 | 通常テスト | テーブルドリブンテスト |
|------|-----------|---------------------|
| ケース追加 | ロジックも書き直す | テーブルに1行追加するだけ |
| ロジックの重複 | ある | **なし** |
| バグ再現追加 | 関数を新しく書く | テーブルに1行追加 |
| 可読性 | ケース数が増えると低下 | **ケース数でも維持** |
| エラーメッセージ | どのケースかわかりにくい | **name で明確に** |

---

## 問題文
`main_test.go` の `TestIsEven` 関数を完成させてください。

1. `IsEven(tt.input)` を呼び出して結果を `got` に代入してください。
2. `got` と `tt.expected` を比較し、異なる場合に `t.Errorf` でエラーを出力してください。

### テストの実行
```bash
cd Go_m/Week3/Q20
go test main.go main_test.go
```

### 実行方法
```bash
cd /project/workspace/Week3/Q20
go run main.go
```

### テストでの確認
```bash
cd /project/workspace/Week3/Q20
go test main.go main_test.go
```
## ヒント
- `for _, tt := range tests` でテーブルをループする（`tt` は各テストケース）
- `t.Errorf("IsEven(%d) = %v; want %v", tt.input, got, tt.expected)` の形式が分かりやすい
- 構造体のスライス `[]struct{ ... }{ ... }` でテーブルを定義する

## 達成基準
- [ ] テーブルドリブンテストのパターンを理解できる
- [ ] 構造体のスライスでテストケースを定義できる
- [ ] 複数のテストケースをループで一括検証できる




