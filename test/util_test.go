package test

import (
	"crypto/aes"
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/Go-SQL-Driver/MySQL"
	"os"
	"reflect"
	"runtime"
	"testing"
	"time"
	"util/database"
	"util/think"
	"util/thinkCrypto"
	"util/thinkMath"
)

var num = 123456.1415926
var s = 1
var t = time.Now()

type TestStruct struct {
	TestId    int             `json:"testId"`
	Num       sql.NullInt64   `json:"num"`
	Str       sql.RawBytes    `json:"str"`
	Sum       sql.NullInt64   `json:"sum"`
	Timestamp mysql.NullTime  `json:"timestamp"`
	Version   sql.NullInt64   `json:"version"`
	Datetime  mysql.NullTime  `json:"datetime"`
	Date      mysql.NullTime  `json:"date"`
	FloatNum  sql.NullFloat64 `json:"floatNum"`
}

func TestUtil(test *testing.T) {
	fmt.Println(errors.New("hello").Error())
	fmt.Println(time.Now().Unix())
	var i int64 = 63665625600
	fmt.Println(time.Unix(i, 0).String())
}

func TestKeepDecimal(t *testing.T) {
	num := []float64{2.754999, 3.2236, 1.3686, 3.098, 4.897, 5.91, 1.34, 0.001, 9.999999, 9.090909, 5.45}
	keep := []int{0, 1, 2, 3, 4, 5}
	for i := 0; i < len(num); i++ {
		for j := 0; j < len(keep); j++ {
			fmt.Println(num[i], "\t\t", keep[j], "\t\t", thinkMath.KeepDecimal(num[i], keep[j]), "\t\t", thinkMath.Round(num[i], keep[j]))
		}
	}
}

func TestRound(t *testing.T) {
	num := 9.9999999999
	for i := 0; i < 9; i++ {
		num -= 1.1111111111
		num = thinkMath.Round(num, 10)
		fmt.Println(num)
	}

}

func four() {
	strOne := "16.94"
	strTwo := "16.21"
	fmt.Println(strOne < strTwo)
}

func checkClose() {
	db, err := sql.Open("mysql", "root:Bank123456().pass@tcp(localhost:3306)/test?parseTime=true&loc=Local")
	think.IsNil(err)
	defer db.Close()

	var datetime mysql.NullTime
	rows, err := db.Query("select datetime from test")
	for rows.Next() {
		rows.Scan(&datetime)
	}
	fmt.Println(datetime)

	file, err := os.OpenFile("./test.txt", os.O_WRONLY|os.O_CREATE, 0766)
	defer file.Close()
	file.WriteString("hello")
}

func getColumnsType() {
	db := database.SetConn("root:Bank123456().pass@tcp(localhost:3306)/test?parseTime=true&loc=Local")
	defer db.Close()

	var datetime mysql.NullTime
	rows := database.Select(nil, "select datetime from test")
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&datetime)
	}
	fmt.Println(datetime)
	//ScanType(DatabaseType)
	//int32(非空数字)
	//sql.NullInt64(可能为空的数字)
	//sql.RawBytes(字符串)
	//sql.NullInt64
	//mysql.NullTime(timestamp)
	//sql.NullInt64
	//mysql.NullTime(datetime)
	//mysql.NullTime(date)
	//sql.NullFloat64
	database.GetColumnsType("select * from test")
}

func getFuncName(f func(time.Time) string) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}

func one(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func two(t time.Time) string {
	return t.String()[0:19]
}

// 编码
// Base64编码在RFC2045中定义，它被定义为：
// Base64内容传送编码被设计用来把任意序列的8位字节描述为一种不易被人直接识别的形式
func encode(bytes []byte) string {
	return base64.StdEncoding.EncodeToString(bytes)
}

// 解码
func decode(str string) []byte {
	dataByte, err := base64.StdEncoding.DecodeString(str)
	think.IsNil(err)
	return dataByte
}
func print() {
	fmt.Println(time.Now())
}

func decrypt() {
	// 最初明文
	keyT := "example key 1234"
	cipherT := "some data for test,please check 1:2:3:4"
	// 用于发送的编码明文
	keyForTran := encode([]byte(keyT))
	cipherForTran := encode([]byte(cipherT))
	//
	keyByte := decode(keyForTran)
	cipherByte := decode(cipherForTran)
	ivByte := cipherByte[:aes.BlockSize]
	// 加密
	result := encode(thinkCrypto.AESEncrypt(cipherByte, keyByte, ivByte))
	// 发送 key,result,iv
	// 解密
	fmt.Println(string(decode(encode(thinkCrypto.AESDecrypt(result, keyForTran, encode(ivByte))))))
}
