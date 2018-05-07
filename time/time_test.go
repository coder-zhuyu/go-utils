package time

import (
	"testing"
)

func TestStringToTimeStamp(t *testing.T) {
	timeStampStr := "2018-05-07 00:00:00"
	timeStamp, err := StringToTimeStamp(timeStampStr)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(timeStamp)
}

func TestStringToTimeStampLocal(t *testing.T) {
	timeStampStr := "2018-05-07 00:00:00"
	timeStamp, err := StringToTimeStampLocal(timeStampStr)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(timeStamp)
}