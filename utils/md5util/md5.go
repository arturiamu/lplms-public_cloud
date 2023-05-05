package md5util

import (
	"crypto/md5"
	"fmt"
)

func Byte2MD5(bytes []byte) []byte {
	hash := md5.New()
	_, err := hash.Write(bytes)
	if err != nil {
		panic(err)
	}
	sum := hash.Sum(nil)
	return sum
}

func Byte2MD5Str(bytes []byte) string {
	return fmt.Sprintf("%x", Byte2MD5(bytes))
}

// String2MD5Str get 32byte md5util
func String2MD5Str(str string) string {
	return Byte2MD5Str([]byte(str))
}
