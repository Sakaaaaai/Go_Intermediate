package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func (u User) Profile() string {
	return fmt.Sprintf("%sさんは%d歳です", u.Name, u.Age)
}

type Admin struct {
	User
	Role string
}

func CreateAdminAndProfile() (Admin, string) {
	admin := Admin{
		User: User{Name: "太郎", Age: 30},
		Role: "SuperUser",
	}

	profile := admin.Profile()
	
	return admin, profile
}

func main() {
	admin, profile := CreateAdminAndProfile()
	fmt.Println("Role:", admin.Role)
	fmt.Println("Profile:", profile)
}
