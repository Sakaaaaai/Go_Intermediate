# 問17｜日付と時刻の操作

## 目標
Goの `time` パッケージを使って、現在時刻の取得・フォーマット・パース（文字列から変換）を行う方法を学びます。

## 今回の演習で学ぶこと
この問題では、Goの時刻処理の**独特のフォーマット定義方法**を理解することが最も重要です。

---

### Goの時刻フォーマットが「変わっている」理由

Python では `YYYY-MM-DD`、PHP では `Y-m-d` などの記号でフォーマットを指定します。しかし Goは全く異なるアプローチを採用しています。

**なぜ Goはこのアプローチを採用したのか？**

従来の `YYYY` などの記号は直感的ではなく、「`M` は月か分か？」などの混乱を生みやすいという問題がありました。Goは「**実際の日時例**でフォーマットを表現する」という哲学を採用し、混乱を減らしました。

---

**1. Goの時刻フォーマットの独特な仕様**

Goは「特定の基準日時」を使います。その基準は：

**2006年1月2日（月曜日）15時04分05秒 MST（UTC-7）**

```
年   → 2006
月   → 01
日   → 02
時   → 15（24時間表記）/ 03（12時間表記）
分   → 04
秒   → 05
```

**覚え方**: `1, 2, 3, 4, 5, 6` という1〜6の数字の並びになっています（3だけ15時として例外）。

| 要素 | Go の値 | 他言語での記号 |
|------|--------|----------------|
| 年 | `2006` | `YYYY` |
| 月（2桁） | `01` | `MM` |
| 日（2桁） | `02` | `DD` |
| 時（24h） | `15` | `HH` |
| 分 | `04` | `mm` |
| 秒 | `05` | `ss` |

---

**2. 時刻のフォーマット**

```go
t := time.Now() // 現在時刻を取得

// 様々なフォーマット例
fmt.Println(t.Format("2006/01/02 15:04:05"))  // 例: "2024/03/15 14:30:00"
fmt.Println(t.Format("2006年1月2日"))          // 例: "2024年3月15日"
fmt.Println(t.Format("Jan 2, 2006"))           // 例: "Mar 15, 2024"
fmt.Println(t.Format("15:04"))                 // 例: "14:30"（時分のみ）

// RFC3339（ISO 8601）形式：APIでよく使われる
fmt.Println(t.Format(time.RFC3339))  // "2024-03-15T14:30:00+09:00"
```

---

**3. 文字列から時刻へのパース**

```go
dateStr := "2024-03-15"
t, err := time.Parse("2006-01-02", dateStr) // フォーマット文字列, 入力文字列
if err != nil {
    log.Fatal(err)
}
fmt.Println(t.Year(), t.Month(), t.Day())  // 2024 March 15

// タイムゾーンに注意
// time.Parse は UTC として解釈する
// ローカルタイムとして扱いたい場合は time.ParseInLocation を使う
```

---

**4. 時間の計算**

```go
t := time.Now()
future := t.Add(24 * time.Hour)     // 24時間後
past := t.Add(-7 * 24 * time.Hour)  // 7日前

diff := future.Sub(t)                // 2つの時刻の差（Duration型）
fmt.Println(diff)                    // "24h0m0s"
fmt.Println(diff.Hours())            // 24.0
fmt.Println(diff.Minutes())          // 1440.0

// 比較
fmt.Println(future.After(t))  // true（future の方が後）
fmt.Println(past.Before(t))   // true（past の方が前）
```

---

**5. 日本のタイムゾーン（JST）を使う**

```go
// JST（日本標準時）での時刻取得
jst, _ := time.LoadLocation("Asia/Tokyo")
now := time.Now().In(jst)
fmt.Println(now.Format("2006年1月2日 15:04:05 MST"))
// "2024年3月15日 23:30:00 JST"
```

---

## 問題文
`main.go` の `___` の部分を埋めてプログラムを完成させてください。

1. `FormatTime(t time.Time) string`：`"2006/01/02 15:04:05"` 形式の文字列に変換して返してください。
2. `ParseTime(s string) time.Time`：`"2006-01-02"` 形式の文字列を `time.Time` に変換して返してください。

### 実行方法
```bash
cd /project/workspace/Week3/Q17
go run main.go
```

### テストでの確認
```bash
cd /project/workspace/Week3/Q17
go test main.go main_test.go
```
## ヒント
- `t.Format("2006/01/02 15:04:05")` でフォーマット
- `time.Parse("フォーマット", "入力文字列")` でパース
- フォーマット文字列：`2006`=年, `01`=月, `02`=日, `15`=時, `04`=分, `05`=秒
- 「1月2日、3時4分5秒、2006年」と 1〜6 の順番で覚えると便利

## 達成基準
- [ ] Goの時刻フォーマットのルール（基準日時）を説明できる
- [ ] `Format` で時刻を文字列に変換できる
- [ ] `Parse` で文字列を時刻に変換できる




