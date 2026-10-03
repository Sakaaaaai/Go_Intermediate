package main

import (
	"encoding/json"
	"fmt"
)

// 1. User 構造体を定義し、JSONタグを設定してください
// ヒント: JSONタグは `json:"キー名"` のように書きます（バッククォートに注意）
type User struct {
	Name string ___
	Age  int    ___
}

// 2. 構造体をJSON文字列に変換する関数
func EncodeUser(u User) (string, error) {
	// ヒント: json.Marshal を使って struct(u) から JSONのバイト配列に変換します
	jsonData, err := ___
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

// 3. JSON文字列を構造体に変換する関数
func DecodeUser(jsonStr string) (User, error) {
	var u User
	// ヒント: json.Unmarshal を使って JSON文字列 から構造体に変換します
	// 第一引数はバイト配列（[]byte(jsonStr)）、第二引数は代入先へのポインタ（&u）です
	err := ___
	if err != nil {
		return User{}, err
	}
	return u, nil
}

func main() {
	originalUser := User{Name: "Gopher", Age: 10}

	jsonString, _ := EncodeUser(originalUser)
	fmt.Println("Encoded JSON:", jsonString)

	decodedUser, _ := DecodeUser(jsonString)
	fmt.Printf("Decoded Struct: %+v\n", decodedUser)
}
