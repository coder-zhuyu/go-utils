package autoinc

import (
	"testing"
)

func TestAutoInc_Id(t *testing.T) {
	var ai *AutoInc
	ai = New(0, 1, 100)

	for i := 0; i <= 500; i++ {
		t.Log(i, ":", ai.Id())
	}
}