package greeting

import "testing"

// func TestGreetingYourName(t *testing.T) {
// 	given := "Hello, Bob"
// 	want := "Hello, Bob"

// 	get := Greet(given)

// 	if want != get {
// 		t.Errorf("give a name %q want greeting %q, but got %q", given, want, get)
// 	}
// }

func TestGreetingMyfriend(t *testing.T) {
	given := ""
	want := "Hello, my friend."

	get := Greet(given)

	if want != get {
		t.Errorf("sawadee krub")
	}
}

func TestGreetingCaptital(t *testing.T) {
	given := "BOB"
	want := "HELLO, BOB."

	get := Greet(given)

	if want != get {
		t.Errorf("sawadee krub")
	}
}
