package errors

import (
    "testing"
)

func f() error {
    return New(500, "internal server error")
}

func TestError_Error(t *testing.T) {
    e := f()
    t.Log(e)
}
