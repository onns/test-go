package string

import (
	"log"
	"testing"
)

/*
@Time : 2021/10/6 22:32
@Author : onns
@File : string/test_string_test.go
*/

func Test_getCurrentIp(t *testing.T) {
	r := getCurrentIp()
	rr := String(r)
	log.Print(len(r))
	log.Print([]byte(r))
	log.Print([]byte("101.87.56.106"))
	log.Print(len("101.87.56.106"))
	log.Print(string([]byte{10}))
	log.Print([]byte("\n"))

	log.Print(len(StringValue(rr)))
}
