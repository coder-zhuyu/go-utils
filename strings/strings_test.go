package strings

import (
	"testing"
)

func TestRuneLen(t *testing.T) {
	s := "你好，世界!"
	lth := RuneLen(s)
	t.Log("len=", lth)
}

func TestRuneSubString(t *testing.T) {
	s := "你好，世界!"
	t.Log("sub=", RuneSubString(s, 0, 6))
	t.Log("sub=", RuneSubString(s, 0, 10))
	t.Log("sub=", RuneSubString(s, 7, 10))
}
