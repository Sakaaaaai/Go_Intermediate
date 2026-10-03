package main

import "testing"

func TestIsEven(t *testing.T) {
	tests := []struct {
		input    int
		expected bool
	}{
		{2, true},
		{3, false},
		{0, true},
		{-4, true},
		{-5, false},
	}

	for _, tt := range tests {
		// 1. IsEven(tt.input) を呼び出して結果を got に代入してください
		// ヒント: テスト対象の関数 IsEven を呼び出します
		___ := ___

		// 2. got と tt.expected を比較し、異なる場合に t.Errorf でエラーを出力してください
		// ヒント: t.Errorf("IsEven(%d) = %v; want %v", tt.input, got, tt.expected) のように書きます
		if ___ {
			___
		}
	}
}
