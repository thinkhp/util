package thinkTimer

import (
	"time"
)

func GetNextWeekday(now time.Time, weekday time.Weekday) time.Time {
	var next time.Time
	// 判断是否为星期日
	if now.Weekday() == weekday {
		// 判断是否为星期日的00:00:00
		if now.Hour() == 0 && now.Minute() == 0 && now.Second() == 0 {
			return now
		} else {
			next = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
			next = next.Add(time.Hour * 24 * 7)
		}
	} else {
		for now.Weekday() != weekday {
			now = now.Add(time.Hour * 24)
		}
		next = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	}

	return next
}

func GetNextMinute(now time.Time) time.Time {
	if now.Second() == 0 {
		return now
	} else {
		for now.Second() != 0 {
			now = now.Add(time.Second)
			//fmt.Println(now)
		}
	}

	return now
}

func GetNextSecond(now time.Time) time.Time {
	if now.Nanosecond() == 0 {
		return now
	} else {
		now = time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second()+1, 0, now.Location())
	}

	return now
}

func ThinkTimer(nextTime func(time.Time) time.Time, f func(time.Duration)) {
	go func() {
		for {
			next := nextTime(time.Now())
			d := next.Sub(time.Now())
			f(d)
			// time.Sleep(d)
			t := time.NewTimer(d)
			<-t.C
		}
	}()
}
