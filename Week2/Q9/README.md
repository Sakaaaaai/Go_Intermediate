# 問9｜WaitGroupで複数ゴルーチンを待機する

## 目標
`sync.WaitGroup` を使い、複数のゴルーチンすべての処理が終わるのを安全に待機する方法を学びます。

## 今回の演習で学ぶこと
この問題では、「全員が仕事を終えるまで待つ」ための仕組みを学びます。

---

### Q8の `time.Sleep` の問題点をより詳しく理解する

Q8では `time.Sleep` でゴルーチンの完了を待ちましたが、これには大きな問題があります：

```go
go doWork1() // 100ms かかる
go doWork2() // 500ms かかる？それとも50ms？
go doWork3() // 不明

time.Sleep(1 * time.Second) // 「たぶん1秒あれば終わるだろう」という賭け

// 問題1: 短すぎたら処理が途中で切れる
// 問題2: 長すぎたら無駄に待つ
// 問題3: ネットワーク状況によって時間が変わる
```

**「全員が完了したら進む」を正確に表現する仕組み**が `sync.WaitGroup` です。

---

**1. time.Sleep による待機の問題点**

問8では `time.Sleep` でゴルーチンの完了を待ちましたが、実際の処理時間は事前にわかりません。長すぎると無駄な待機、短すぎると処理が途中で打ち切られます。

---

**2. sync.WaitGroup の仕組み**

`WaitGroup` は内部にカウンターを持ちます。工事現場の「入場者カウンター」をイメージするとわかりやすいです：

```
Add(3)  ← 「3人が現場に入ります」
Done()  ← 「1人が退場しました（カウント-1）」
Wait()  ← 「全員が退場するまでゲートで待ちます」
```

具体的なコード：

```go
var wg sync.WaitGroup

wg.Add(3) // 3つのゴルーチンを起動するよ（カウント=3）

for i := 1; i <= 3; i++ {
    go func(id int) {
        defer wg.Done() // 終わったら必ず呼ばれる（カウント-1）
        fmt.Printf("Worker %d 完了\n", id)
    }(i)
}

wg.Wait() // カウントが0になるまでここで待機
fmt.Println("全員完了！") // ← ここに来るのは全ゴルーチン完了後
```

---

**3. defer wg.Done() が重要**

`defer` を使うと、関数がパニックで終了した場合でも必ず `Done` が呼ばれます。`defer` なしだと、処理の途中でエラーが起きると `Wait()` が永遠に終わらなくなります（**デッドロック**）。

```go
// 良い例（defer を使う）
func worker(wg *sync.WaitGroup) {
    defer wg.Done()  // ← 一番最初に書く。何があっても必ず呼ばれる
    // ... 処理 ...
}

// 危険な例（defer を使わない）
func worker(wg *sync.WaitGroup) {
    // ... 処理 ...
    if err != nil {
        return // Done を呼び忘れ！Wait() が永遠に待ち続ける！
    }
    wg.Done()
}
```

---

**4. WaitGroup の使用上のルール**

```
1. Add は ゴルーチン起動の前に呼ぶ
   ※ 起動後に Add すると競合が起きる可能性がある

2. Done はゴルーチンの中で（defer で）呼ぶ

3. Wait は全 Add に対応する Done が呼ばれた後に完了する

4. ゴルーチン間で WaitGroup を共有する場合はポインタで渡す
   ※ var wg sync.WaitGroup は正しい（コピーしない）
```

---

**5. WaitGroup とチャネルの使い分け**

| 目的 | 使うもの |
|------|---------|
| 「完了を待つだけ」 | `WaitGroup` がシンプル |
| 「完了と一緒にデータを受け取りたい」 | チャネル（Q10）が適切 |
| 「タイムアウト付きで待ちたい」 | Context と select（Q12/Q14）|

---

## 問題文
`main.go` の `___` の部分を埋めてプログラムを完成させてください。

1. `sync.WaitGroup` のインスタンス `wg` を作成してください。
2. 3つのゴルーチン起動前に、同時に `wg.Add(3)` を呼んでください。
3. 各ゴルーチン内で `defer wg.Done()` を呼んでください。
4. `main` 最後に `wg.Wait()` を書いて全ゴルーチン完了を待ってください。

### 実行方法
```bash
cd /project/workspace/Week2/Q9
go run main.go
```

### テストでの確認
```bash
cd /project/workspace/Week2/Q9
go test main.go main_test.go
```
## ヒント
- `var wg sync.WaitGroup` でゼロ値のまま使える（`new` は不要）
- `defer wg.Done()` をゴルーチンの一番最初に書くのが慣例
- `wg.Add(n)` は起動する**前に**呼ぶこと

## 達成基準
- [ ] `sync.WaitGroup` で全ゴルーチンの完了を待機できる
- [ ] `Add` / `Done` / `Wait` の役割を説明できる
- [ ] `defer` で安全に `Done` を呼べる




