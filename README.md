# Go言語 中級講座（全30問）

インターフェース・ゴルーチン（並行処理）・標準ライブラリ・Web サーバー構築を、
手を動かしながら学ぶ自習用教材です。各問題は自動テストで答え合わせができます。

## ここから始める

### https://codespaces.new/Sakaaaaai/Go_Intermediate

1. 上のリンクを開き、**Create codespace** を押します（GitHub アカウントが必要です）。
2. 1〜2 分で、Go が入った VS Code がブラウザ上で開きます。インストール作業は不要です。
3. 下の「演習一覧」から問題を選び、`README.md` を読んで `main.go` を完成させます。
4. ターミナルで実行して、テストで答え合わせをします。

```bash
cd /project/workspace/Week1/Q1
go run main.go
go test main.go main_test.go
```

詳しい進め方は [guide.md](guide.md) を読んでください。

> **なぜ StackBlitz ではないのか**: StackBlitz はブラウザ内で Node.js を動かす仕組みで、Go のコンパイラを持たないため
> `go run` / `go test` を実行できません。この講座は Go をそのまま動かせる GitHub Codespaces を使います。
> 無料枠（月 120 コア時間 / 15GB）で足ります。

## ローカルで動かす場合

[Go](https://go.dev/dl/) をインストールしたうえで:

```bash
git clone https://github.com/Sakaaaaai/Go_Intermediate.git
cd Go_Intermediate/Week1/Q1
go run main.go
go test main.go main_test.go
```

## 構成

| ファイル | 内容 |
| --- | --- |
| `README.md`（各問） | 問題文・ヒント・解説 |
| `main.go` | ここを完成させます（`___` の部分） |
| `answer.go` | 模範解答（詰まったときの参考） |
| `main_test.go` | 自動テスト |

## 演習一覧

| 問 | タイトル | フォルダ |
| --- | --- | --- |
| 1 | インターフェースの基礎 | [開く](https://github.com/Sakaaaaai/Go_Intermediate/tree/main/Week1/Q1) |
| 2 | 空インターフェースと型アサーション | [開く](https://github.com/Sakaaaaai/Go_Intermediate/tree/main/Week1/Q2) |
| 3 | 型スイッチ（Type Switch） | [開く](https://github.com/Sakaaaaai/Go_Intermediate/tree/main/Week1/Q3) |
| 4 | 構造体の埋め込み（Embedding） | [開く](https://github.com/Sakaaaaai/Go_Intermediate/tree/main/Week1/Q4) |
| 5 | 標準インターフェース（fmt.Stringer） | [開く](https://github.com/Sakaaaaai/Go_Intermediate/tree/main/Week1/Q5) |
| 6 | エラーのラップとアンラップ | [開く](https://github.com/Sakaaaaai/Go_Intermediate/tree/main/Week1/Q6) |
| 7 | 高度なエラー判定（errors.Is / errors.As） | [開く](https://github.com/Sakaaaaai/Go_Intermediate/tree/main/Week1/Q7) |
| 8 | ゴルーチン（Goroutine）基礎 | [開く](https://github.com/Sakaaaaai/Go_Intermediate/tree/main/Week2/Q8) |
| 9 | WaitGroupで複数ゴルーチンを待機する | [開く](https://github.com/Sakaaaaai/Go_Intermediate/tree/main/Week2/Q9) |
| 10 | チャネル（Channel）基礎 | [開く](https://github.com/Sakaaaaai/Go_Intermediate/tree/main/Week2/Q10) |
| 11 | バッファ付きチャネル | [開く](https://github.com/Sakaaaaai/Go_Intermediate/tree/main/Week2/Q11) |
| 12 | select文とタイムアウト | [開く](https://github.com/Sakaaaaai/Go_Intermediate/tree/main/Week2/Q12) |
| 13 | 排他制御（Mutex） | [開く](https://github.com/Sakaaaaai/Go_Intermediate/tree/main/Week2/Q13) |
| 14 | Context入門 | [開く](https://github.com/Sakaaaaai/Go_Intermediate/tree/main/Week2/Q14) |
| 15 | ファイルI/O（読み書き） | [開く](https://github.com/Sakaaaaai/Go_Intermediate/tree/main/Week3/Q15) |
| 16 | JSONのパースと生成 | [開く](https://github.com/Sakaaaaai/Go_Intermediate/tree/main/Week3/Q16) |
| 17 | 日付と時刻の操作 | [開く](https://github.com/Sakaaaaai/Go_Intermediate/tree/main/Week3/Q17) |
| 18 | HTTPクライアント | [開く](https://github.com/Sakaaaaai/Go_Intermediate/tree/main/Week3/Q18) |
| 19 | 単体テストを自分で書く | [開く](https://github.com/Sakaaaaai/Go_Intermediate/tree/main/Week3/Q19) |
| 20 | テーブルドリブンテスト | [開く](https://github.com/Sakaaaaai/Go_Intermediate/tree/main/Week3/Q20) |
| 21 | ベンチマークテスト | [開く](https://github.com/Sakaaaaai/Go_Intermediate/tree/main/Week3/Q21) |
| 22 | HTTPサーバーの基礎 | [開く](https://github.com/Sakaaaaai/Go_Intermediate/tree/main/Week4/Q22) |
| 23 | リクエストとレスポンス | [開く](https://github.com/Sakaaaaai/Go_Intermediate/tree/main/Week4/Q23) |
| 24 | マルチプレクサ（ルーティング） | [開く](https://github.com/Sakaaaaai/Go_Intermediate/tree/main/Week4/Q24) |
| 25 | JSON APIの作成 | [開く](https://github.com/Sakaaaaai/Go_Intermediate/tree/main/Week4/Q25) |
| 26 | ミドルウェアの実装 | [開く](https://github.com/Sakaaaaai/Go_Intermediate/tree/main/Week4/Q26) |
| 27 | データベース（モック）の結合 | [開く](https://github.com/Sakaaaaai/Go_Intermediate/tree/main/Week4/Q27) |
| 28 | 総合演習（1）設計と事前準備 | [開く](https://github.com/Sakaaaaai/Go_Intermediate/tree/main/Week4/Q28) |
| 29 | 総合演習（2）一覧の取得と作成（GET / POST） | [開く](https://github.com/Sakaaaaai/Go_Intermediate/tree/main/Week4/Q29) |
| 30 | 総合演習（3）詳細取得と削除（GET / DELETE） | [開く](https://github.com/Sakaaaaai/Go_Intermediate/tree/main/Week4/Q30) |
