# 問15｜ファイルI/O（読み書き）

## 目標
Goの標準パッケージ `os` を使って、テキストファイルへの書き込みと読み込みを行う方法を学びます。

## 今回の演習で学ぶこと
この問題では、「処理結果をファイルに保存したり、設定ファイルから読み込む」という実務で多用されるファイル操作の基礎を学びます。

---

### なぜファイルI/Oが重要か

プログラムは通常、実行が終わるとメモリ上のデータが消えます。データを**永続化**（プログラム終了後もデータを保持すること）するために、ファイルへの書き込みが必要です。

よくある使い場面：
- **設定ファイルの読み込み**: `config.json` や `settings.txt` を読み込んで設定値を取得
- **ログファイルへの書き込み**: アプリケーションの動作記録をファイルに保存
- **CSVやデータファイルの処理**: 分析用データの読み込みと加工結果の書き込み
- **シリアライズ**: 構造体の状態をファイルに保存して後から復元

---

**1. ファイルへの書き込み（シンプルな方法）**

```go
// os.WriteFile: ファイル全体を一度に書き込む（Go 1.16以降）
content := "書き込む内容\n"
err := os.WriteFile("output.txt", []byte(content), 0644)
if err != nil {
    log.Fatal(err) // エラーが起きたらログを出して終了
}
```

- 第1引数: ファイルパス（存在しなければ新規作成）
- 第2引数: 書き込む内容（`[]byte` バイト列）
- 第3引数 `0644`: Unixのファイルパーミッション
  - `0`: 8進数表記を示すプレフィックス
  - `6`: オーナー（自分）は読み書き可 (`4+2=6`)
  - `4`: グループは読み取り可
  - `4`: その他のユーザーは読み取り可

---

**2. ファイルからの読み込み**

```go
// os.ReadFile: ファイル全体を一度に読み込む（Go 1.16以降）
data, err := os.ReadFile("output.txt") // data は []byte
if err != nil {
    log.Fatal(err)
}
content := string(data) // []byte を string に変換して表示
fmt.Println(content)
```

`[]byte` と `string` の変換は頻出パターンです：
- ファイル → Goのプログラム: `data, _ := os.ReadFile(...)` → `string(data)`
- Goのプログラム → ファイル: `"文字列"` → `[]byte("文字列")` → `os.WriteFile(...)`

---

**3. より細かい読み書き（行単位など）**

長大なファイルを逐次的に読むには `os.Open` と `bufio.Scanner` の組み合わせが一般的です。`os.ReadFile` は全データを一度にメモリに読み込むため、巨大ファイルには不向きです。

```go
file, err := os.Open("data.txt")
if err != nil { log.Fatal(err) }
defer file.Close() // 必ず閉じる（忘れるとリソースリーク）

scanner := bufio.NewScanner(file)
for scanner.Scan() {
    line := scanner.Text() // 1行取得
    fmt.Println(line)      // 1行ずつ処理
}
```

---

**4. defer file.Close() が重要**

ファイルを開いたら、関数が終わるときに必ず閉じないとリソース（**ファイルディスクリプタ**）がリークします。

ファイルディスクリプタとはOS が管理するファイルへの接続情報で、数に上限があります。閉じ忘れると上限に達して新しいファイルを開けなくなります。

```go
file, err := os.Open("data.txt")
if err != nil { log.Fatal(err) }
defer file.Close()  // 関数の最後に自動的に呼ばれる

// ... ファイルの処理 ...
// 処理がどこで終わっても（return や panic でも）Close が呼ばれる
```

---

**5. os.WriteFile vs os.OpenFile**

| 方法 | 使いどころ | 特徴 |
|------|-----------|------|
| `os.WriteFile` | 全データを一度に書く | シンプル、ファイルを丸ごと上書き |
| `os.OpenFile` | 追記・細粒度の制御 | `os.O_APPEND` で追記可能 |
| `bufio.Writer` | 効率的な書き込み | バッファリングで高速 |

---

## 問題文
`main.go` の `___` の部分を埋めてプログラムを完成させてください。

1. `WriteFile(filename, content string)` 関数を作成し、`os.WriteFile` でファイルに文字列を書き込んでください。
2. `ReadFile(filename string) string` 関数を作成し、`os.ReadFile` でファイルを読み込んで文字列を返してください。
3. `main` 関数で `"test.txt"` に `"Hello Go!"` を書き込み、読み込んで出力してください。

### 実行方法
```bash
cd /project/workspace/Week3/Q15
go run main.go
```

### テストでの確認
```bash
cd /project/workspace/Week3/Q15
go test main.go main_test.go
```
## ヒント
- `os.WriteFile` と `os.ReadFile` はGo 1.16以降で使える便利な関数
- `[]byte("文字列")` で文字列をバイト列に変換する
- `string(data)` でバイト列を文字列に変換する
- ファイルパーミッション `0644` はほとんどの場合でOK

## 達成基準
- [ ] ファイルへの書き込みができる
- [ ] ファイルからの読み込みができる
- [ ] `defer file.Close()` の重要性を説明できる




