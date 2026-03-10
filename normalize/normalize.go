package normalize

import "strings"

func Clean(s string) string {
	tokens := strings.Fields(s)
	if len(tokens) == 0 {
		return ""
	}
	return strings.ToLower(strings.Join(tokens, " "))
}
