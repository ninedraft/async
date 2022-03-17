package testutils

import (
	"fmt"
	"testing"
)

func RequireEq[E comparable](test testing.TB, expected, got E, message string, args ...any) {
	test.Helper()
	if expected != got {
		msg := fmt.Sprintf(message, args...)
		test.Fatalf("%s: %v is expected, got %v", msg, expected, got)
	}
}

func RequireOneOf[E comparable](test testing.TB, expected []E, got E, message string, args ...any) {
	test.Helper()
	if len(expected) == 0 {
		test.Fatalf("expected must be a non-empty slice")
	}
	for _, exp := range expected {
		if exp == got {
			return
		}
	}
	msg := fmt.Sprintf(message, args...)
	test.Fatalf("%s: one of %v is expected, got %v", msg, expected, got)
}
