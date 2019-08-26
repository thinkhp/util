package tools

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestFmt(t *testing.T) {
	var err error
	fileName := fmtSH("../", []string{"github", "gopkg.in", "gui"})
	fmt.Println(os.Getwd())
	fileName, err = filepath.Abs(fileName)
	if err != nil {
		panic(err)
	}
	fmt.Println("fmt.sh:", fileName)

	cmd := exec.Command("sh", fileName)
	//cmd.Path, err = os.Getwd()
	//fmt.Println("工作目录", cmd.Path)
	if err != nil {
		panic(err)
	}
	err = cmd.Run()
	if err != nil {
		panic(err)
	}
}
