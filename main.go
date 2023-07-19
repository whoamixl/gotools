package main

import (
	"fmt"
	"github.com/whoamixl/gotools/utils/StrUtil"
)

func main() {
	s := "我是{}，今年{}岁"
	str := StrUtil.Format(s, "yxl", "18")
	fmt.Println(str)
}
