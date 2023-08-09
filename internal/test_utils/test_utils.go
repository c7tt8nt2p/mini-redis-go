package test_utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func ShouldPanic(t *testing.T, f func()) {
	t.Helper()
	defer func() { _ = recover() }()
	f()
	t.Errorf("should have panicked but did not")
}

func ShouldPanicWithError(t *testing.T, f func(), expectedErr string) {
	t.Helper()
	defer func() {
		err := recover()
		assert.Equal(t, expectedErr, err)
	}()
	f()
	t.Errorf("should have panicked but did not")
}
