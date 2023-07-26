package main

import (
	"fmt"
	"github.com/whoamixl/gotools/utils/IdcardUtil"
)

func main() {
	//a := NumberUtil.Add(0.336, 0.0053)
	//fmt.Println(a)
	//b := NumberUtil.Sub(0.336, 0.0053)
	//fmt.Println(b)
	//c := NumberUtil.Mul(0.5, 0.3)
	//fmt.Println(c)
	//d := NumberUtil.Div(0.6, 0.2)
	//fmt.Println(d)
	//a1 := NumberUtil.RoundFloat(4.225233, 2)
	//fmt.Println(a1)
	//a2 := NumberUtil.TruncateFloat(4.225233, 2)
	//fmt.Println(a2)
	//b1, _ := NumberUtil.RoundFloatString("4.33252", 3)
	//fmt.Println(b1)
	//b2, _ := NumberUtil.TruncateFloatString("4.33252", 3)
	//fmt.Println(b2)
	//fmt.Println("////////////////////")
	//
	//c1 := NumberUtil.IsNumber("223.33")
	//c2 := NumberUtil.IsNumber("223")
	//c3 := NumberUtil.IsNumber("sss")
	//fmt.Println(c1, c2, c3)
	//fmt.Println(NumberUtil.GenerateRandomNumber(1, 59))
	//
	//slice1 := []int{1, 2, 3}
	//slice2 := []int{4, 5, 6}
	//slice3 := []int{7, 8, 9}
	//result := SliceUtil.AddAll(slice1, slice2, slice3)
	//fmt.Println(result)
	//
	//fmt.Println(SliceUtil.Range(1, 22, 2))
	//fmt.Println(SliceUtil.Split([]int{2, 3, 6, 7, 34, 34, 6, 7, 3}, 3))
	//
	//keys := []string{"a", "b", "c"}
	//values := []int{1, 2, 3}
	//result2, _ := SliceUtil.Zip(keys, values)
	//fmt.Println(result2)

	//type S struct {
	//	A string
	//	B int
	//}
	//cc1 := []string{"2", "33", "5"}
	//cc2 := []int{2, 4, 5}
	//cc3 := []S{S{A: "11", B: 2}, S{A: "33", B: 3}}
	//fmt.Println(SliceUtil.Contains(cc1, "2"))
	//fmt.Println(SliceUtil.Join(cc2, "---"))
	//fmt.Println(SliceUtil.ToString(cc3))
	/////
	//
	//fmt.Println(RandomUtil.RandomInt(2, 99))
	//
	//err, ds := RandomUtil.RandomBytes(22)
	//fmt.Println(err, ds)
	//dds := NumberUtil.GenerateRandomNumber(22, 33)
	//fmt.Println(dds)

	//a := []string{"22", "33", "34", "35"}
	//fmt.Println(RandomUtil.RandomEle(a))
	//fmt.Println(RandomUtil.RandomUniqueEleSet(a, 2))
	//fmt.Println(RandomUtil.RandomString(6))
	//
	//fmt.Println(IdUtil.UUID())
	//fmt.Println(IdUtil.NewSnowflake(1))

	//regex := `(\d{4})-(\d{2})-(\d{2})`
	//content := "My birthday is 1990-01-01 and my friend's birthday is 1995-12-31."
	//groupIndex := 0
	//
	//pattern := regexp.MustCompile(regex)
	//matches := pattern.FindAllStringSubmatch(content, -1)
	//fmt.Println(matches)
	//matchedStrings := ReUtil.FindAll(regex, content, groupIndex)
	//matchedString := ReUtil.Find(regex, content, groupIndex)
	//if len(matchedStrings) > 0 {
	//	println("Matched strings:", matchedStrings)
	//	for _, i := range matchedStrings {
	//		fmt.Println(i)
	//	}
	//} else {
	//	println("No matches found.")
	//}
	//fmt.Println(matchedString)
	card, b := IdCardUtil.NewIdCard("513721199809102292")
	if !b {
		fmt.Println(card.BirthDay)
		fmt.Println(card.Province)
	}

}
