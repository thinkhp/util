package thinkTimer

import (
	"errors"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"
	"util/think"
	"util/thinkLog"
)

// 立即执行
func TaskNow(nextTime string, f ...func()) {
	for i := 0; i < len(f); i++ {
		f[i]()
	}
	for true {
		task(nextTime, f...)
	}

}
func Task(nextTime string, f ...func()) {
	for true {
		task(nextTime, f...)
	}
}

func task(nextTime string, f ...func()) {
	defer func() {
		if r := recover(); r != nil {
			thinkLog.ErrorLog.Println("[recover] 程序已恢复")
		}
	}()
	now := time.Now()
	next, err := GetNextTime(nextTime)
	think.Check(err)
	for i := 0; i < len(f); i++ {
		funcName := runtime.FuncForPC(reflect.ValueOf(f[i]).Pointer()).Name()
		thinkLog.DebugLog.Println(funcName, "\t\t nextTime:", next)
	}
	time.Sleep(next.Sub(now))
	for i := 0; i < len(f); i++ {
		f[i]()
	}
}

// 堆叠second来计算nextTime,效率奇差无比
func GetNextTimeByAddSecond(now time.Time, timeSlice []int) *time.Time {
	for {
		now = now.Add(time.Second)
		//fmt.Println(now)
		if isFormat(now, timeSlice) {
			return &now
		}
	}
}

func GetNextTime(expression string) (*time.Time, error) {
	now := time.Now()
	// 由表达式("* * * * * *")获取[]int
	timeSlice := getDate(expression)
	// 参数校验
	if len(timeSlice) != 6 {
		return nil, errors.New("表达式格式错误")
	}
	if timeSlice[5] != -999 && now.Year() > timeSlice[5] {
		return nil, errors.New("表达式年数错误")
	}
	fixLen, next := getInitTimeSlice(timeSlice, now)
	//fmt.Println("next:",next)
	//fmt.Println("fixLen:",fixLen)
	nextFixed := next
	if fixLen == 6 {
		next = next.Add(time.Second)
		return &next, nil
	}
	if fixLen == 0 {
		return &next, nil
	}
	for next.Sub(now) < 0 {
		// 找到可以未固定(-999)的 timeUnit
		i := 0
		for i = 0; i < len(timeSlice); i++ {
			// !-999 原始被固定
			// -888 已经被固定了,防止再次被固定
			if timeSlice[i] == -999 {
				timeSlice[i] = -888
				break
			}
		}
		//fmt.Println("timeSlice:", timeSlice, " i:", i)
		// 根据 i 的值确定要改变的 timeUnit
		switch i {
		case 0:
			{
				for next.Sub(now) < 0 {
					next = next.Add(time.Second)
					// 产生进位:回滚,本位置于init,break
					if next.Minute() != nextFixed.Minute() {
						next = next.Add(-time.Minute)
						next = time.Date(next.Year(), next.Month(), next.Day(), next.Hour(), next.Minute(), 0, 0, next.Location())
						break
					}
				}
				// 进位或者

			}
		case 1:
			{
				for next.Sub(now) < 0 {
					next = next.Add(time.Minute)
					// 产生进位:回滚,本位置于init,break
					if next.Hour() != nextFixed.Hour() {
						next = next.Add(-time.Hour)
						next = time.Date(next.Year(), next.Month(), next.Day(), next.Hour(), 0, next.Second(), 0, next.Location())
						break
					}
				}
			}
		case 2:
			{
				for next.Sub(now) < 0 {
					next = next.Add(time.Hour)
					// 产生进位:回滚,本位置于init,break
					if next.Day() != nextFixed.Day() {
						next = next.AddDate(0, 0, -1)
						next = time.Date(next.Year(), next.Month(), next.Day(), 0, next.Minute(), next.Second(), 0, next.Location())
						break
					}
				}
			}
		case 3:
			{
				for next.Sub(now) < 0 {
					next = next.AddDate(0, 0, 1)
					// fmt.Println("afterOneDay:", next)
					// 产生进位:回滚,本位置于init,break
					if next.Month() != nextFixed.Month() {
						next = next.AddDate(0, -1, 0)
						next = time.Date(next.Year(), next.Month(), 1, next.Hour(), next.Minute(), next.Second(), 0, next.Location())
						break
					}
				}
			}
		case 4:
			{
				for next.Sub(now) < 0 {
					next = next.AddDate(0, 1, 0)
					// 产生进位:回滚,本位置于init,break
					if next.Year() != nextFixed.Year() {
						next = next.AddDate(-1, 0, 0)
						next = time.Date(next.Year(), 1, next.Day(), next.Hour(), next.Minute(), next.Second(), 0, next.Location())
						break
					}
				}
			}
		case 5:
			{
				for next.Sub(now) < 0 {
					next = next.AddDate(1, 0, 0)
				}
			}
		default:
			return nil, errors.New("outOfTimeRange 或者 时间格式错误")
		}
	}
	return &next, nil
}

// 初始化时间
func getInitTimeSlice(timeSlice []int, now time.Time) (int, time.Time) {
	// 非* 与 非* 之间要填充为 initTimeUnit
	// 非 * 位填写int[]对应的值
	// 其他位与 now 一致
	fixLen := 0
	initTimeSlice := make([]int, 0, len(timeSlice))
	startFlag := false
	start := -1
	end := -1
	for i := 0; i < len(timeSlice); i++ {
		if timeSlice[i] == -999 {
			// 计算固定位(数字位)的个数
			fixLen++
			if start != -1 {
				startFlag = true
			}
		} else {
			if !startFlag {
				start = i
			} else {
				end = i
			}
		}
		initTimeSlice = append(initTimeSlice, timeSlice[i])
		// 如果前后均有 flag 且不相邻,则中途所有的数字置为 init
		if start != -1 && end != -1 && (end-start) > 1 {
			for i := start + 1; i < end; i++ {
				initTimeSlice[i] = initTimeUnit(i)
			}
			// 状态回复至start刚开始时
			startFlag = false
			start = end
			end = -1
		}
	}
	var year, month, day, hour, minute, second *timeUnit
	second = getTimeUnit(initTimeSlice, 0, now)
	minute = getTimeUnit(initTimeSlice, 1, now)
	hour = getTimeUnit(initTimeSlice, 2, now)
	day = getTimeUnit(initTimeSlice, 3, now)
	month = getTimeUnit(initTimeSlice, 4, now)
	year = getTimeUnit(initTimeSlice, 5, now)

	next := time.Date(year.num, time.Month(month.num), day.num, hour.num, minute.num, second.num, 0, now.Location())
	//next = time.Date(next.Year(),next.Month(),next.Day(),next.Hour(),next.Minute(),next.Second(),0,next.Location())

	//fmt.Println("second:",second)
	//fmt.Println("minute:",minute)
	//fmt.Println("hour:",hour)
	//fmt.Println("day:",day)
	//fmt.Println("month:",month)
	//fmt.Println("year:",year)

	return fixLen, next
}

func getDate(str string) []int {
	slice := strings.Split(str, " ")
	resultSlice := make([]int, 0)
	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i] == "*" {
			resultSlice = append(resultSlice, -999)
		} else {
			temp, err := strconv.Atoi(slice[i])
			if err != nil {
				panic(err)
			}
			resultSlice = append(resultSlice, temp)
		}
	}
	// fmt.Println("resultSlice",resultSlice)
	return resultSlice
}

type timeUnit struct {
	isFixed bool
	num     int
}

func getTimeUnit(timeSlice []int, index int, now time.Time) *timeUnit {
	var timeUnit = new(timeUnit)
	if timeSlice[index] == -999 {
		timeUnit.isFixed = false
		switch index {
		case 5:
			timeUnit.num = now.Year()
		case 4:
			timeUnit.num = int(now.Month())
		case 3:
			timeUnit.num = now.Day()
		case 2:
			timeUnit.num = now.Hour()
		case 1:
			timeUnit.num = now.Minute()
		case 0:
			timeUnit.num = now.Second()
		}
	} else {
		timeUnit.isFixed = true
		timeUnit.num = timeSlice[index]
	}
	return timeUnit
}

func isFormat(now time.Time, timeSlice []int) bool {
	if timeSlice[0] != -999 && now.Second() != timeSlice[0] {
		return false
	}
	if timeSlice[1] != -999 && now.Minute() != timeSlice[1] {
		return false
	}
	if timeSlice[2] != -999 && now.Hour() != timeSlice[2] {
		return false
	}
	if timeSlice[3] != -999 && now.Day() != timeSlice[3] {
		return false
	}
	if timeSlice[4] != -999 && int(now.Month()) != timeSlice[4] {
		return false
	}
	if timeSlice[5] != -999 && now.Year() != timeSlice[5] {
		return false
	}
	return true
}

func initTimeUnit(index int) int {
	switch index {
	case 0, 1, 2: // 时分秒的初始时间
		return 0
	case 3, 4, 5: // 年月日的初始时间
		return 1
	}
	return -1
}
