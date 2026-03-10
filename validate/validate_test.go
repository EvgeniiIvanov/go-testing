package validate

import (
	"testing"
)

// func TestValidateNamePossitive(t *testing.T)
func TestValidateNamePossitive(t *testing.T) {
	if err := ValidateName("Non Empty String"); err != nil {
		t.Errorf("We get an error %q, but we pass string 'Non Empty String'", err)
	}
}

func TestValidateNameEmptyString(t *testing.T) {
	err := ValidateName("")
	if err == nil {
		t.Fatal("Error is expected but not returned")
	}
	want := "empty name"
	if err.Error() != want {
		t.Errorf("We get %q, but want to get %q", err, want)
	}
}
