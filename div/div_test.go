package div

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDiv(t *testing.T) {
	t.Run("Division by zero", func(t *testing.T) {
		result, err := Div(10, 0)

		require.Error(t, err)
		require.ErrorIs(t, err, errorDivisionByZero)
		assert.Equal(t, 0, result)
	})

	cases := []struct {
		name           string
		a, b, expected int
	}{
		{"Positive", 10, 2, 5},
		{"Negative", -9, 3, -3},
		{"Division by one", 10, 1, 10},
		{"0 division", 0, 10, 0},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			result, err := Div(c.a, c.b)

			require.NoError(t, err)
			assert.Equal(t, c.expected, result)
		})
	}
}
