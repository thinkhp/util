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
	// 设定日志格式
	DebugLog.SetFlags(log.LstdFlags | log.Lshortfile)
	DebugLog.SetPrefix("[DEBUG] ")
	WarnLog.SetFlags(log.LstdFlags | log.Lshortfile)
	WarnLog.SetPrefix("[WARN] ")
	ErrorLog.SetFlags(log.LstdFlags | log.Lshortfile)
	ErrorLog.SetPrefix("[ERROR] ")
	SystemLog.SetFlags(log.LstdFlags | log.Lshortfile)
	SystemLog.SetPrefix("[SYSTEM] ")
	// 指定默认输出
	defaultOutput()
}

func defaultOutput() {
	DefaultSetting = true
	DebugLog.SetOutput(os.Stdout)
	WarnLog.SetOutput(os.Stdout)
	ErrorLog.SetOutput(os.Stderr)
	SystemLog.SetOutput(os.Stdout)
}

// 必须等待日志初始化后完成后才可以继续运行
func SetLogFileTask() {
	DefaultSetting = false
	duration := time.Hour * 24
	//duration := time.Minute
	setLogFile(duration)
	go setLogFileTask(duration)
}

// 每天更改日志指向文件
func setLogFileTask(duration time.Duration) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("[recover] 日志已恢复,文件均指向标准输出")
			// 取消日志
			defaultOutput()
		}
	}()
	for true {
		// 设定输出频率:每天
		next := time.Now().Add(duration)
		switch duration {
		case time.Hour * 24: // 日志的密度以天为单位
			next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
		case time.Hour: // 日志的密度以小时为单位
			next = time.Date(next.Year(), next.Month(), next.Day(), next.Hour(), 0, 0, 0, next.Location())
		case time.Minute: // 日志的密度以分钟为单位
			next = time.Date(next.Year(), next.Month(), next.Day(), next.Hour(), next.Minute(), 0, 0, next.Location())
		case time.Second: // 日志的密度以秒为单位
			next = time.Date(next.Year(), next.Month(), next.Day(), next.Hour(), next.Minute(), next.Second(), 0, next.Location())
		}
		fmt.Println("util/thinkLog.setLogFile", "  nextTime:", next)
		time.Sleep(next.Sub(time.Now()))
		setLogFile(duration)
	}
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
	DebugLogFile = logFile("debug", fileNameUnSuffix)
	WarnLogFile = logFile("debug", fileNameUnSuffix)
	ErrorLogFile = logFile("error", fileNameUnSuffix)

	mu.Lock()
	DebugLog.SetOutput(DebugLogFile)
	WarnLog.SetOutput(WarnLogFile)
	ErrorLog.SetOutput(ErrorLogFile)
	mu.Unlock()

	//
	if !DefaultSetting {
		defer olderDebugLogFile.Close()
		defer olderWarnLogFile.Close()
		defer olderErrorLogFile.Close()
	}
}

// 因为 import cycle not allowed(thinkLog, thinkTime)
func GetTimeStringForFileName(now time.Time) string {
	return now.Format("2006-01-02T15-04-05.999999999")
}

// 获取日志文件的指针*os.File
func logFile(level string, fileNameUnSuffix string) *os.File {
	// 如果path指定了一个已经存在的目录，MkdirAll不做任何操作并返回nil。
	err := os.MkdirAll(LogDir, os.ModePerm)
	if err != nil {
		panic(err)
	}
	// e.g: ./log/debug20180304123456.log
	fileName := LogDir + level + fileNameUnSuffix + ".log"
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	// 本次打开的文件如果close,log日志无法写入
	// file.Close()
	return file
}

// 在Linux中,因为不存在文件保护,所以要检查文件名所对应的指针是否存在
func checkFile(filename string) {
	_, err := os.Stat(filename)
	if err != nil {
		panic(err)
	}
}

//Deprecated
func (t *thinkDebugLogger) PrintSQL(sqlString string, params []interface{}) {
	printStr := "params: "
	for _, str := range params {
		printStr += fmt.Sprint(str)
		printStr += ","
	}
	thinkString.ReplaceLastRune(&printStr, 0)
	t.Println("query:", sqlString)
	t.Println(printStr)
}

//Deprecated
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


