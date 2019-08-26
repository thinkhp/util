package timeUtil

import (
	"fmt"
	"testing"
	"time"
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
