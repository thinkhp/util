package thinkHttp

import (
	"fmt"
	"net/http"
	"strings"
)

type ThinkRequest struct {
	Url      string
	Function func(http.ResponseWriter, *http.Request)
	Role     int
}

var RequestMap = make(map[string]*ThinkRequest)

// 注册路由
func AddHttpFunc(url string, function func(http.ResponseWriter, *http.Request), role ...int) {
	if len(role) == 0 {
		RequestMap[url] = &ThinkRequest{url, function, 0}
	} else {
		RequestMap[url] = &ThinkRequest{url, function, role[0]}
	}
}

func FindHttpFunc(url string) (func(http.ResponseWriter, *http.Request), bool) {
	// url携带?和=
	if strings.Contains(url, "?") && strings.Contains(url, "=") {
		i := strings.Index(url, "?")
		url = url[:i]
	}
	fmt.Println("匹配URL", url)
	f, ok := RequestMap[url]
	if ok {
		return f.Function, true
	}
	return nil, false
}

// 注册静态文件服务器
func AddFileServer(url string) {

}

//权限控制(函数权限,999所需权限最小,0需要超级管理员)
//if requestMap[url].role < 0{
//	think.GetResponseJson(w,false,"not page",404)
//	return
//}

//
//import (
//	"reflect"
//	"fmt"
//	"strings"
//)
//
//type User struct {
//	User
//	role int
//}
//
//type ControllerFunction struct {
//	url string
//	function *func()
//	role int
//}
//
//var RequestMap map[string]ControllerFunction
//var UserMap map[string]User
//
//
//func get(){
//	var typeName reflect.Type
//	typeName =
//		reflect.Zero()
//}
//
//type Controllers struct {
//	ControllersType []reflect.Type
//	BaseController []BaseController
//}
//
//// 智能路由
//// url 权限 func
//type BaseController struct {
//	ControllerType reflect.Type
//	url string
//	ControllerFunction *[]ControllerFunction
//}
//
//
//
//// 注册路由
//func (c *Controllers)addAuto(obj interface{}){
//	// 通过添加一个空对象来注册controller
//	c.ControllersType = append(c.ControllersType, reflect.TypeOf(obj))
//	// 将该controller 的所有方法进行自动路由化
//	// BaseController
//	baseController := new(BaseController)
//	// BaseController 的 type
//	controllerType := reflect.TypeOf(obj)
//	// BaseController 的 url
//	packageName := "generate."
//	controllerSuffix := "Controller"
//	url := strings.TrimLeft(controllerType.String(),packageName)
//	url = strings.TrimRight(url,controllerSuffix)
//	// BaseController 的
//	controllerFunctionSlice := make([]ControllerFunction,0)
//	for i := 0;i < controllerType.NumMethod();i++ {
//		controllerFunction := new(ControllerFunction)
//		controllerFunction.role = 999
//		controllerFunction.url = controllerType.Method(i).Name.
//			methodType := controllerType.Method(i)
//	}
//	controllerType.
//
//
//
//}
//
//func (b *BaseController)initRequestMap(){
//
//	for v := range b.Controllers {
//		typeValue := reflect.ValueOf(v)
//		typeName := typeValue.Type()
//
//		// 获取该
//	}
//	// 获取所有的BaseController
//	var baseController BaseController
//	var typeName reflect.Type
//	typeName = reflect.TypeOf(baseController)
//	fmt.Println("************************type")
//	fmt.Println(typeName)
//	fmt.Printf("%T\n",baseController)
//	fmt.Println("************************value")
//	var typeValue reflect.Value
//	typeValue = reflect.ValueOf(baseController)
//	fmt.Println(typeValue)
//	//RequestMap[] = BaseController{}
//}
//
