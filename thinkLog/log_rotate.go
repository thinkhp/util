package thinkLog

import (
	"github.com/pkg/errors"
	"os"
	"path"
	"time"
)

const Day = time.Hour * 24
const Hour = time.Hour
const Minute = time.Minute
const Second = time.Second

var allLogFile = make(map[level]*os.File)

func SetLogRotateTask(d time.Duration) {
	allLogRotate(d)
	go logRotateTask(d)
}

// 必须等待日志初始化后完成后才可以继续运行
func SetLogFileTask() {
	//duration := time.Hour * 24
	//duration := time.Minute
	d := time.Hour * 24
	allLogRotate(d)
	go logRotateTask(d)
}

// 每天更改日志指向文件
func logRotateTask(duration time.Duration) {
	for true {
		// 设定输出频率:每天
		next := getNextTime(duration)
		TraceLog.Println("util/thinkLog.setLogFile", "  nextTime:", next)
		time.Sleep(next.Sub(time.Now()))

		allLogRotate(duration)
	}
}

func allLogRotate(duration time.Duration){
	// create all new output file
	// get all redirect level
	reLvs := make(map[level]bool)
	for _, l := range allLog {
		if l.(*lLogger).Logger.Writer() == stdNull {
			continue
		}
		reLvs[l.(*lLogger).redirectLv] = true
	}
	// create all redirect level file
	for k, _ := range reLvs {
		allLogFile[k] = createNewOutput(k, duration)
	}
	for _, l := range allLog {
		// set new output and close old output
		l.(*lLogger).logRotate()
	}
}


func createNewOutput(redirectLv level,duration time.Duration) *os.File {
	// new output
	fileNameUnSuffix := getFileNamePrefix(duration)
	fileName := path.Join(cfgLogPath, getLevelMsg(redirectLv)+fileNameUnSuffix+".log")
	newOutput, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0660)
	if err != nil {
		ErrorLog.Printf("%+v", errors.Wrap(err, "open new logger file err:"))
		return nil
	}
	TraceLog.Println("open new log file:lv", getLevelMsg(redirectLv), newOutput.Name())
	// 本次打开的文件如果close,log日志无法写入
	// newOutput.Close()

	return newOutput
}

func (l *lLogger) logRotate() {
	oldOutput := l.Writer()
	defer func() {
		if r := recover(); r != nil {
			l.SetOutput(oldOutput)
			TraceLog.Printf("[recover] 日志已恢复,文件均指向标准输出%+v\n", r)
		}
	}()
	if oldOutput == stdNull {
		return
	}
	// set new output and close old output
	l.SetOutput(allLogFile[l.redirectLv])
	switch oldOutput {
	case os.Stdout, os.Stderr, stdNull:
	default:
		f := oldOutput.(*os.File)
		TraceLog.Println("close log file:lv", getLevelMsg(l.lv), f.Name())
		f.Close()
	}
}

func getNextTime(duration time.Duration) time.Time {
	next := time.Now().Add(duration)
	switch duration {
	case Day: // 日志的密度以天为单位
		next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
	case Hour: // 日志的密度以小时为单位
		next = time.Date(next.Year(), next.Month(), next.Day(), next.Hour(), 0, 0, 0, next.Location())
	case Minute: // 日志的密度以分钟为单位
		next = time.Date(next.Year(), next.Month(), next.Day(), next.Hour(), next.Minute(), 0, 0, next.Location())
	case Second: // 日志的密度以秒为单位
		next = time.Date(next.Year(), next.Month(), next.Day(), next.Hour(), next.Minute(), next.Second(), 0, next.Location())
	}

	return next
}

func getFileNamePrefix(duration time.Duration) string {
	var fileNameUnSuffix string
	now := time.Now()
	switch duration {
	case Day: // 日志的密度以天为单位
		fileNameUnSuffix = now.Format(time.RFC3339)[:10]
	case Hour: // 日志的密度以小时为单位
		fileNameUnSuffix = now.Format(time.RFC3339)[:13]
	case Minute: // 日志的密度以分钟为单位
		fileNameUnSuffix = now.Format(time.RFC3339)[:16]
	case Second: // 日志的密度以秒为单位
		fileNameUnSuffix = now.Format(time.RFC3339)[:19]
	}
	return fileNameUnSuffix
}

func getLevelMsg(l level) string {
	switch l {
	case TraceLevel:
		return "trace"
	case DebugLevel:
		return "debug"
	case InfoLevel:
		return "info"
	case WarnLevel:
		return "warn"
	case ErrorLevel, FatalLevel, PanicLevel:
		return "error"
	default:
		return ""
	}
}
