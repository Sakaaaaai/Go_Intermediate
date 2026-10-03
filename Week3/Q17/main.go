package main

import (
	"fmt"
	"time"
)

// 1. FormatTime は time.Time を "yyyy/mm/dd hh:mm:ss" の文字列にして返す
func FormatTime(t time.Time) string {
	// ヒント: t.Format() を使い、Goの基準日時（2006/01/02 15:04:05）を指定します
	return ___
}

// 2. ParseTime は "yyyy-mm-dd" の文字列を time.Time に変換して返す
func ParseTime(s string) (time.Time, error) {
	// ヒント: time.Parse() を使い、第一引数に基準日時（2006-01-02）、第二引数に文字列 s を渡します
	return ___
}

func main() {
	now := time.Now()
	fmt.Println("現在時刻:", FormatTime(now))

	dateStr := "2023-12-31"
	parsed, err := ParseTime(dateStr)
	if err != nil {
		fmt.Println("エラー:", err)
		return
	}
	fmt.Println("パースされた日付:", parsed.Year(), "年", parsed.Month(), "月", parsed.Day(), "日")
}
