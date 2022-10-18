package service

import (
	"checkCDN/client/utils"
	"fmt"
)

func Ping(domain string)(string)  {
	result := utils.PingExec(domain)
	fmt.Print("result",result)
	return result
}