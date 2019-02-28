package thinkHttp

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"util/think"
	"util/thinkLog"
)

func GetRequestBody(r *http.Request) []byte {
	body, err := ioutil.ReadAll(r.Body)
	think.IsNil(err)
	defer r.Body.Close()

	fmt.Println()
	for i := 0; i < len(body); i++ {
		fmt.Print(body[i], ",")
	}
	fmt.Println()
	thinkLog.DebugLog.PrintParams(r.URL.String(), "receive, body", string(body))

	return body
}
