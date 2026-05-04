package concurrency

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandlePanicError_NoPanic(t *testing.T) {
	assert.NotPanics(t, func() {
		defer HandlePanicError("no-panic")
	})
}

func TestHandlePanicError_StringPanic(t *testing.T) {
	assert.NotPanics(t, func() {
		defer HandlePanicError("string-panic")
		panic("something went wrong")
	})
}

func TestHandlePanicError_ErrorPanic(t *testing.T) {
	assert.NotPanics(t, func() {
		defer HandlePanicError("error-panic")
		panic(errors.New("an error"))
	})
}
