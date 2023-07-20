package main

import (
	"fmt"
	"github.com/whoamixl/gotools/utils/HexUtil"
)

func main() {
	str := HexUtil.EncodeHexStr("hello,world")
	fmt.Println(str)
	str1, err := HexUtil.DecodeHexStr(str)
	if err == nil {
		fmt.Println(str1)
	}
}
