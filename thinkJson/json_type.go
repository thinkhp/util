package thinkJson

//要将json数据解码写入一个接口类型值，函数会将数据解码为如下类型写入接口：
//Bool                   对应JSON布尔类型
//float64                对应JSON数字类型
//string                 对应JSON字符串类型
//[]interface{}          对应JSON数组
//map[string]interface{} 对应JSON对象
//nil                    对应JSON的null
type JsonBool bool
type JsonFloat float64
type JsonString string
type JsonList []interface{}
type JsonMap map[string]interface{}
type JsonInterface interface{}
