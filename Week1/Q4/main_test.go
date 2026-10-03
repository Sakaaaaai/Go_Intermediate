package main

import "testing"

func TestCreateAdminAndProfile(t *testing.T) {
	admin, profile := CreateAdminAndProfile()

	if admin.Role != "SuperUser" {
		t.Errorf("Admin Role = %q; want \"SuperUser\"", admin.Role)
	}
	if admin.Name != "太郎" || admin.Age != 30 {
		t.Errorf("Admin embedded User = %+v; want {Name: \"太郎\", Age: 30}", admin.User)
	}
	expectedProfile := "太郎さんは30歳です"
	if profile != expectedProfile {
		t.Errorf("Profile() = %q; want %q", profile, expectedProfile)
	}
}
