package thinkTimer

import (
	"github.com/beevik/guid"
	"sync"
	"time"
	"util/think"
)

// 任务池
// 管理内部任务的启动,暂停,终止
// 方法内部 go
type ThinkTasks struct {
	locker sync.RWMutex
	// 运行开关
	tasks map[string]*ThinkTask
	//
	startSwitch bool
}

// 任务
type ThinkTask struct {
	TaskId   int
	TaskGuid string
	TaskName string

	// 状态相关
	Flag       bool
	CreateTime time.Time
	UpdateTime time.Time
	StartTime  time.Time
	EndTime    time.Time

	//
	Params map[string]interface{}
}

func (ts *ThinkTasks) Init() *ThinkTasks {
	ts.tasks = make(map[string]*ThinkTask)
	ts.startSwitch = false

	return ts
}

func (ts *ThinkTasks) GetTasks() map[string]*ThinkTask {
	return ts.tasks
}

func (ts *ThinkTasks) AddTaskNow(nextTime string, params map[string]interface{}, f ...func()) *ThinkTask {
	// 立即运行函数f
	for i := 0; i < len(f); i++ {
		f[i]()
	}
	return ts.AddTask(nextTime, params, f...)
}

func (ts *ThinkTasks) AddTask(nextTime string, params map[string]interface{}, f ...func()) *ThinkTask {
	// 添加任务
	t := new(ThinkTask)
	t.CreateTime = time.Now()
	t.TaskName = time.Now().String()[:19]
	t.TaskGuid = guid.New().StringUpper()
	t.Params = params

	ts.setTasks(t)
	ts.setFlags(t.TaskGuid, true)
	go func() {
		for ts.getFlags(t.TaskGuid) {
			next, err := GetNextTime(nextTime)
			think.IsNil(err)
			task(*next, f...)
		}
		t.EndTime = time.Now()
		ts.delTasks(t.TaskGuid)
	}()

	return t
}

func (ts *ThinkTasks) StopTask(taskGuid string) {
	if ts.getFlags(taskGuid) {
		ts.setFlags(taskGuid, false)
	}
}

func (ts *ThinkTasks) setTasks(t *ThinkTask) {
	ts.locker.Lock()
	defer ts.locker.Unlock()
	ts.tasks[t.TaskGuid] = t
}
func (ts *ThinkTasks) getTasks(taskGuid string) *ThinkTask {
	ts.locker.RLock()
	defer ts.locker.RUnlock()

	return ts.tasks[taskGuid]
}
func (ts *ThinkTasks) delTasks(taskGuid string) {
	ts.locker.Lock()
	defer ts.locker.Unlock()

	delete(ts.tasks, taskGuid)
}
func (ts *ThinkTasks) setFlags(taskGuid string, flag bool) {
	t := ts.getTasks(taskGuid)
	if t != nil {
		t.Flag = flag
	}

}
func (ts *ThinkTasks) getFlags(taskGuid string) bool {
	t := ts.getTasks(taskGuid)
	if t != nil {
		return t.Flag
	}
	return false
}
