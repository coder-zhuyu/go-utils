package time

import (
	tm "time"
)

func StringToTimeStamp(timeStampStr string) (tm.Time, error) {
	timeStamp, err := tm.Parse("2006-01-02 15:04:05", timeStampStr)
	return timeStamp, err
}

func StringToTimeStampLocal(timeStampStr string) (tm.Time, error) {
	timeStamp, err := tm.ParseInLocation("2006-01-02 15:04:05", timeStampStr, tm.Local)
	return timeStamp, err
}
