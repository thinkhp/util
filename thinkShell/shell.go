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

//func ExecShell() {
// StdoutPipe方法返回一个在命令Start后与命令标准输出关联的管道。Wait方法获知命令结束后会关闭这个管道，一般不需要显式的关闭该管道。
//stdin, err := cmd.StdinPipe()
//think.Check(err)
//stdin.Write()
//stdin.Close()
//
//stdin, err = cmd.StdinPipe()
//think.Check(err)
//stdin.Write([]byte("use mysql"))
//stdin.Close()
//
//stdin, err = cmd.StdinPipe()
//think.Check(err)
//stdin.Write([]byte("select * from user"))
//stdin.Close()
//
//stdout, err := cmd.StdoutPipe()
//think.Check(err)
//str, err := ioutil.ReadAll(stdout)
//think.Check(err)
//stdout.Close()
//
//fmt.Println(str)
//}

//func execShell() {
//	cmd := exec.Command("mysql","-uroot","-p")
//	//显示运行的命令
//	fmt.Println(cmd.Args)
//	// StdoutPipe方法返回一个在命令Start后与命令标准输出关联的管道。Wait方法获知命令结束后会关闭这个管道，一般不需要显式的关闭该管道。
//	stdin, err := cmd.StdinPipe()
//	think.Check(err)
//	stdin.Write([]byte("Bank123456().pass"))
//	stdin.Close()
//}
////阻塞式的执行外部shell命令的函数,等待执行完毕并返回标准输出
//func execShell(s string) (string, error){
//	//函数返回一个*Cmd，用于使用给出的参数执行name指定的程序
//	cmd := exec.Command("/bin/bash", "-c", s)
//
//	//读取io.Writer类型的cmd.Stdout，再通过bytes.Buffer(缓冲byte类型的缓冲器)将byte类型转化为string类型(out.String():这是bytes类型提供的接口)
//	var out bytes.Buffer
//	cmd.Stdout = &out
//
//	//Run执行c包含的命令，并阻塞直到完成。  这里stdout被取出，cmd.Wait()无法正确获取stdin,stdout,stderr，则阻塞在那了
//	err := cmd.Run()
//	think.Check(err)
//
//
//	return out.String(), err
//}
//
//func execCommand(commandName string, params []string) bool {
//	//函数返回一个*Cmd，用于使用给出的参数执行name指定的程序
//	cmd := exec.Command(commandName, params...)
//
//	//显示运行的命令
//	fmt.Println(cmd.Args)
//	//StdoutPipe方法返回一个在命令Start后与命令标准输出关联的管道。Wait方法获知命令结束后会关闭这个管道，一般不需要显式的关闭该管道。
//	stdout, err := cmd.StdoutPipe()
//
//	if err != nil {
//		fmt.Println(err)
//		return false
//	}
//
//	cmd.Start()
//	//创建一个流来读取管道内内容，这里逻辑是通过一行一行的读取的
//	reader := bufio.NewReader(stdout)
//
//	//实时循环读取输出流中的一行内容
//	for {
//		line, err2 := reader.ReadString('\n')
//		if err2 != nil || io.EOF == err2 {
//			break
//		}
//		fmt.Println(line)
//	}
//
//	//阻塞直到该命令执行完成，该命令必须是被Start方法开始执行的
//	cmd.Wait()
//	return true
//}
//
////不需要执行命令的结果与成功与否，执行命令马上就返回
//func exec_shell_no_result(command string) {
//	//处理启动参数，通过空格分离 如：setsid /home/luojing/gotest/src/test_main/iwatch/test/while_little &
//	command_name_and_args := strings.FieldsFunc(command, splite_command)
//	//开始执行c包含的命令，但并不会等待该命令完成即返回
//	cmd.Start()
//	if err != nil {
//		fmt.Printf("%v: exec command:%v error:%v\n", get_time(), command, err)
//	}
//	fmt.Printf("Waiting for command:%v to finish...\n", command)
//	//阻塞等待fork出的子进程执行的结果，和cmd.Start()配合使用[不等待回收资源，会导致fork出执行shell命令的子进程变为僵尸进程]
//	err = cmd.Wait()
//	if err != nil {
//		fmt.Printf("%v: Command finished with error: %v\n", get_time(), err)
//	}
//	return
//}
