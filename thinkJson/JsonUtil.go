package thinkJson

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
	"time"
	"util/think"
	"util/timeUtil"
)

type JsonObject map[string]interface{}

//要将json数据解码写入一个接口类型值，函数会将数据解码为如下类型写入接口：
//Bool                   对应JSON布尔类型
//float64                对应JSON数字类型
//string                 对应JSON字符串类型
//[]interface{}          对应JSON数组
//map[string]interface{} 对应JSON对象
//nil                    对应JSON的null
func GetJsonObject(data []byte) JsonObject {
	jsonObject := make(map[string]interface{})
	// !!!json.Unmarshal第二个参数为指针,jsonObject
	json.Unmarshal(data, &jsonObject)
	return jsonObject
}

// 以下函数含有参数校验:若为nil,panic()
func (jsonObject JsonObject) GetObject(key string) JsonObject {
	jsonObjectUnder, ok := jsonObject[key].(map[string]interface{})
	if ok {
		return jsonObjectUnder
	} else {
		return nil
	}

}

func (jsonObject JsonObject) GetBool(key string) bool {
	return jsonObject[key].(bool)
}

func (jsonObject JsonObject) GetString(key string) string {
	str, ok := jsonObject[key].(string)
	if ok {
		return str
	} else {
		return ""
	}
}

func (jsonObject JsonObject) GetInt(key string) int {
	return int(jsonObject[key].(float64))
}

func (jsonObject JsonObject) GetFloat64(key string) float64 {
	return jsonObject[key].(float64)
}

func (jsonObject JsonObject) GetTime(key string) time.Time {
	datetime := jsonObject[key].(string)
	return timeUtil.GetTimeFromString(datetime)
}

func (jsonObject JsonObject) GetList(key string) []JsonObject {
	list := jsonObject[key].([]interface{})
	var jsonObjectSlice = make([]JsonObject, 0, len(list))
	for i := 0; i < len(list); i++ {
		jsonObjectSlice = append(jsonObjectSlice, list[i].(map[string]interface{}))
	}
	return jsonObjectSlice
}

func (jsonObject JsonObject) GetStruct(ptr interface{}) {
	v := reflect.ValueOf(ptr).Elem() // the struct variable
	// 获取结构体字段
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i) // a reflect.StructField
		name := fieldInfo.Name
		// json,tag
		tag := fieldInfo.Tag
		jsonName := tag.Get("json")
		if jsonName == "" {
			jsonName = name
		}
		//
		field := v.FieldByName(name)
		switch field.Kind() {
		case reflect.String:
			field.SetString(jsonObject.GetString(jsonName))
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			field.SetInt(int64(jsonObject.GetInt(jsonName)))
		case reflect.Float64, reflect.Float32:
			field.SetFloat(jsonObject.GetFloat64(jsonName))
		case reflect.Bool:
			field.SetBool(jsonObject.GetBool(jsonName))
		default:
			field.SetPointer(nil)
		}
	}
}

func GetJsonObjectFromRequest(r *http.Request) JsonObject {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	think.Check(err)

	return GetJsonObject(body)
}
