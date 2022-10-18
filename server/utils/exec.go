package utils

import (
	"bytes"
	"fmt"
	"github.com/mahonia"
	"os/exec"
	"runtime"
)

func convertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}

func PingExec(domain string)(result string)  {
	sysType := runtime.GOOS
	var command = &exec.Cmd{}
	if sysType == "windows"{
		command = exec.Command("ping","-n","1", domain)
	}else if sysType == "linux"{
		command = exec.Command("ping","-c","1", domain)
	}

	outinfo := bytes.Buffer{}
	command.Stdout = &outinfo
	err := command.Start()
	if err != nil{
		fmt.Println(err.Error())
	}
	if err = command.Wait();err!=nil{
		fmt.Println(err.Error())
		return err.Error()
	}else{
		//fmt.Println(command.ProcessState.Pid())
		//fmt.Println(command.ProcessState.Sys().(syscall.WaitStatus).ExitCode)
		fmt.Println(convertToString(outinfo.String(),"gbk", "utf-8"))
		return convertToString(outinfo.String(),"gbk", "utf-8")
	}
}
