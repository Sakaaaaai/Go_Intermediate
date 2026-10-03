package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func EncodeUser(u User) (string, error) {
	jsonData, err := json.Marshal(u)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

func DecodeUser(jsonStr string) (User, error) {
	var u User
	err := json.Unmarshal([]byte(jsonStr), &u)
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
