package thinkFile

import (
	"os"
	"path/filepath"
	"strings"
	"util/think"
		)

//权限说明：
//O_RDONLY int = syscall.O_RDONLY // 只读
//O_WRONLY int = syscall.O_WRONLY // 只写
//O_RDWR int = syscall.O_RDWR // 读写
//O_APPEND int = syscall.O_APPEND // 在文件末尾追加，打开后cursor在文件结尾位置
//O_CREATE int = syscall.O_CREAT // 如果不存在则创建
//O_EXCL int = syscall.O_EXCL //与O_CREATE一起用，构成一个新建文件的功能，它要求文件必须不存在
//O_SYNC int = syscall.O_SYNC // 同步方式打开，没有缓存，这样写入内容直接写入硬盘，系统掉电文件内容有一定保证
//O_TRUNC int = syscall.O_TRUNC // 打开并清空文件

// The defined file mode bits are the most significant bits of the FileMode.
// The nine least-significant bits are the standard Unix rwxrwxrwx permissions.
// The values of these bits should be considered part of the public API and
// may be used in wire protocols or disk representations: they must not be
// changed, although new bits might be added.
//const (
//	// The single letters are the abbreviations
//	// used by the String method's formatting.
//	ModeDir        FileMode = 1 << (32 - 1 - iota) // d: is a directory
//	ModeAppend                                     // a: append-only
//	ModeExclusive                                  // l: exclusive use
//	ModeTemporary                                  // T: temporary file; Plan 9 only
//	ModeSymlink                                    // L: symbolic link
//	ModeDevice                                     // D: device file
//	ModeNamedPipe                                  // p: named pipe (FIFO)
//	ModeSocket                                     // S: Unix domain socket
//	ModeSetuid                                     // u: setuid
//	ModeSetgid                                     // g: setgid
//	ModeCharDevice                                 // c: Unix character device, when ModeDevice is set
//	ModeSticky                                     // t: sticky
//
//	// Mask for the type bits. For regular files, none will be set.
//	ModeType = ModeDir | ModeSymlink | ModeNamedPipe | ModeSocket | ModeDevice
//
//	ModePerm FileMode = 0777 // Unix permission bits
//)
func OpenFile(filePath string, fileName string, flag int) *os.File {
	CreatePath(filePath)
	fileFullName := filePath + fileName
	// os.O_WRONLY|os.O_CREATE|os.O_APPEND
	// 以只写方式打开文件
	// 如果不存在，则创建
	// 在文件末尾追加
	file, err := os.OpenFile(fileFullName, flag, 0766)
	think.Check(err)
	// 本次打开的文件如果close,log日志无法写入
	// file.Close()
	return file
}

func CreatePath(path string) {
	// 如果不存在,则创建目录
	_, err := os.Stat(path)
	if err != nil {
		if !os.IsNotExist(err) {
			think.Check(err)
		} else {
			// 创建文件夹
			err := os.MkdirAll(path, os.ModePerm)
			think.Check(err)
		}
	}
}

// basePath: ././XXX/YYY
// return: /XXX/YYY/ZZZ/AAA/
func GetAbsPathWith(basePath string) string {
	basePath = strings.Replace(basePath, "/", string(os.PathSeparator), -1)
	path, err := filepath.Abs(basePath)
	think.Check(err)

	if strings.HasSuffix(path, string(os.PathSeparator)) {
		return path
	} else {
		return path + string(os.PathSeparator)
	}
}

// 遍历当前目录下的文件
func ListFile(path string,suffix string) []string{
	allFile := make([]string, 0)
	filePath := GetAbsPathWith(path)
	// 遍历filePath下的所有文件以及目录,ls .sql 文件
	filepath.Walk(filePath, func(path string, info os.FileInfo, err error) error {
		if info == nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(path, suffix) {
			allFile = append(allFile, path)
			return nil
		} else {
			return nil
		}
		return nil
	})

	return allFile
}
