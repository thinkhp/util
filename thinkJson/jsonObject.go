package thinkJson

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"util/think"
)

type ErrNotGetValue struct {
	o   interface{}
	key string
}

func (e ErrNotGetValue) Error() string {
	return fmt.Sprintf("json:error not get '%s' from '%s'", reflect.TypeOf(e.o).String(), e.key)
}

type JsonObject map[string]interface{}

//要将json数据解码写入一个接口类型值，函数会将数据解码为如下类型写入接口：
//Bool                   对应JSON布尔类型
//float64                对应JSON数字类型
//string                 对应JSON字符串类型
//[]interface{}          对应JSON数组
//map[string]interface{} 对应JSON对象
//nil                    对应JSON的null
func GetJsonObject(data []byte) (JsonObject, error) {
	jsonObject := make(map[string]interface{})
	// !!!json.Unmarshal第二个参数为指针,jsonObject
	err := json.Unmarshal(data, &jsonObject)
	return jsonObject, err
}

func MustGetJsonObject(data []byte) JsonObject {
	jsonObject := make(map[string]interface{})
	// !!!json.Unmarshal第二个参数为指针,jsonObject
	err := json.Unmarshal(data, &jsonObject)
	if err != nil {
		panic(err)
	}
	return jsonObject
}

func MustGetList(data []byte) []interface{}{
	list := make([]interface{}, 0)
	// !!!json.Unmarshal第二个参数为指针,jsonObject
	err := json.Unmarshal(data, &list)
	if err != nil {
		panic(err)
	}
	return list
}

func GetList(data []byte) ([]interface{}, error) {
	list := make([]interface{}, 0)
	// !!!json.Unmarshal第二个参数为指针,jsonObject
	err := json.Unmarshal(data, &list)
	return list, err
}


// Deprecated:
// 因为要尽量减少工具包之间的关联性,所以要将该功能拆分
func GetJsonObjectFromRequest(r *http.Request) JsonObject {
	body, err := ioutil.ReadAll(r.Body)
	think.IsNil(err)
	defer r.Body.Close()

	return MustGetJsonObject(body)
}

func (jsonObject JsonObject) TransMapStringString() map[string]string {
	m := make(map[string]string)
	for k, v := range jsonObject {
		m[k] = v.(string)
	}

	return m
}

// 覆盖
func (jsonObject JsonObject) SetObject(k string, v interface{}) JsonObject {
	jsonObject[k] = v
	return jsonObject
}

// 不覆盖
func (jsonObject JsonObject) SetObjectNotCover(k string, v interface{}) JsonObject {
	_, ok := jsonObject[k]
	if ok {
		return jsonObject
	} else {
		jsonObject[k] = v
		return jsonObject
	}
}
