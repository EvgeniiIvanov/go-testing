package even

import "testing"

func TestIsEvenTrue(t *testing.T) {
	got := IsEven(666)
	want := true

	if got != want {
		t.Errorf("IsEven(666) = %v, want %v", got, want)
	}
}

func TestIsEvenFalse(t *testing.T) {
        got := IsEven(13)
        want := false

        if got != want {
                t.Errorf("IsEven(13) = %v, want %v", got, want)
        }
}

func TestIsEvenZero(t *testing.T) {
        got := IsEven(0)
        want := true

        if got != want {
                t.Errorf("IsEven(0) = %v, want %v", got, want)
        }
}

func TestIsEvenNegativeTrue(t *testing.T) {
        got := IsEven(-666)
        want := true

        if got != want {
                t.Errorf("IsEven(-666) = %v, want %v", got, want)
        }
}

func TestIsEvenNegaitveFalse(t *testing.T) {
        got := IsEven(-13)
        want := false

        if got != want {
                t.Errorf("IsEven(-13) = %v, want %v", got, want)
        }
}

// Before this was too much for simple function

func TestIsEven(t *testing.T) {
        if !IsEven(8) {
		t.Error("We are waiting true for 8, but getting false")
	}

	if IsEven(17) {
                t.Error("We are waiting false for 17, but getting true")
        }

	if !IsEven(0) {
                t.Error("We are waiting true for 0, but getting false")
        }

	if !IsEven(-88) {
                t.Error("We are waiting true for -88, but getting false")
        }

	if IsEven(-67) {
                t.Error("We are waiting false for -67, but getting true")
        }
}
