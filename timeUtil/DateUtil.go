package timeUtil

import (
	"time"
	"util/think"
)

// 方法命名规范:
// Date 	   		日期 			time.Time
// DateString		日期字符串 		string:yyyy-mm-dd
// DateTime    		时间 			time.Time
// DateTimeString   时间字符串 		string: yyyy-mm-dd hh:mm:ss
// Timestamp(S)   	时间戳(默认,秒级别)	int64
// TimestampMS   	时间戳(毫秒级别)		int64
// Time	   			go类 			time.Time
// 2006-01-02 03:04:05
// 一天的毫秒数
const oneDayMS int64 = 24 * 60 * 60 * 1000

// 一天的秒数
const oneDay int64 = 24 * 60 * 60

type FileName string

func GetTimeStringForFileName(now time.Time) FileName {
	return FileName(now.Format("20060102T150405"))
}
func (f FileName) TillDay() string {
	return string(f)[:8]
}
func (f FileName) TillHour() string {
	return string(f)[:11]
}
func (f FileName) TillMinute() string {
	return string(f)[:13]
}
func (f FileName) TillSecond() string {
	return string(f)[:15]
}

func GetDate(now time.Time, i int) time.Time {
	now = now.AddDate(0, 0, i)
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
}

// GetDate 的不同实现
// 存在时区不一致的问题
func getDate(now time.Time, i int) time.Time {
	return time.Unix((now.Unix()/oneDay+int64(i))*oneDay, 0)
}
func GetYesterday(now time.Time) time.Time {
	return GetDate(now, -1)
}

func GetYesterdayString(now time.Time) string {
	return GetYesterday(now).String()[0:10]
}

// 获取日期 yyyy-mm-dd[:10]
// 获取截取字符串
func GetDateString(time time.Time) string {
	timeString := time.String()
	return timeString[:10]
}

// 获取日期 yyyy-mm-dd hh:mm:s[:19]
func GetDateTimeString(time time.Time) string {
	time.Format("2006-01-02 03:04:05 PM")
	// time.Format("02/01/2006 15:04:05 PM")
	// time.Format("2006-01-02 15:04:05")
	timeString := time.String()
	return timeString[:19]
}

//const (
//	ANSIC       = "Mon Jan _2 15:04:05 2006"
//	UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
//	RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
//	RFC822      = "02 Jan 06 15:04 MST"
//	RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
//	RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
//	RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
//	RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
//	RFC3339     = "2006-01-02T15:04:05Z07:00"
//	RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
//	Kitchen     = "3:04PM"
//	// Handy time stamps.
//	Stamp      = "Jan _2 15:04:05"
//	StampMilli = "Jan _2 15:04:05.000"
//	StampMicro = "Jan _2 15:04:05.000000"
//	StampNano  = "Jan _2 15:04:05.000000000"
//)
//
//const (
//	_                        = iota
//	stdLongMonth             = iota + stdNeedDate  // "January"
//	stdMonth                                       // "Jan"
//	stdNumMonth                                    // "1"
//	stdZeroMonth                                   // "01"
//	stdLongWeekDay                                 // "Monday"
//	stdWeekDay                                     // "Mon"
//	stdDay                                         // "2"
//	stdUnderDay                                    // "_2"
//	stdZeroDay                                     // "02"
//	stdHour                  = iota + stdNeedClock // "15"
//	stdHour12                                      // "3"
//	stdZeroHour12                                  // "03"
//	stdMinute                                      // "4"
//	stdZeroMinute                                  // "04"
//	stdSecond                                      // "5"
//	stdZeroSecond                                  // "05"
//	stdLongYear              = iota + stdNeedDate  // "2006"
//	stdYear                                        // "06"
//	stdPM                    = iota + stdNeedClock // "PM"
//	stdpm                                          // "pm"
//	stdTZ                    = iota                // "MST"
//	stdISO8601TZ                                   // "Z0700"  // prints Z for UTC
//	stdISO8601SecondsTZ                            // "Z070000"
//	stdISO8601ShortTZ                              // "Z07"
//	stdISO8601ColonTZ                              // "Z07:00" // prints Z for UTC
//	stdISO8601ColonSecondsTZ                       // "Z07:00:00"
//	stdNumTZ                                       // "-0700"  // always numeric
//	stdNumSecondsTz                                // "-070000"
//	stdNumShortTZ                                  // "-07"    // always numeric
//	stdNumColonTZ                                  // "-07:00" // always numeric
//	stdNumColonSecondsTZ                           // "-07:00:00"
//	stdFracSecond0                                 // ".0", ".00", ... , trailing zeros included
//	stdFracSecond9                                 // ".9", ".99", ..., trailing zeros omitted
//
//	stdNeedDate  = 1 << 8             // need month, day, year
//	stdNeedClock = 2 << 8             // need hour, minute, second
//	stdArgShift  = 16                 // extra argument in high bits, above low stdArgShift
//	stdMask      = 1<<stdArgShift - 1 // mask out argument
//)

//
// 从MySQL取得的时间字符串格式为2006-01-02T15:04:05Z
func GetTimeFromSqlString(date string) time.Time {
	loc, _ := time.LoadLocation("Local")
	// ParseInLocation类似Parse但有两个重要的不同之处。
	// 第一，当缺少时区信息时，Parse将时间解释为UTC时间，而ParseInLocation将返回值的Location设置为loc；
	// 第二，当时间字符串提供了时区偏移量信息时，Parse会尝试去匹配本地时区，而ParseInLocation会去匹配loc
	t, err := time.ParseInLocation("2006-01-02T15:04:05Z", date, loc)
	think.IsNil(err)
	return t
}

// 第一个参数是格式，第二个是要转换的时间字符串
// 对于空字符串,time为零值
func GetTimeFromString(date string) time.Time {
	//loc, _ := time.LoadLocation("Local")
	loc := time.Local
	var t time.Time
	var err error
	switch len(date) {
	case 10:
		t, err = time.ParseInLocation("2006-01-02", date, loc)
	case 19:
		t, err = time.ParseInLocation("2006-01-02 15:04:05", date, loc)
	case 20:
		t, err = time.ParseInLocation("2006-01-02T15:04:05Z", date, loc)
	}
	think.IsNil(err)

	return t
}
