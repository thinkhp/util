package thinkLog

// 2018-07-12解决:
// 1.log的Output需要每天定时更新;(使用定时器)
// 2.写入文件会覆盖之前的文件;
import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
	"util/thinkString"
)

var DebugLog thinkDebugLogger
var WarnLog thinkWarnLogger
var ErrorLog thinkErrorLogger
var SystemLog thinkSystemLogger
var LogDir = "./log/"
var DebugLogFile *os.File
var WarnLogFile *os.File
var ErrorLogFile *os.File
var SystemLogFile *os.File
var DefaultSetting = true

type thinkDebugLogger struct {
	log.Logger
}
type thinkWarnLogger struct {
	log.Logger
}
type thinkErrorLogger struct {
	log.Logger
}
type thinkSystemLogger struct {
	log.Logger
}

func init() {
	DebugLog.SetFlags(log.LstdFlags | log.Lshortfile)
	DebugLog.SetPrefix("[DEBUG] ")
	WarnLog.SetFlags(log.LstdFlags | log.Lshortfile)
	WarnLog.SetPrefix("[WARN] ")
	ErrorLog.SetFlags(log.LstdFlags | log.Lshortfile)
	ErrorLog.SetPrefix("[ERROR] ")
	SystemLog.SetFlags(log.LstdFlags | log.Lshortfile)
	SystemLog.SetPrefix("[SYSTEM] ")
	DebugLog.SetOutput(os.Stdout)
	WarnLog.SetOutput(os.Stdout)
	ErrorLog.SetOutput(os.Stderr)
	SystemLog.SetOutput(os.Stdout)
}

// 必须等待日志初始化后完成后才可以继续运行
func SetLogFileTask() {
	settingOk := make(chan bool)
	go func() {
		for true {
			setLogFileTask()
		}
	}()
	go func(setting chan bool) {
		for DefaultSetting {
			time.Sleep(time.Millisecond)
		}
		setting <- true
		defer close(setting)
	}(settingOk)
	<-settingOk
}

// 每天更改日志指向文件
func setLogFileTask() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("[recover] 日志已恢复")
		}
	}()
	duration := time.Hour * 24
	setLogFile(duration)
	next := time.Now().Add(duration)
	next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
	fmt.Println("util/thinkLog.setLogFile", "  nextTime:", next)
	time.Sleep(next.Sub(time.Now()))
}

func setLogFile(duration time.Duration) {
	var mu sync.Mutex
	// 获取之前的日志文件的指针
	var olderDebugLogFile *os.File = DebugLogFile
	var olderWarnLogFile *os.File = WarnLogFile
	var olderErrorLogFile *os.File = ErrorLogFile

	var fileNameUnSuffix string
	now := time.Now()
	switch duration {
	case time.Hour * 24: // 日志的密度以天为单位
		fileNameUnSuffix = GetTimeStringForFileName(now)[:10]
	case time.Hour: // 日志的密度以小时为单位
		fileNameUnSuffix = GetTimeStringForFileName(now)[:13]
	case time.Minute: // 日志的密度以分钟为单位
		fileNameUnSuffix = GetTimeStringForFileName(now)[:16]
	case time.Second: // 日志的密度以秒为单位
		fileNameUnSuffix = GetTimeStringForFileName(now)[:19]
	}
	if LogDir != "" {
		logFile := func(level string) *os.File {
			// 如果path指定了一个已经存在的目录，MkdirAll不做任何操作并返回nil。
			err := os.MkdirAll(LogDir, os.ModePerm)
			if err != nil {
				panic(err)
			}
			fileName := LogDir + level + fileNameUnSuffix + ".log"
			file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0766)
			if err != nil {
				panic(err)
			}
			// 本次打开的文件如果close,log日志无法写入
			// file.Close()
			return file
		}
		DebugLogFile = logFile("debug")
		WarnLogFile = logFile("debug")
		ErrorLogFile = logFile("error")

		mu.Lock()
		DebugLog.SetOutput(DebugLogFile)
		WarnLog.SetOutput(WarnLogFile)
		ErrorLog.SetOutput(ErrorLogFile)
		mu.Unlock()
	}
	if !DefaultSetting {
		olderDebugLogFile.Close()
		olderWarnLogFile.Close()
		olderErrorLogFile.Close()
	}
	DefaultSetting = false
}

func (t *thinkDebugLogger) PrintSQL(sqlString string, params []interface{}) {
	t.Println("query:", sqlString)
	printStr := "params: "
	for _, str := range params {
		printStr += fmt.Sprint(str)
		printStr += ","
	}
	thinkString.ReplaceLastRune(&printStr, 0)
	t.Println(printStr)
}

// 请求信息,请求参数
func (t *thinkDebugLogger) PrintParams(url, paramKind string, params ...string) {
	log := "\n"
	log += url + "\n"
	log += "***************************** " + paramKind + " ****************************\n"
	for i := 0; i < len(params); i++ {
		log += params[i] + "\n"
	}
	log += "***************************** " + paramKind + " ****************************"
	t.Println(log)
}

// 因为 import cycle not allowed(thinkLog, thinkTime)
func GetTimeStringForFileName(now time.Time) string {
	return now.Format("2006-01-02T15-04-05.999999999")
}
