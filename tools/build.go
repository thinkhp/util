package tools

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

type Project struct {
	Name         string //项目名称,同时也是输出的可执行文件名称
	Path         string //项目所在的绝对路径
	MainFileName string //main 文件的名称
}

var ErrGOOS = errors.New("不支持的操作系统")
const GOOSWin = "windows"
const GOOSMac = "darwin"
const GOOSLinux = "linux"

func (p Project) Build(aimGOOS string) (output []byte, err error) {
	exeName := p.Name
	mainFile := p.Path + p.MainFileName
	binDir := p.Path
	suffix := ""
	cmds := make([]string, 0, 5)
	switch aimGOOS {
	case GOOSWin: //如果要编译运行在 win 下的可执行文件,加后缀 .exe
		exeName += ".exe"
	}
	// 生成脚本(.sh 或者 .bat 文件)
	switch runtime.GOOS {
	case GOOSMac:
		suffix = ".sh"
		cmds = append(cmds, "#!/usr/bin/env bash\n")
		cmds = append(cmds, "cd "+binDir+"\n")
		cmds = append(cmds, fmt.Sprintf("GOOS=%s go build -o %s %s\n", aimGOOS, exeName, mainFile))
		cmds = append(cmds, "exit 0")
	case GOOSWin:
		suffix = ".bat"
		cmds = append(cmds, binDir[:2]+"\n") //盘符
		cmds = append(cmds, "cd "+binDir+"\n")
		cmds = append(cmds, "set GOOS="+aimGOOS+"\n")
		cmds = append(cmds, fmt.Sprintf("go build -o %s %s", exeName, mainFile)+"\n")
		cmds = append(cmds, "set GOOS=windows\n")
	case GOOSLinux:
		suffix = ".sh"
		cmds = append(cmds, "#!/bin/bash\n")
		cmds = append(cmds, "cd "+binDir+"\n")
		cmds = append(cmds, fmt.Sprintf("GOOS=%s go build -o %s %s", aimGOOS, exeName, mainFile)+"\n")
		cmds = append(cmds, "exit 0")
	default:
		return nil, ErrGOOS
	}

	file, err := os.OpenFile("../build"+suffix, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0770)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	fileName, _ := filepath.Abs(file.Name())
	fmt.Println("脚本文件路径:",fileName)


	fmt.Println("脚本内容:")
	for _, cmd := range cmds {
		_, err := file.WriteString(cmd)
		if err != nil {
			return nil, err
		}
		fmt.Print(cmd)
	}
	file.Close()


	// 执行脚本
	cmd := exec.Command(fileName)
	//fmt.Println("cmd:",cmd.Args)
	errput := make([]byte, 0)
	cmd.Stdout = bytes.NewBuffer(output)
	cmd.Stderr = bytes.NewBuffer(errput)
	if err := cmd.Run(); err != nil {
		return nil, err
	}

	if len(errput) != 0 {
		return output, errors.New(string(errput))
	}
	//fmt.Println(string(output))
	return output, nil
}
