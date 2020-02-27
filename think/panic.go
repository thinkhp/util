package think

import (
	"runtime/debug"
	"util/thinkLog"
)

//func Check(err error) {
//	if err != nil {
//		thinkLog.ErrorLog.Println(err)
//		thinkLog.ErrorLog.Println(string(debug.Stack()))
//		panic(err)
//	}
//}
func Panic(err error) {
	if err != nil {
		panic(err)
	}
}

func IsNil(err error) {
	if err != nil {
		thinkLog.ErrorLog.Println(err)
		thinkLog.ErrorLog.Println(string(debug.Stack()))
		panic(err)
	}
}
