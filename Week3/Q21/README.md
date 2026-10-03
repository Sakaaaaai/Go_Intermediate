# 問21｜ベンチマークテスト

## 目標
Goの `testing` パッケージのベンチマーク機能を使い、2つの実装のどちらが処理速度が速いかを計測・比較します。

## 今回の演習で学ぶこと
この問題では、「コードがどれだけ速いか」を**数値として測定する**方法を学びます。これはパフォーマンス最適化の第一歩です。

---

### なぜベンチマークが必要か

「このコードは速い」「あっちの実装の方が効率的だ」という議論は、**数値なしでは主観的な意見**にすぎません。ベンチマークを使うと：

- 2つの実装を**客観的な数値**で比較できる
- 最適化後に本当に改善されたかを確認できる
- 「最適化したつもりが実は遅くなっていた」を防げる

例えば、1から N までの合計を計算する2つの方法：

```go
// ループで計算（O(n)）
func SumLoop(n int) int {
    sum := 0
    for i := 1; i <= n; i++ {
        sum += i
    }
    return sum
}

// 数式で計算（O(1)）
func SumFormula(n int) int {
    return n * (n + 1) / 2
}
```

どちらが速いか「明らか」かもしれませんが、ベンチマークで**実際に何倍速いか**を確認できます。

---

**1. ベンチマークテストの書き方**

関数名を `Benchmark〇〇(b *testing.B)` にし、`b.N` 回ループの中で計測したい処理を呼び出します。`b.N` はGoのテストフレームワークが自動的に最適な繰り返し回数を決めます。

```go
func BenchmarkSumLoop(b *testing.B) {
    for i := 0; i < b.N; i++ {
        SumLoop(100) // 計測したい関数を b.N 回呼ぶ
    }
}

func BenchmarkSumFormula(b *testing.B) {
    for i := 0; i < b.N; i++ {
        SumFormula(100)
    }
}
```

> **`b.N` について**: Goの `testing` パッケージが「信頼できる計測結果を得るために必要な回数」を自動で決めます。最初は少ない回数（例: 100回）で実行し、結果が安定しないなら回数を増やします。あなたが数値を指定する必要はありません。

---

**2. ベンチマークの実行コマンド**

```bash
# -bench=. で全ベンチマーク関数を実行（. は正規表現で「すべて」を意味）
go test -bench=.

# 特定のベンチマーク関数だけ実行
go test -bench=BenchmarkSumLoop

# 実行秒数を指定（デフォルト1秒。精度を上げたい場合は長くする）
go test -bench=. -benchtime=5s

# 通常テストも一緒に実行しない場合（ベンチマークのみ）
go test -bench=. -run=^$  # -run=^$ は "何もマッチしない" の正規表現
```

---

**3. ベンチマーク結果の読み方**

実行すると以下のような出力が得られます：

```
BenchmarkSumLoop-12      1000000          1045 ns/op
BenchmarkSumFormula-12   1000000000          0.5714 ns/op
```

各列の意味：

| 列 | 説明 | 例 |
|----|------|-----|
| `BenchmarkSumLoop-12` | 関数名 + 論理CPU数 | `12` コアを使用 |
| `1000000` | 実行された回数（b.N の最終値） | 100万回実行 |
| `1045 ns/op` | 1回あたりの処理時間 | 1045ナノ秒 |

> `SumFormula` は `SumLoop` よりも **約1828倍高速**（1045 / 0.5714 ≈ 1829）

---

**4. b.ResetTimer の活用**

セットアップ（準備）に時間がかかる場合、計測前にタイマーをリセットします。

```go
func BenchmarkWithSetup(b *testing.B) {
    // セットアップ（大量データ生成など）
    data := generateLargeData()  // これは計測対象外にしたい
    
    b.ResetTimer() // ← ここまでの時間を計測から除外してリセット
    
    for i := 0; i < b.N; i++ {
        process(data)  // これだけを計測したい
    }
}
```

---

**5. ベンチマーク vs ユニットテスト の使い分け**

| | ユニットテスト | ベンチマーク |
|--|------------|------------|
| 目的 | 「正しく動くか」を確認 | 「どれだけ速いか」を計測 |
| 実行方法 | `go test ./...` | `go test -bench=.` |
| 関数命名 | `TestXxx(t *testing.T)` | `BenchmarkXxx(b *testing.B)` |
| 普段の実行 | 毎回 | パフォーマンスが問題になったとき |

---

## 問題文
`main_test.go` の `BenchmarkLoop` と `BenchmarkFormula` を完成させてください。

1. `for i := 0; i < b.N; i++ { ... }` のループを作成してください。
2. `BenchmarkLoop` では `SumLoop(100)` を、`BenchmarkFormula` では `SumFormula(100)` を呼び出してください。

### ベンチマークの実行
```bash
cd Go_m/Week3/Q21
go test -bench=.
```

### 実行方法
```bash
cd /project/workspace/Week3/Q21
go run main.go
```

### テストでの確認
```bash
cd /project/workspace/Week3/Q21
go test main.go main_test.go
```
## ヒント
- 通常テストと違い、ベンチマーク関数は `-bench` フラグなしでは実行されない
- `b.N` はGoが自動設定する → 自分で数値を入れなくてよい
- ベンチマーク結果の `ns/op` が小さいほど速い

## 達成基準
- [ ] ベンチマーク関数の書き方（`Benchmark〇〇(b *testing.B)`）を理解している
- [ ] `b.N` の意味を説明できる
- [ ] `go test -bench=.` でベンチマーク結果を出力できる




