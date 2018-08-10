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

func setLogFileTask() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("[recover] 日志已恢复")
		}
	}()
	setLogFile()
	time.Sleep(time.Second)
	now := time.Now()
	next := now.AddDate(0, 0, 1)
	next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
	fmt.Println("util/thinkLog.setLogFile", "  nextTime:", next)
	time.Sleep(next.Sub(time.Now()))
}

func setLogFile() {
	var mu sync.Mutex
	// 获取之前的日志文件的指针
	var tempDebugLogFile *os.File = DebugLogFile
	var tempWarnLogFile *os.File = WarnLogFile
	var tempErrorLogFile *os.File = ErrorLogFile
	// 日志的密度以天为单位
	fileNameUnSuffix := time.Now().String()[:10]
	// 日志密度以分钟为单位
	//getTimeString := func() string{
	//	now := time.Now()
	//	return now.Format("20060102150405")
	//}
	//fileNameUnSuffix := getTimeString()
	//nextTime := "* * * * * 0"
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
		tempDebugLogFile.Close()
		tempWarnLogFile.Close()
		tempErrorLogFile.Close()
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
