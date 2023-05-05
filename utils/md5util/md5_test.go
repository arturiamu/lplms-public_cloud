package md5util

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	str := "123"
	fmt.Println(String2MD5Str(str))
}

func TestBT(t *testing.T) {
	str := "123"
	bt := []byte(str)
	str2 := string(bt)
	fmt.Println(str)
	fmt.Println(bt)
	fmt.Println(str2)
}
