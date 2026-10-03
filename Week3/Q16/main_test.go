package main

import (
	"strings"
	"testing"
)

func TestJSONEncodeDecode(t *testing.T) {
	user := User{Name: "TestUser", Age: 25}

	// エンコードのテスト
	encoded, err := EncodeUser(user)
	if err != nil {
		t.Fatalf("EncodeUser failed: %v", err)
	}

	if !strings.Contains(encoded, `"name":"TestUser"`) || !strings.Contains(encoded, `"age":25`) {
		t.Errorf("Encoded JSON does not seem correct: %s", encoded)
	}

	// デコードのテスト
	decoded, err := DecodeUser(encoded)
	if err != nil {
		t.Fatalf("DecodeUser failed: %v", err)
	}

	if decoded.Name != "TestUser" || decoded.Age != 25 {
		t.Errorf("Decoded User is incorrect, got %+v", decoded)
	}
}
