package string

import (
	"log"
	"strings"
	"testing"
)

/*
@Time : 2021/10/7 15:19
@Author : onns
@File : string/test_trim_test.go
*/

func Test(t *testing.T) {
	a := "\t \b\r\n 这是一段中文\t\r\n "
	out("a", a, "")
	b := strings.Trim(a, " ")
	out("b", b, " ")
	c := strings.Trim(a, "\t ")
	out("c", c, "\t ")
}

func out(n, s, r string) {
	log.Printf("string %s is %s,[]byte is %v, length %d, trim bytes %v", n, s, []byte(s), len(s), []byte(r))
}
