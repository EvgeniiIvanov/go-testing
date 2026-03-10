package safe

import (
	"testing"
)

func TestMustAtPositive(t *testing.T) {
	s := []string{"first", "second", "third"}
	i := MustAt(s, 1)
	if i != "second" {
		t.Errorf("Expected 'second' but got '%s'", i)
	}
}

func TestMustAtIndexLessZero(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("Panic is expected but not thrown")
		} else if r != "index out of range" {
			t.Errorf("We want panic with message 'index out of range' but got '%s'", r)
		}
	}()

	s := []int{1, 2, 3}
	_ = MustAt(s, -2)
}

func TestMustAtOutOfRange(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("Panic is expected but not thrown")
		} else if r != "index out of range" {
			t.Errorf("We want panic with message 'index out of range' but got '%s'", r)
		}
	}()

	s := []int{1, 2, 3}
	_ = MustAt(s, 4)
}
