package timeUtil

import (
	"fmt"
	"testing"
	"time"
	"log"
)

func TestTimeStamp(t *testing.T) {
	now := time.Now()
	sec := now.Unix()
	nsec := now.UnixNano()
	fmt.Println(sec, nsec)
	fmt.Println(time.Unix(0, nsec).String())
}

func TestParseTimestamp(t *testing.T) {
	//1563180381 1563180381 831919800
	//nsec := int64(1563180324004 * 1000 * 1000)
	nsec := int64(1553584550004 * 1000 * 1000)
	fmt.Println(time.Unix(0, nsec).String())
	//math.Pow10()
}

func TestAge(t *testing.T) {
	//s := GetTimeFromString("2019-11-30")
	//e := GetTimeFromString("2019-12-27")
	s := GetTimeFromString("1980-07-16")
	e := GetTimeFromString("2019-12-22")

	if e.Sub(s) < 0 {
		log.Panic("err")
	}
	d := e.Sub(s)
	log.Println(d.Hours()/24)

	// 计算用户实际年龄时，应以出生日期和保单生效日为计算起止点，
	// 如果保单生效日刚好在生日前一天，则认为用户还没到year (DOB)-year (保单生效日)这个数值，而是应该在这个数值的基础上-1岁
	age := e.Year() - s.Year()
	month := e.Month() - s.Month()
	day := e.Day() - s.Day()
	//if month < 0 || (month == 0 && day < 0) {
	//	age--
	//}
	log.Println(age, int(month), day)
}