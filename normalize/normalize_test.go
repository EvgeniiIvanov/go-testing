package normalize

import (
	"testing"
)

func TestClean(t *testing.T) {
	cases := []struct {
		name string
		s    string
		want string
	}{
		{"normalized string test", "already normalized string", "already normalized string"},
		{"empty string test", "", ""},
		{"several spaces and tabs test", "  there are  some\tspaces   here \t  ", "there are some spaces here"},
		{"different registers test", "We hAve diFferent Registers HerE", "we have different registers here"},
		{"non letter characters test", "1234567890!@#$%^&*()_ accidental words  1234567890!@#$%^&*()_", "1234567890!@#$%^&*()_ accidental words 1234567890!@#$%^&*()_"},
		{"punctuation marks test", "Hello! How are you? Well, it's fine. I'm just a test.", "hello! how are you? well, it's fine. i'm just a test."},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := Clean(c.s)
			if got != c.want {
				t.Errorf("Clean(%s) = %s, but we want %s", c.s, got, c.want)
			}
		})
	}
}
