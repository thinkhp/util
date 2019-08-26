package tools

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"util/thinkFile"
)

type fileName struct {
	dir  string
	file string
}

// 生成 fmt.sh 文件
// rootPath 整理根目录
func fmtSH(rootPath string, excludes []string) string {
	rootPath, err := filepath.Abs(rootPath)
	if err != nil {
		panic(err)
	}
	fmt.Println("待整理的根目录", rootPath)
	toFmtFiles := make(map[string][]string)

	// 找到该目录下所有后缀为 .go 的文件
	allFile := thinkFile.ListFile(rootPath, ".go")
T:
	for _, f := range allFile {
		index := strings.LastIndex(f, "/")
		dir := f[:index] //文件目录
		// 不整理包名包含 excludes 下的包
		for _, exclude := range excludes {
			if strings.Contains(dir, exclude) {
				continue T
			}
		}
		name := f[index+1:]                             //文件名称
		toFmtFiles[dir] = append(toFmtFiles[dir], name) //添加到待整理的list
	}

	s := ""
	for k, v := range toFmtFiles {
		s += "cd " + k + "\n"
		for i := range v {
			s += "go fmt " + v[i] + "\n"
		}
	}

	f, err := os.OpenFile("./fmt.sh", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0770)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.Write([]byte(s))

	return f.Name()
}
