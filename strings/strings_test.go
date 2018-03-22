package strings_test

import (
	"testing"
	strUtils "github.com/coder-zhuyu/go-utils/strings"
)

func TestRuneLen(t *testing.T) {
	s := "你好，世界!"
	lth := strUtils.RuneLen(s)
	t.Log("len=", lth)
}

func TestRuneSubString(t *testing.T) {
	s := "你好，世界!"
	t.Log("sub=", strUtils.RuneSubString(s, 0, 6))
	t.Log("sub=", strUtils.RuneSubString(s, 0, 10))
	t.Log("sub=", strUtils.RuneSubString(s, 7, 10))
}
