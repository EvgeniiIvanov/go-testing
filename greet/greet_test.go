package greet

import (
	"testing"
)

func assertString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("We get %q, but want %q", got, want)
	}
}

func assertError(t testing.TB, got error, want bool) {
	t.Helper()
	if want == (got == nil) {
		if !want {
			t.Fatalf("We get an error %q, but it is not expected", got)
		} else {
			t.Fatal("The error was expected but we didn't get it")
		}
	}
}

func TestHello(t *testing.T) {
	cases := []struct {
		testName  string
		inputName string
		wantName  string
		wantError bool
	}{
		{"basic case #1", "Go", "Hello, Go", false},
		{"basic case #2", "World", "Hello, World", false},
		{"unicode characters", "Гофер", "Hello, Гофер", false},
		{"save spaces", " Guy  ", "Hello,  Guy  ", false},
		{"empty string", "", "", true},
	}

	for _, c := range cases {
		t.Run(c.testName, func(t *testing.T) {
			got, err := Hello(c.inputName)
			assertError(t, err, c.wantError)
			assertString(t, got, c.wantName)
		})
	}
}
