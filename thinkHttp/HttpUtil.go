package thinkHttp

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"util/think"
	"util/thinkJson"
)

func Send(method, url string, paramsMap map[string]interface{}, headersMap map[string]string) (thinkJson.JsonObject, error) {
	var request *http.Request
	var err error
	// 设置method,url,body
	request, err = http.NewRequest(method, url, makeParams(paramsMap))
	think.Check(err)
	// 设置header
	for k, v := range headersMap {
		request.Header.Set(k, v)
	}
	// 发送
	response, err := http.DefaultClient.Do(request)
	think.Check(err)
	// 解析回应
	if response.StatusCode != 200 {
		return nil, errors.New(response.Status)
	}
	body, err := ioutil.ReadAll(response.Body)
	think.Check(err)
	jsonObject := thinkJson.GetJsonObject(body)

	return jsonObject, nil
}

func makeParams(paramsMap map[string]interface{}) io.Reader {
	paramsJson, err := json.Marshal(paramsMap)
	think.Check(err)
	return bytes.NewReader([]byte(paramsJson))
}
