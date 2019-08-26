package thinkString

import (
	"bytes"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/url"
	"testing"
)

func TestFirstRuneLarge(t *testing.T) {
	fmt.Println(FirstRuneLarge("hello"))
}

func TestUrlE(t *testing.T) {
	fmt.Println(url.QueryEscape("https://https://btl188.com/fengqiu/code/get"))
}

func TestUrlUn(t *testing.T) {
	fmt.Println(url.QueryUnescape("https%3A%2F%2Fbtl188.com%2Ffengqiu%2Fcode%2Fget"))
}

func TestUUID(t *testing.T) {
	fmt.Println(UUID(32))
}

func TestGB2312(t *testing.T) {
	utf8Str := "hello"
	utf8Byte := []byte(utf8Str)

	gbByte, err := UTF82GB2312(utf8Byte)
	if err != nil {
		panic(err)
	}
	gbStr := string(gbByte)

	fmt.Println(utf8Byte)
	fmt.Println(gbByte)
	fmt.Println(gbStr)

}

func TestGetGb2312(t *testing.T) {
	e := simplifiedchinese.HZGB2312.NewEncoder()
	gb, _ := e.Bytes([]byte("// 此处为字符串string,结束"))
	fmt.Println(gb)
}

func UTF82GB2312(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

// 此处为字符串string,结束 [47 47 32 230 173 164 229 164 132 228 184 186 229 173 151 231 172 166 228 184 178 115 116 114 105 110 103 44 231 187 147 230 157 159]
// �˴�Ϊ�ַ���string,���� [47 47 32 180 203 180 166 206 170 215 214 183 251 180 174 115 116 114 105 110 103 44 189 225 202 248]
// ~{4K4&N*WV7{4.~}string,~{=aJx [47 47 32 126 123 52 75 52 38 78 42 87 86 55 123 52 46 126 125 115 116 114 105 110 103 44 126 123 61 97 74 120]
func TestFile(t *testing.T) {
	var check = func(err error) {
		if err != nil {
			panic(err)
		}
	}
	utf8Byte, err := ioutil.ReadFile("./utf8")
	check(err)
	gb2312Byte, err := ioutil.ReadFile("./gb2312")
	check(err)

	gb2312Byte1, err := UTF82GB2312(utf8Byte)
	check(err)

	fmt.Println(string(utf8Byte), utf8Byte)
	fmt.Println(string(gb2312Byte), gb2312Byte)
	fmt.Println(string(gb2312Byte1), gb2312Byte1)
}
