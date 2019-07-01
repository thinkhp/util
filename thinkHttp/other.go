package thinkHttp

import (
	"fmt"
	"net/http"
	"runtime/debug"
	"util/thinkLog"
)

// 默认 deferHttp
func DeferRecoverHttp(w http.ResponseWriter) {
	if r := recover(); r != nil {
		thinkLog.ErrorLog.Println(r)
		thinkLog.ErrorLog.Println(string(debug.Stack()))
		thinkLog.ErrorLog.Println("[recover] 程序已恢复")
		WriteJsonFail(w, 50000, fmt.Sprint(r))
	}
}

// 重定向
func ToRedirect(w http.ResponseWriter, newPath string) {
	thinkLog.DebugLog.Println("redirect to", newPath)
	w.Header().Set("Location", newPath)
	w.WriteHeader(301)
}


