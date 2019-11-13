package thinkHttp

import (
	"io/ioutil"
	"net/http"
	"util/think"
	"util/thinkLog"
)

func GetRequestBody(r *http.Request) []byte {
	body, err := ioutil.ReadAll(r.Body)
	think.IsNil(err)
	defer r.Body.Close()
	thinkLog.DebugLog.Println(SprintRequestReceive(r.Method, r.URL.String(), r.Header, body))

	return body
}
