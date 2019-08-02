package timeUtil

import (
	"fmt"
	"testing"
	"time"
)

func TestGetTimeStringForFileName(t *testing.T) {
	var test = func(duration time.Duration) {
		var fileNameUnSuffix string
		now := time.Now()
		switch duration {
		case time.Hour * 24:
			fileNameUnSuffix = GetTimeStringForFileName(now).TillDay()
		case time.Hour:
			fileNameUnSuffix = GetTimeStringForFileName(now).TillHour()
		case time.Minute:
			fileNameUnSuffix = GetTimeStringForFileName(now).TillMinute()
		case time.Second:
			fileNameUnSuffix = GetTimeStringForFileName(now).TillSecond()
		}

		fmt.Println(fileNameUnSuffix)
	}

	test(time.Hour * 24)
	test(time.Hour)
	test(time.Minute)
	test(time.Second)

}


func TestGetTimeFromString(t *testing.T) {
	fmt.Println(GetTimeFromString(""))
}
