package main

import (
	"fmt"
	"github.com/whoamixl/gotools/utils/NumberUtil"
	"github.com/whoamixl/gotools/utils/SliceUtil"
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

	c1 := NumberUtil.IsNumber("223.33")
	c2 := NumberUtil.IsNumber("223")
	c3 := NumberUtil.IsNumber("sss")
	fmt.Println(c1, c2, c3)
	fmt.Println(NumberUtil.GenerateRandomNumber(1, 59))

	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6}
	slice3 := []int{7, 8, 9}
	result := SliceUtil.AddAll(slice1, slice2, slice3)
	fmt.Println(result)

	fmt.Println(SliceUtil.Range(1, 22, 2))
	fmt.Println(SliceUtil.Split([]int{2, 3, 6, 7, 34, 34, 6, 7, 3}, 3))

	keys := []string{"a", "b", "c"}
	values := []int{1, 2, 3}
	result2, _ := SliceUtil.Zip(keys, values)
	fmt.Println(result2)

}
