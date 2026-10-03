package main

import (
	"fmt"
	"time"
)

func FormatTime(t time.Time) string {
	return t.Format("2006/01/02 15:04:05")
}

func ParseTime(s string) (time.Time, error) {
	return time.Parse("2006-01-02", s)
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
