package textstat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordCount(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want map[string]int
	}{
		{"ascii basic test", "yes, No, yes, Yes", map[string]int{"yes": 3, "no": 1}},
		{"ascii + numbers + spaces test", "  fuz1   buz2  fuz1 buz2   ", map[string]int{"fuz1": 2, "buz2": 2}},
		{"symbol separator test", "yes&Yes$#%No;no", map[string]int{"yes": 2, "no": 2}},
		{"empty string test", "", map[string]int{}},
		{"separators only test", "  ;,.{}!$?/  ", map[string]int{}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := WordCount(c.in)
			assert.Equal(t, c.want, got)
		})
	}
}
