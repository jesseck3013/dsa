package testutils

import "testing"

func AssertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("expected no error, but got %v", err)
	}
}

func AssertError(t *testing.T, want, got error) {
	t.Helper()
	if want != got {
		t.Errorf("expected error: %v, got %v", want, got)
	}
}

func AssertValue[T comparable](t *testing.T, want, got T) {
	t.Helper()
	if got != want {
		t.Errorf("expected %v, got %v", want, got)
	}
}
