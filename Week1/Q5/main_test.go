package main

import (
	"fmt"
	"testing"
)

func TestBookStringer(t *testing.T) {
	b := Book{Title: "テスト本", Author: "テスト著者"}

	expected := "『テスト本』(著: テスト著者)"
	
	// String() メソッドを直接呼び出して確認
	if got := b.String(); got != expected {
		t.Errorf("b.String() = %q; want %q", got, expected)
	}

	// fmt.Sprintf("%s", b) でも確認（%v ではなく %s や %v で String() が呼ばれるか）
	if got := fmt.Sprintf("%v", b); got != expected {
		t.Errorf("fmt.Sprintf(%%v) = %q; want %q", got, expected)
	}
}
