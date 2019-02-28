package thinkJson

import (
	"reflect"
	"time"
	"util/timeUtil"
)

// 以下函数含有参数校验:若为nil,panic()
func (jsonObject JsonObject) MustGetJsonObject(key string) (o JsonObject) {
	o, ok := jsonObject[key].(map[string]interface{})
	if !ok {
		panic(ErrNotGetValue{o, key})
	}
	return o
}

func (jsonObject JsonObject) MustGetBool(key string) (flag bool) {
	flag, ok := jsonObject[key].(bool)
	if !ok {
		panic(ErrNotGetValue{flag, key})
	}
	return flag
}

func (jsonObject JsonObject) MustGetString(key string) (str string) {
	str, ok := jsonObject[key].(string)
	if !ok {
		panic(ErrNotGetValue{str, key})
	}
	return str
}

func (jsonObject JsonObject) MustGetInt(key string) (i int) {
	f, ok := jsonObject[key].(float64)
	if !ok {
		panic(ErrNotGetValue{i, key})
	}
	return int(f)
}

func (jsonObject JsonObject) MustGetFloat64(key string) (f float64) {
	f, ok := jsonObject[key].(float64)
	if !ok {
		panic(ErrNotGetValue{f, key})
	}
	return f
}

func (jsonObject JsonObject) MustGetTime(key string) (t time.Time) {
	s, ok := jsonObject[key].(string)
	if !ok {
		panic(ErrNotGetValue{s, key})
	}
	return timeUtil.GetTimeFromString(s)
}

func (jsonObject JsonObject) MustGetList(key string) (l []JsonObject) {
	list, ok := jsonObject[key].([]interface{})
	if !ok {
		panic(ErrNotGetValue{l, key})
	}
	l = make([]JsonObject, 0, len(list))
	for i := 0; i < len(list); i++ {
		l = append(l, list[i].(map[string]interface{}))
	}
	return l
}

func (jsonObject JsonObject) MustGetStringList(key string) (l []string) {
	list, ok := jsonObject[key].([]interface{})
	if !ok {
		panic(ErrNotGetValue{l, key})
	}
	l = make([]string, 0, len(list))
	for i := 0; i < len(list); i++ {
		l = append(l, list[i].(string))
	}
	return l
}

func (jsonObject JsonObject) MustGetStruct(ptr interface{}) {
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
			str := jsonObject.MustGetString(jsonName)
			field.SetString(str)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			num := jsonObject.MustGetInt(jsonName)
			field.SetInt(int64(num))
		case reflect.Float64, reflect.Float32:
			num := jsonObject.MustGetFloat64(jsonName)
			field.SetFloat(num)
		case reflect.Bool:
			flag := jsonObject.MustGetBool(jsonName)
			field.SetBool(flag)
		default:
			field.SetPointer(nil)
		}
	}
}
