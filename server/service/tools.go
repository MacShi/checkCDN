package service

import (
	"checkCDN/server/modle"
	"encoding/json"
	"fmt"
	"github.com/pelletier/go-toml"
	"io/ioutil"
	"net/http"
	"strings"

	//"net/http"
	"checkCDN/server/config"
)

func Ping(domain string)(aa *modle.CdnResult)  {
	defer func() {
		err:=recover()
		if err!=nil{
			fmt.Println(err)
		}
	}()
	cdnResult:=modle.CdnResult{}
	ips :=[]string{}
	lossRateList :=[]string{}
	//result := utils.PingExec(domain)
	urls := config.Conf.Get("client.infos").([]*toml.Tree)
	for _,value :=range urls{
		clientAddr := value.Get("host").(string)
		clientPort := value.Get("port").(string)
		clientToken := value.Get("token").(string)
		ipAddr,lossRate,err:=reqPing(clientAddr,clientPort,clientToken,domain)
		if err!=nil{
			fmt.Println(err.Error())
			break
		}
		if !isContain(ips,ipAddr){
			ips = append(ips, ipAddr)
		}
		lossRateList = append(lossRateList,lossRate)
	}
	if len(ips)>1{
		cdnResult.HaveCDN="Yes"
		cdnResult.IpAddr =ips
		cdnResult.LossRate = lossRateList
	}else {
		cdnResult.HaveCDN="No"
		cdnResult.IpAddr =ips
		cdnResult.LossRate = lossRateList
	}
	return &cdnResult
}
//
func reqPing(clientAddr string,clientPort string,clientToken string,host string) (ip string,loss string,err error) {
	url := fmt.Sprintf("http://%s:%s/apis/ping?token=%s&host=%s",clientAddr,clientPort,clientToken,host)
	req , err:=http.Get(url)
	if err!=nil{
		fmt.Println(err.Error())
		return "nil","nil",err
	}
	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		fmt.Println(err.Error())
		return "nil","nil",err
	}
	var respMap  map[string]interface{}
	str2mapErr := json.Unmarshal([]byte(body),&respMap)
	if str2mapErr!=nil{
		fmt.Println(str2mapErr.Error())
		return "nil","nil",err
	}
	str1 :=strings.Split(respMap["data"].(string),"bytes of data")[0]
	str2 := strings.Split(str1," (")[1]
	ipAddr := strings.Split(str2,") ")[0]
	str3 := strings.Split(respMap["data"].(string)," packet loss")[0]
	lossRate := strings.Split(str3,"received, ")[1]
	fmt.Println(ipAddr,lossRate)
	return ipAddr,lossRate,nil
}

func isContain(items []string, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}
