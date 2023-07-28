package main

import (
	"fmt"
	"github.com/whoamixl/gotools/utils/CreditCodeUtil"
	"github.com/whoamixl/gotools/utils/DesensitizedUtil"
)

func main() {

	fmt.Println(DesensitizedUtil.Email("18899999923@qq.com"))
	fmt.Println(DesensitizedUtil.Password("18899999923@qq.com"))
	fmt.Println(CreditCodeUtil.CheckSocietyId("GJ23054659193LH21T"))
	fmt.Println(CreditCodeUtil.GenerateRandomSocialCreditCode())
}
