package thinkHttp

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"util/think"
	"util/thinkJson"
	"util/thinkLog"
)

type Response struct {
	// 业务状态码
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func WriteStatus(w http.ResponseWriter, statusCode int){
	w.WriteHeader(statusCode)
}

// 返回HTML
func WriteHTMLPage(w http.ResponseWriter, path string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	data, err := ioutil.ReadFile(path)
	think.IsNil(err)
	w.Write(data)
}

// 返回HTML
func WriteHTML(w http.ResponseWriter, body []byte) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(body)
}

// 返回JSON
func WriteJsonOk(w http.ResponseWriter, data interface{}) {
	writeJson(w, 20000, "ok", data)
}
func WriteJsonFail(w http.ResponseWriter, code int, msg string) {
	writeJson(w, code, msg, nil)
}

func writeJson(w http.ResponseWriter, code int, msg string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	response := Response{code, msg, data}
	// json化,私有属性无法json化
	bytes, err := json.Marshal(response)
	think.IsNil(err)
	thinkLog.DebugLog.PrintParams("", "response, body", string(bytes))
	w.Write(bytes)
}

func WriteJson(w http.ResponseWriter, data interface{})  {
	w.Header().Set("Content-Type", "application/json")
	bytes := thinkJson.MustMarshal(data)
	thinkLog.DebugLog.PrintParams("", "response, body", string(bytes))
	w.Write(bytes)
}
