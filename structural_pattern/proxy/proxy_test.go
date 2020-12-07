package proxy

import "testing"

func TestUserLoginProxy1(t *testing.T) {
	proxy := UserLoginProxy1{}

	got := proxy.login("Hello ", "World")
	want := "Hello World UserLoginProxy1"
	if got != want {
		t.Fatalf("Expected %q, got %q", want, got)
	}
}

func TestUserLoginProxy2(t *testing.T) {
	userLogin := UserLogin{}
	proxy := UserLoginProxy2{userLogin: &userLogin}
	got := proxy.login("Hello ", "World")
	want := "Hello World UserLoginProxy2"
	if got != want {
		t.Fatalf("Expected %q, got %q", want, got)
	}
}
