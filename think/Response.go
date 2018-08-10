package think

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"runtime/debug"
	"util/thinkLog"
)

type Response struct {
	ReturnCode int         `json:"returnCode"`
	ReturnMsg  string      `json:"returnMsg"`
	Data       interface{} `json:"data"`
}

// 默认 defer
func DeferRecover(w http.ResponseWriter) {
	if r := recover(); r != nil {
		thinkLog.ErrorLog.Println("[recover] 程序已恢复")
		GetResponseJsonFail(w, 500, "server error", 500)
	}
}

func DeferCommon() {
	if r := recover(); r != nil {
		thinkLog.ErrorLog.Println(r)
		thinkLog.ErrorLog.Println(string(debug.Stack()))
	}
}

func DeferTransaction(tx *sql.Tx) {
	if r := recover(); r != nil {
		tx.Rollback()
		thinkLog.ErrorLog.Println(r)
		thinkLog.ErrorLog.Println(string(debug.Stack()))
	} else {
		tx.Commit()
	}
}
func GetResponseJsonFail(w http.ResponseWriter, code int, msg string, status int) {
	getResponseJson(w, code, msg, nil, status)
}

func GetResponseJsonOK(w http.ResponseWriter, obj interface{}) {
	getResponseJson(w, 200, "ok", obj)
}

// status 服务器或客户端 error ,返回状态码:200,403...,500...
// suc 对于输出内容的判断,true:输出成功;bool:错误信息
// obj 返回主体
// suc, code, msg, data, status
func getResponseJson(w http.ResponseWriter, code int, msg string, obj interface{}, status ...int) {
	// 状态码
	if len(status) != 0 {
		w.WriteHeader(status[0])
	}
	// 返回主体
	w.Header().Set("Content-Type", "application/json")
	var response Response
	response = Response{code, msg, obj}
	// json化,私有属性无法json化
	str, err := json.Marshal(response)
	Check(err)
	w.Write(str)
}

func Check(err error) {
	if err != nil {
		//thinkLog.ErrorLog.Println(string(debug.Stack()))
		thinkLog.ErrorLog.Panic(err)
	}
}
