package textutil

import (
	"io"
	"strings"
	"testing"
)

func assertStringEqual(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("We got %q, but want %q", got, want)
	}
}

func assertIntEqual(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("We got %q, but want %q", got, want)
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestNormalizeSpaces(t *testing.T) {
	cases := []struct {
		name     string
		original string
		want     string
	}{
		{"Empty string", "", ""},
		{"Spaces and tabs only", "  \t  \t \t", ""},
		{"Remove spaces inside, before and after", " Something   happenes ", "Something happenes"},
		{"Normolized text with punctuations", "Hello, World!", "Hello, World!"},
		{"Unicode string with punctuations", " Готов  поспорить, \t что тест  пройден!\t", "Готов поспорить, что тест пройден!"},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			got := NormalizeSpaces(c.original)
			assertStringEqual(t, got, c.want)
		})
	}
}

func TestCountLines(t *testing.T) {
	cases := []struct {
		name string
		line string
		want int
	}{
		{"0 lines", "", 0},
		{"1 empty line", "\n", 1},
		{"Simple 3 lines", "line1\nline2\nline3", 3},
		{"Simple 3 lines with new line at the end", "line1\nline2\nline3\n", 3},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			r := strings.NewReader(c.line)
			cl, err := CountLines(r)

			assertNoError(t, err)
			assertIntEqual(t, cl, c.want)

		})
	}
}

func TestCountLines_ReadError(t *testing.T) {
	r := &errorReader{err: io.ErrUnexpectedEOF}
	_, err := CountLines(r)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestCountLines_LineTooLong(t *testing.T) {
	// Строка длиннее 64KB — превышает дефолтный буфер Scanner
	longLine := strings.Repeat("a", 70*1024) + "\n"
	r := strings.NewReader(longLine)
	_, err := CountLines(r)
	if err == nil {
		t.Fatal("expected buffer too long error")
	}
}

type errorReader struct {
	err error
}

func (e *errorReader) Read(p []byte) (n int, err error) {
	return 0, e.err
}
