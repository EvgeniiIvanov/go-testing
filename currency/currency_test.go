package currency

import "testing"

type mockConverterStatic struct {
	lastAmount float64
	lastFrom   string
	lastTo     string
	calls      int
}

func (m *mockConverterStatic) Convert(amount float64, from, to string) float64 {
	m.calls++
	m.lastAmount = amount
	m.lastFrom = from
	m.lastTo = to
	return 42.0
}

func TestCurrency(t *testing.T) {
	converter := mockConverterStatic{}

	t.Run("Basic scenario", func(t *testing.T) {
		result := PriceIn(12.0, "USD", "RUB", &converter)
		if result != 42.0 {
			t.Errorf("Got %0.2f, but wanted 42.0", result)
		}
		if converter.calls != 1 {
			t.Errorf("Got %d, but we call 1 time", converter.calls)
		}
	})

	t.Run("Args consistency case", func(t *testing.T) {
		PriceIn(100.0, "EUR", "USD", &converter)
		if converter.lastAmount != 100.0 {
			t.Errorf("Got %v, but wanted 100.0", converter.lastAmount)
		}
		if converter.lastFrom != "EUR" {
			t.Errorf("Got %q, but wanted 'EUR'", converter.lastFrom)
		}
		if converter.lastTo != "USD" {
			t.Errorf("Got %q, but wanted 'USD", converter.lastTo)
		}
	})

	t.Run("Negative amount case", func(t *testing.T) {
		PriceIn(-12.0, "USD", "RUB", &converter)
		if converter.lastAmount != -12.0 {
			t.Errorf("Got %0.2f, but wanted -12.0", converter.lastAmount)
		}
	})

	t.Run("Zero amount case", func(t *testing.T) {
		PriceIn(0.0, "USD", "RUB", &converter)
		if converter.lastAmount != 0.0 {
			t.Errorf("Got %0.2f, but wanted 0.0", converter.lastAmount)
		}
	})
}
