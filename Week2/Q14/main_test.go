package main

import (
	"bytes"
	"context"
	"io"
	"os"
	"strings"
	"testing"
	"time"
)

func TestContextTimeout(t *testing.T) {
	// タイムアウトするケース
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	if !strings.Contains(output, "処理が中断されました") {
		t.Errorf("Expected output to contain '処理が中断されました', got %q", output)
	}
	
	if !strings.Contains(output, "context deadline exceeded") {
		t.Errorf("Expected context deadline exceeded error message")
	}

	// 成功するケース（1秒で終わる処理にする）
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	done := make(chan struct{})
	
	oldStdout2 := os.Stdout
	r2, w2, _ := os.Pipe()
	os.Stdout = w2

	go func() {
		time.Sleep(1 * time.Second)
		close(done)
	}()
	slowProcess(ctx, done)

	w2.Close()
	os.Stdout = oldStdout2

	var buf2 bytes.Buffer
	io.Copy(&buf2, r2)
	output2 := buf2.String()
	
	if !strings.Contains(output2, "処理が正常に完了しました") {
		t.Errorf("Expected output to contain '処理が正常に完了しました', got %q", output2)
	}
}
