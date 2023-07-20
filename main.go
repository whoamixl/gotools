package main

import (
	"fmt"
	"github.com/whoamixl/gotools/utils/NumberUtil"
)

func main() {
	fmt.Println(0.336 + 0.0053)
	d := NumberUtil.Add(0.336, 0.0053)
	fmt.Println(d)

}
