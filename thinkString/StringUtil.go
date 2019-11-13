package thinkString

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var emailPattern = regexp.MustCompile("[\\w!#$%&'*+/=?^_`{|}~-]+(?:\\.[\\w!#$%&'*+/=?^_`{|}~-]+)*@(?:[\\w](?:[\\w-]*[\\w])?\\.)+[a-zA-Z0-9](?:[\\w-]*[\\w])?")

func SizeFormat(size float64) string {
	units := []string{"Byte", "KB", "MB", "GB", "TB"}
	n := 0
	for size > 1024 {
		size /= 1024
		n += 1
	}

	return fmt.Sprintf("%.2f %s", size, units[n])
}

func IsEmail(b []byte) bool {
	return emailPattern.Match(b)
}

// 生成随机字符串
// 非并发安全,基于时间戳的随机都是不安全的
func UUID(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := make([]byte, 0)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < l; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	//r := rand.New(rand.NewSource(time.Now().UnixNano()))
	//for i := 0; i < lens; i++ {
	//	result = append(result, bytes[r.Intn(len(bytes))])
	//}
	return string(result)
}

//const hextable = "0123456789abcdef"
func RandString(l int) (string, error) {
	b := make([]byte, l)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n, err := r.Read(b)
	if n != len(b) || err != nil {
		return "", fmt.Errorf("Could not successfully read from the system CSPRNG:%w ", err)
	}
	return hex.EncodeToString(b)[:l], nil
}

func RandStringSafe(l int) (string, error){
	b := make([]byte, l)
	n, err := rand.Read(b)
	if n != len(b) || err != nil {
		return "", fmt.Errorf("Could not successfully read from the system CSPRNG:%w ", err)
	}
	return hex.EncodeToString(b)[:l], nil
}

func FirstRuneLarge(s string) string {
	s = string(s[0]+'A'-'a') + s[1:]
	return s
}

// 替换字符串的最后一位字符
func ReplaceLastRune(s *string, replace rune) {
	charBuff := []rune(*s)
	charBuff[len(charBuff)-1] = replace
	*s = string(charBuff)
}

// 本方法取自 go标准库中 rows.Scan()
func AsString(src interface{}) string {
	switch v := src.(type) {
	case string:
		return v
	case []byte:
		return string(v)
	}
	rv := reflect.ValueOf(src)
	switch rv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(rv.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(rv.Uint(), 10)
	case reflect.Float64:
		return strconv.FormatFloat(rv.Float(), 'g', -1, 64)
	case reflect.Float32:
		return strconv.FormatFloat(rv.Float(), 'g', -1, 32)
	case reflect.Bool:
		return strconv.FormatBool(rv.Bool())
	}
	return fmt.Sprintf("%v", src)
}

// 首字母小写,驼峰转换
func underlineToUpperCaseWithout(underline string) string {
	underlineArray := strings.Split(underline, "_")
	var upperStr string = ""
	for y := 0; y < len(underlineArray); y++ {
		if y == 0 {
			upperStr += string(underlineArray[0])
			continue
		}
		temp := []rune(underlineArray[y])
		for i := 0; i < len(temp); i++ {
			if i == 0 {
				temp[i] -= 32
			}
		}
		upperStr += string(temp)
	}

	return upperStr
}
