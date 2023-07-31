package main

import (
	"fmt"
	"github.com/whoamixl/gotools/net/HttpUtil"
)

func main() {
	c := HttpUtil.NewClient()
	m := map[string]string{"ip": "47.100.247.211", "port": "22"}
	r, _ := c.Url("http://test.inis.cn/api/other/ping").Method(HttpUtil.GET).Param(m).Send()
	fmt.Println(r.StatusCode, r.Body)
}
