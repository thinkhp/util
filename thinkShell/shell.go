package thinkShell

import (
	"bytes"
	"os"
	"os/exec"
	"util/think"
	"util/thinkFile"
	"util/thinkLog"
)

func RunCommand(name string, args []string) []byte {
	cmd := exec.Command(name, args...)
	thinkLog.DebugLog.Println("执行的COMMAND:", cmd.Args)

	var output bytes.Buffer
	var error bytes.Buffer
	cmd.Stdout = &output
	cmd.Stderr = &error
	cmd.Run()

	if len(error.Bytes()) != 0 {
		thinkLog.ErrorLog.Panic(error.String())
	}
	return output.Bytes()
}

func RunSH(filePath, startFileName string, commands []string) []byte {
	startFileName = "start.sh"
	file := thinkFile.OpenFile(filePath, startFileName, os.O_CREATE|os.O_WRONLY)
	var buffer bytes.Buffer
	for i := 0; i < len(commands); i++ {
		buffer.WriteString(commands[i])
	}
	file.Write(buffer.Bytes())
	defer file.Close()

	output := RunCommand(filePath+startFileName, nil)
	thinkLog.DebugLog.Println(filePath+startFileName, "running")

	return output
}

func RunBash(commands []string, stopBash chan bool) []byte {
	cmd := exec.Command("/bin/bash")
	var input bytes.Buffer
	for i := 0; i < len(commands); i++ {
		command := commands[i]
		thinkLog.DebugLog.Println("command", i, command)
		input.WriteString(command)
	}
	cmd.Stdin = bytes.NewReader(input.Bytes())

	var output bytes.Buffer
	var errOutput bytes.Buffer
	cmd.Stdout = &output
	cmd.Stderr = &errOutput
	// Bash 后台运行程序时,会检测一个动作,主动释放Bash
	if stopBash != nil {
		cmd.Start()
		//fmt.Println("have a nohup")
		flag, ok := <-stopBash
		if !flag || !ok {
			thinkLog.ErrorLog.Panic(errOutput.String())
		}
	} else {
		cmd.Run()
		cmd.Wait()
		if len(errOutput.Bytes()) != 0 {
			thinkLog.ErrorLog.Println(errOutput.String())
		}
		/// 无法校验cmd的真正运行是否错误
		//if len(output.Bytes()) == 0 && len(errOutput.Bytes()) != 0{
		//	thinkLog.WarnLog.Println(output.String())
		//	// string(debug.Stack())
		//}
	}
	cmdPid := cmd.Process.Pid
	cmd.Process.Kill()
	thinkLog.DebugLog.Println("[bash] <defunct> PID:", cmdPid)

	return output.Bytes()
}

// 思路:写一个 bat 文件,然后执行该文件
// 暂时只能在goProject/bin目录下运行命令
// file必须在运行外的函数执行,类似于db
func createBat(commands []string, filePath, fileName string) string {
	// 生成文件
	file, err := os.OpenFile(filePath+fileName, os.O_WRONLY|os.O_CREATE, 0766)
	think.Check(err)
	var buffer bytes.Buffer
	for i := 0; i < len(commands); i++ {
		buffer.WriteString(commands[i])
	}
	file.Write(buffer.Bytes())
	defer file.Close()
	return fileName
}
func RunBat(commands []string, batPath, batName string) []byte {
	createBat(commands, batPath, batName)
	// 执行bat
	//path, err := exec.LookPath(batName)
	//think.Check(err)
	//thinkLog.DebugLog.Println("fortune is available at ", batPath)
	output := RunCommand(batPath+batName, nil)

	return output
}