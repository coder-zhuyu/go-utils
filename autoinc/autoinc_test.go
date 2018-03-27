package autoinc_test

import (
	"testing"
	"github.com/coder-zhuyu/go-utils/autoinc"
)

func TestAutoInc_Id(t *testing.T) {
	var ai *autoinc.AutoInc
	ai = autoinc.New(0, 1, 100)

	for i := 0; i <= 500; i++ {
		t.Log(i, ":", ai.Id())
	}
}