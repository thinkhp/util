package thinkString

import (
	"fmt"
	"net/url"
	"testing"
)

func TestFirstRuneLarge(t *testing.T) {
	fmt.Println(FirstRuneLarge("hello"))
}

func TestUrlE(t *testing.T)  {
	fmt.Println(url.QueryEscape("https://https://btl188.com/fengqiu/code/get"))
}

func TestUrlUn(t *testing.T)  {
	fmt.Println(url.QueryUnescape("https%3A%2F%2Fbtl188.com%2Ffengqiu%2Fcode%2Fget"))
}