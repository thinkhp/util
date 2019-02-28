package thinkHttp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"util/think"
	"util/thinkJson"
	"util/thinkLog"
)

func ExampleSendGet() {
	url := "http://www.baidu.com"
	resBody, err := SendGet(url, nil)
	think.IsNil(err)
	res := thinkJson.MustGetJsonObject(resBody)
	haveAccessToken := res["access_token"]
	fmt.Println(haveAccessToken)
}

func ExampleSendPost() {
	url := "https://www.baidu.com"
	paramsMap := map[string]interface{}{
		"type":   "news",
		"offset": 0,
		"count":  20,
	}

	resBody, err := SendPost(url, nil,  thinkJson.MustMarshal(paramsMap))
	think.IsNil(err)
	res := thinkJson.MustGetJsonObject(resBody)
	haveAccessToken := res["access_token"]
	fmt.Println(haveAccessToken)
}

func ExampleSend() {
	client := &http.Client{}
	urlHead := "https://wxyx.pandains.com/kj"
	// url
	var url = urlHead + "/cus_bal_dtl/detail?source=1005&sourceId=2"
	// params
	var paramMap = map[string]interface{}{
		"principal": "1",
	}
	// header
	var headerMap = map[string]string{
		"Content-Type": "application/thinkJson",
		"superman":     "superman",
	}
	// *取值，&取址
	params := thinkJson.MustMarshal(paramMap)
	thinkLog.DebugLog.PrintParams(url, "request params", string(params))
	request, err := http.NewRequest("POST", url, bytes.NewReader(params))
	think.IsNil(err)
	// 设置header
	for k, v := range headerMap {
		request.Header.Set(k, v)
	}
	//request.Body.Read(paramByte)
	//request.Body.Close()
	// 发送请求
	response, _ := client.Do(request)
	// 解析回应
	var jsonMap = make(map[string]interface{})
	if response != nil && response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		err := json.Unmarshal(body, &jsonMap)
		if err != nil {
			thinkLog.ErrorLog.Println(err)
			//return 0, nil
		}
	}
	if len(jsonMap) == 0 {
		//return 0, nil
	}
	//return 1, jsonMap
}
