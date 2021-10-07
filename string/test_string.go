package string

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
@Time : 2021/10/6 22:32
@Author : onns
@File : string/test_string.go
*/


func getCurrentIp() (ip string) {
	responseClient, err := http.Get("http://members.3322.org/dyndns/getip") // 获取外网 IP
	if err != nil {
		fmt.Printf("获取外网 IP 失败，请检查网络\n")
		panic(err)
	}
	defer responseClient.Body.Close()
	body, _ := ioutil.ReadAll(responseClient.Body)
	ip = fmt.Sprintf("%s", string(body))
	// ip = strings.Trim(ip,"\n")
	return
}

func String(a string) *string {
	return &a
}

func StringValue(a *string) string {
	if a == nil {
		return ""
	}
	return *a
}