package thinkHttp

import (
	"io/ioutil"
	"net/http"
	"util/think"
	"util/thinkLog"
)

// TODO requestID
func GetRequestBody(r *http.Request) []byte {
	body, err := ioutil.ReadAll(r.Body)
	think.IsNil(err)
	defer r.Body.Close()
	l := new(logger)
	thinkLog.DebugLog.Println(l.SprintRequestReceive(r.Method, r.URL.String(), r.Header, body))

	return body
}

func GetRequestBodyWithLogger(r *http.Request) ([]byte,*logger) {
	body, err := ioutil.ReadAll(r.Body)
	think.IsNil(err)
	defer r.Body.Close()
	l := new(logger).Init()
	thinkLog.DebugLog.Println(l.SprintRequestReceive(r.Method, r.URL.String(), r.Header, body))

	return body, l
}
