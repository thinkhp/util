package think

import (
	"util/thinkLog"
)

func DeferRecoverCommon() {
	if r := recover(); r != nil {
		thinkLog.ErrorLog.Println("[recover] 程序已恢复:", r)
	}
}