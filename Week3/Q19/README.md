# 問19｜単体テストを自分で書く

## 目標
Goの標準テストフレームワーク `testing` パッケージを使い、自分で書いた関数に対するテストコードを実際に書いてみます。

## 今回の演習で学ぶこと
この問題では、これまで「裏側で実行されていたテストコード」を自分で書く体験をします。

---

### なぜテストを書くのか

「テストを書く」というのは一見、余計な作業に見えますが、実は強力なメリットがあります：

**テストがないと起こること:**
- 修正するたびに「他の部分が壊れていないか」手動で確認が必要
- チームでの開発では、他の人の変更が自分のコードを壊しても気づかない
- リリース直前になってバグが発覚する

**テストがあると:**
- コードを変更したらすぐに `go test` で問題を検出できる
- 「このコードは正しく動いている」という自信を持ってリリースできる
- テストがドキュメントの役割も果たす（使い方を示す）

---

**1. テストファイルのルール**

Goのテストは特別なルールに従います：

```
mypackage/
├── main.go         ← 実装コード
└── main_test.go    ← テストコード（_test.go で終わる）
```

- ファイル名を `_test.go` で終わらせる（例：`main_test.go`）
- テスト関数名を `Test〇〇(t *testing.T)` にする（`Test` で始まる）
- `go test` コマンドを使うとGoが自動で認識して実行する
- テストファイルも同じパッケージに属する（通常は `package main`）

---

**2. テストの基本形**

```go
package main

import "testing"

func TestAdd(t *testing.T) {
    // arrange（準備）: テストに必要なデータを準備
    a, b := 2, 3

    // act（実行）: テスト対象の関数を呼ぶ
    got := Add(a, b)

    // assert（検証）: 期待する結果と比較する
    want := 5
    if got != want {
        t.Errorf("Add(%d, %d) = %d; want %d", a, b, got, want)
        // Errorf: テスト失敗としてマークするが、テスト関数は続ける
    }
}
```

テストの書き方は「**AAA パターン**」と呼ばれる `Arrange → Act → Assert` の3段階で構成すると読みやすいです。

---

**3. t.Errorf と t.Fatalf の違い**

| 関数 | 失敗後の動作 | 使いどころ |
|------|------------|------------|
| `t.Errorf(...)` | 失敗をマークして**残りのテストを続ける** | 複数の条件を全部チェックしたい |
| `t.Fatalf(...)` | 失敗をマークして**その場でテスト関数を終了** | 前提条件が満たされなければ続けても意味がない場合 |
| `t.Logf(...)` | 失敗扱いしない情報ログ | デバッグ情報の記録（`-v` フラグで表示） |

```go
// Errorf の使い方（テストを続ける）
if got != want {
    t.Errorf("Add(2,3) = %d; want %d", got, want)
}
// ← ここに到達する

// Fatalf の使い方（テストを終了する）
if got == nil {
    t.Fatalf("got is nil, cannot continue") // 致命的エラー
}
got.SomeMethod() // ← nil なら呼べないので、Fatalf が適切
```

---

**4. テストの実行コマンド**

```bash
go test main.go main_test.go    # 特定ファイルを指定
go test ./...                   # 全パッケージを再帰的に実行
go test -v ./...                # 詳細ログ付きで実行（各テスト関数名も表示）
go test -run TestAdd ./...      # 特定のテスト関数だけ実行
go test -count=1 ./...          # キャッシュを使わず必ず実行
```

---

**5. テストが「成功」したときの出力例**

```
$ go test ./...
ok  	mypackage  0.012s

$ go test -v ./...
--- PASS: TestAdd (0.00s)
PASS
ok  	mypackage  0.012s
```

テストが「失敗」したときは：
```
--- FAIL: TestAdd (0.00s)
    main_test.go:10: Add(2, 3) = 6; want 5
FAIL
```

---

## 問題文
`main_test.go` を開いて、`TestAdd` 関数を完成させてください。

1. `Add(2, 3)` を呼び出して結果を `got` に代入してください。
2. 期待値 `want = 5` と比較して、異なる場合は `t.Errorf` でエラーを表示してください。

### 実行方法
```bash
cd /project/workspace/Week3/Q19
go run main.go
```

### テストでの確認
```bash
cd /project/workspace/Week3/Q19
go test main.go main_test.go
```
## ヒント
- `*testing.T` は「テストのコントローラー」。失敗通知はすべてここ経由で行う
- `t.Errorf("説明: got=%v, want=%v", got, want)` の形式が一般的
- テストファイルのパッケージを実装と同じ `package main` にするのを忘れずに

## 達成基準
- [ ] テストファイルのルール（`_test.go`, `TestXxx`）を理解している
- [ ] `t.Errorf` でテスト失敗を報告できる
- [ ] `go test` コマンドでテストを実行できる




