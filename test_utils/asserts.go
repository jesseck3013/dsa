package testutils

import "testing"

func AssertNoError(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
		panic("expected no err")
	}
}

func AssertError(t *testing.T, want, got error) {
	if want != got {
		t.Errorf("expected error: %v, got %v", want, got)
	}
}

func AssertValue[T comparable](t *testing.T, want, got T) {
	if got != want {
		t.Errorf("expected %v, got %v", want, got)
	}
}
