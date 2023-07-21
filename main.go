package main

import (
	"fmt"
	"github.com/whoamixl/gotools/utils/NumberUtil"
)

func main() {
	a := NumberUtil.Add(0.336, 0.0053)
	fmt.Println(a)
	b := NumberUtil.Sub(0.336, 0.0053)
	fmt.Println(b)
	c := NumberUtil.Mul(0.5, 0.3)
	fmt.Println(c)
	d := NumberUtil.Div(0.6, 0.2)
	fmt.Println(d)
	a1 := NumberUtil.RoundFloat(4.225233, 2)
	fmt.Println(a1)
	a2 := NumberUtil.TruncateFloat(4.225233, 2)
	fmt.Println(a2)
	b1, _ := NumberUtil.RoundFloatString("4.33252", 3)
	fmt.Println(b1)
	b2, _ := NumberUtil.TruncateFloatString("4.33252", 3)
	fmt.Println(b2)
	fmt.Println("////////////////////")
}
