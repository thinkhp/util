package thinkLog

import (
	"testing"
	"time"
	"strconv"
)



func TestSetLogFileTask(t *testing.T) {
	var i int
	SetLogFileTask()
	for true {
		i++
		DebugLog.Println(`test`, strconv.Itoa(i))
		time.Sleep(time.Millisecond * 100)
	}
}
