package thinkTimer

import (
	"fmt"
	"testing"
	"time"
)

func TestTask(t *testing.T){
	ts := New()
	task := ts.AddTask("* * * * * *", printTask)

	fmt.Println(ts.flags)
	time.Sleep(time.Second * 10)
	ts.StopTask(task.TaskGuid)

	fmt.Println(ts.flags)
	time.Sleep(time.Second * 10)
	fmt.Println(ts.flags)
}

func printTask(){
	println(time.Now().String()[:19])
}
