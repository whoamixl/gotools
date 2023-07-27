package main

import (
	"fmt"
	"github.com/whoamixl/gotools/utils/DesensitizedUtil"
)

func main() {
	fmt.Println(DesensitizedUtil.UserId("12344"))
	fmt.Println("qweerrt"[:])
	fmt.Println(DesensitizedUtil.ChineseName("闫小龙"))
	fmt.Println(DesensitizedUtil.IdCard("5137211998091", 1, 2))
	fmt.Println(DesensitizedUtil.FixedPhone("09157518479x"))
	fmt.Println(DesensitizedUtil.MobilePhone("18899999923"))
}
