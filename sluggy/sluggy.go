package sluggy

import (
	"strings"
	"unicode"
)

// Slug нормализует строку для использования в URL.
func Slug(s string) string {
	s = strings.ToLower(s)
	var b strings.Builder
	prevHyphen := false
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			b.WriteRune(r)
			prevHyphen = false
			continue
		}
		if !prevHyphen {
			b.WriteByte('-')
			prevHyphen = true
		}
	}
	out := b.String()
	out = strings.Trim(out, "-")
	// Схлопывание уже обеспечено логикой prevHyphen
	return out
}
