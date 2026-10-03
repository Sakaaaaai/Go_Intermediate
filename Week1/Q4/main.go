package main

import "fmt"

// 1. User 構造体を定義してください
type User struct {
	Name string
	Age  int
}

// 2. User に Profile() メソッドを定義してください
// ヒント: メソッド名は Profile、戻り値は string で "<Name>さんは<Age>歳です" を返します
func ___ {
	return fmt.Sprintf("%sさんは%d歳です", ___, ___)
}

// 3. Admin 構造体に User を埋め込んでください
// ヒント: 型名だけを書くと埋め込みになります
type Admin struct {
	___
	Role string
}

func CreateAdminAndProfile() (Admin, string) {
	// 4. Admin のインスタンスを作成
	admin := Admin{
		User: User{Name: "太郎", Age: 30},
		Role: "SuperUser",
	}

	// 埋め込まれた User の Profile() を呼び出すプロモーション機能を利用してください
	// ヒント: admin.Profile() のように呼び出します
	profile := ___
	
	return admin, profile
}

func main() {
	admin, profile := CreateAdminAndProfile()
	fmt.Println("Role:", admin.Role)
	fmt.Println("Profile:", profile)
}
