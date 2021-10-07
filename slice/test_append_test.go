package slice

import (
	"fmt"
	"testing"
)

/*
@Time : 2021/10/7 15:08
@Author : onns
@File : slice/test_append_test.go
*/

func Test(t *testing.T) {
	i := 0
	temp := []int{0,1}
	test(append(temp[:i], temp[i+1:]...))
	fmt.Println(temp)
}
