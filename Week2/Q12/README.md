# 問12｜select文とタイムアウト

## 目標
複数のチャネルを同時に待ち受ける `select` 文と、`time.After` を組み合わせてタイムアウト処理を実装する方法を学びます。

## 今回の演習で学ぶこと
この問題では、「複数の待ち受けの中で最初に来たものを処理する」というGoらしい並行処理パターンを学びます。

---

### `select` 文が必要な場面

例えば、Webアプリで「外部API呼び出し」と「タイムアウト」を同時に待ちたい場合：

```
- APIが2秒以内に返したら → 成功として処理
- 2秒経ってもAPIが返さなかったら → タイムアウトエラー
```

これを `if/else` や複数の変数で実現しようとすると複雑になります。
`select` 文を使うと、**複数のチャネル操作を同時に待ち受け、最初に準備できたものを実行**できます。

---

**1. select 文の基本**

`select` は `switch` に似た構文ですが、**チャネルの送受信**を条件にします。複数の `case` の中から「最初に準備できたもの」を実行します。

```go
ch1 := make(chan string)
ch2 := make(chan string)

// どこかのゴルーチンが ch1 または ch2 に送信すると想定...

select {
case msg := <-ch1:
    fmt.Println("ch1 から受信:", msg)
case msg := <-ch2:
    fmt.Println("ch2 から受信:", msg)
// どちらも準備できていない間はブロック（待機）
}
```

複数同時に準備できた場合は、**どれが選ばれるかはランダム**です（公平性が保証されている）。

---

**2. time.After でタイムアウト**

`time.After(d)` は「`d` 時間後に現在時刻を1回送信するチャネル」を返します。これを `select` の `case` に組み合わせると、簡単にタイムアウト処理が書けます。

```go
ch := make(chan string)

// 誰かが ch に送信するかもしれないゴルーチンを起動
go func() {
    time.Sleep(3 * time.Second) // 重い処理（3秒かかる）
    ch <- "完了"
}()

select {
case msg := <-ch:
    fmt.Println("処理成功:", msg)
case <-time.After(2 * time.Second): // 2秒待っても来なければ
    fmt.Println("タイムアウトしました") // ← これが実行される（3秒 > 2秒）
}
```

---

**3. select の動作タイムライン**

```
0s: select 文に入る（ch と time.After(2s) を両方同時に監視）
2s: time.After の2秒が経過 → タイムアウトの case が選ばれる
3s: ch に送信されるが、もう select を抜けているので無視
```

---

**4. select の default で非ブロッキング受信**

`default` を書くと、どの `case` も準備できていなかった場合に即座に `default` が実行されます（ブロックしない）。

```go
select {
case msg := <-ch:
    fmt.Println(msg)
default:
    fmt.Println("まだ届いていません") // ← ch に何もなければ即実行
}
```

これを使うと「チャネルにデータがあれば処理、なければスキップ」というポーリング（定期確認）パターンが実現できます。

---

**5. select でできること一覧**

| パターン | 説明 |
|----------|------|
| 複数チャネルの同時待機 | どれかが準備できたら実行 |
| タイムアウト処理 | `time.After` との組み合わせ |
| 非ブロッキング受信 | `default` を使う |
| for + select でイベントループ | 繰り返しチャネルを監視 |

---

## 問題文
`main.go` の `___` の部分を埋めてプログラムを完成させてください。

1. ゴルーチンが2秒後に `"Process finished"` を送信するコードがあります。
2. `select` 文を使い、チャネル `ch` からの受信を待機してください。
3. `time.After` で1秒のタイムアウトを設定してください（処理の2秒より早い）。
4. タイムアウトが先に来ることを確認してください。

### 実行方法
```bash
cd /project/workspace/Week2/Q12
go run main.go
```

### テストでの確認
```bash
cd /project/workspace/Week2/Q12
go test main.go main_test.go
```
## ヒント
- `select { case msg := <-ch: ... case <-time.After(N * time.Second): ... }` の形
- タイムアウト値を処理時間より短くするとタイムアウトが発動する
- `time.After` の戻り値はチャネル（`<-chan time.Time`）なので `case <-time.After(...)` と書く

## 達成基準
- [ ] `select` 文で複数チャネルを待ち受けられる
- [ ] `time.After` でタイムアウト処理が実装できる
- [ ] どの `case` が実行されるかを制御できる




