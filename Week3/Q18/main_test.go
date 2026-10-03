package main

import (
	"strings"
	"testing"
)

func TestFetchWebsite(t *testing.T) {
	// 実際のネットワークアクセスが発生するテスト
	// https://httpbin.org/get へのリクエストが成功するか確認
	content, err := FetchWebsite("https://httpbin.org/get")
	if err != nil {
		t.Fatalf("FetchWebsite failed: %v", err)
	}

	if !strings.Contains(content, "httpbin.org") {
		t.Errorf("Response does not seem to contain expected data (httpbin.org). Got: \n%s", content)
	}

	// 存在しないURLへのアクセスでエラーが返るか確認
	_, err = FetchWebsite("http://this-url-is-invalid-and-should-not-exist.local")
	if err == nil {
		t.Errorf("Expected an error for an invalid URL, but got nil")
	}
}
