package textutil

import (
	"bufio"
	"io"
	"strings"
)

// NormalizeSpaces схлопывает последовательности пробельных символов в один пробел
// и обрезает крайние пробелы.
func NormalizeSpaces(s string) string {
	tokens := strings.Fields(s)
	return strings.Join(tokens, " ")
}

// CountLines считает количество строк в ридере (по разделителю \n).
func CountLines(r io.Reader) (int, error) {
	var n int
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		n++
	}
	return n, scanner.Err()
}
