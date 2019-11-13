package thinkHttp

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"util/think"
	"util/thinkLog"
)

func SendGetWithClient(client *http.Client, url string, headersMap map[string]string) ([]byte, error) {
	return Send(client, http.MethodGet, url, headersMap, nil)
}
func SendPostWithClient(client *http.Client, url string, headersMap map[string]string, params []byte) ([]byte, error) {
	return Send(client, http.MethodPost, url, headersMap, params)
}
func SendGet(url string, headersMap map[string]string) ([]byte, error) {
	return Send(http.DefaultClient, http.MethodGet, url, headersMap, nil)
}
func SendPost(url string, headersMap map[string]string, params []byte) ([]byte, error) {
	return Send(http.DefaultClient, http.MethodPost, url, headersMap, params)
}

//e.g.
// POST Send(http.MethodPost, "http://www.baidu.com", nil, []byte("{'key'='value'}"))
// Get	Send(http.MethodGet, "https://www.baidu.com/s?wd=keyword", nil, nil)
func Send(client *http.Client, method string, url string, headersMap map[string]string, params []byte) ([]byte, error) {
	var request *http.Request
	var err error

	thinkLog.DebugLog.Println(SprintRequest(method, url, headersMap, params))
	// 设置method,url,body
	request, err = http.NewRequest(method, url, bytes.NewReader(params))
	think.IsNil(err)
	// 设置header
	for k, v := range headersMap {
		request.Header.Set(k, v)
	}
	//log.Println("resquest", request)
	// 发送
	//fmt.Println("timeout:", client.Timeout)
	response, err := client.Do(request)
	think.IsNil(err)
	body, err := ioutil.ReadAll(response.Body)
	think.IsNil(err)
	defer response.Body.Close()

	thinkLog.DebugLog.Println(SprintResponse(response.StatusCode, url, response.Header, body))
	//log.Println("response", response)
	// 解析回应
	if response.StatusCode != 200 {
		return body, errors.New(response.Status)
	}

	return body, nil
}