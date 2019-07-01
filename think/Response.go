package think

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	ReturnCode int         `json:"returnCode"`
	ReturnMsg  string      `json:"returnMsg"`
	Data       interface{} `json:"data"`
}

// Deprecated:
func GetResponseJsonFail(w http.ResponseWriter, code int, msg string, status int) {
	getResponseJson(w, code, msg, nil, status)
}

// Deprecated:
func GetResponseJsonOK(w http.ResponseWriter, obj interface{}) {
	getResponseJson(w, 200, "ok", obj)
}

// status 服务器或客户端 error ,返回状态码:200,403...,500...
// suc 对于输出内容的判断,true:输出成功;bool:错误信息
// obj 返回主体
// suc, code, msg, data, status
// Deprecated:
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
	IsNil(err)
	w.Write(str)
}
