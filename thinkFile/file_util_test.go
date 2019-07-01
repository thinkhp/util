package thinkFile

import (
	"fmt"
	"os"
	"strconv"
	"testing"
)

func TestFilePerm(t *testing.T) {
	fmt.Println(os.ModeDir, os.ModePerm)
}

func TestFileUtil(t *testing.T) {
	//打开文件，返回文件指针
	file, error := os.Open("./1.txt")
	if error != nil {
		fmt.Println(error)
	}
	//创建byte的slice用于接收文件读取数据
	buf := make([]byte, 1024)
	//循环读取
	for {
		//Read函数会改变文件当前偏移量
		len, _ := file.Read(buf)

		//读取字节数为0时跳出循环
		if len == 0 {
			break
		}

		fmt.Println(string(buf))
	}
	fmt.Println(file)
	file.Close()

	//以读写方式打开文件，如果不存在，则创建
	file2, error := os.OpenFile("./2.txt", os.O_RDWR|os.O_CREATE, 0766)
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println(file2)
	file2.Close()
	//读取文件内容
	file5, error := os.Open("./1.txt")
	if error != nil {
		fmt.Println(error)
	}
	buf2 := make([]byte, 1024)
	ix := 0
	for {
		//ReadAt从指定的偏移量开始读取，不会改变文件偏移量
		len, _ := file5.ReadAt(buf2, int64(ix))
		ix = ix + len
		if len == 0 {
			break
		}

		fmt.Println(string(buf2))
	}
	file5.Close()

	//写入文件
	file6, error := os.Create("./4.txt")
	if error != nil {
		fmt.Println(error)
	}
	data := "我是数据\r\n"
	for i := 0; i < 10; i++ {
		//写入byte的slice数据
		file6.Write([]byte(data))
		//写入字符串
		file6.WriteString(data)
	}
	file6.Close()

	//写入文件
	file7, error := os.Create("./5.txt")
	if error != nil {
		fmt.Println(error)
	}
	for i := 0; i < 10; i++ {
		//按指定偏移量写入数据
		ix := i * 64
		file7.WriteAt([]byte("我是数据"+strconv.Itoa(i)+"\r\n"), int64(ix))
	}
	file7.Close()

	//删除文件
	del := os.Remove("./1.txt")
	if del != nil {
		fmt.Println(del)
	}

	//删除指定path下的所有文件
	delDir := os.RemoveAll("./dir")
	if delDir != nil {
		fmt.Println(delDir)
	}
}
