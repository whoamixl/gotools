package main

import (
	"fmt"
	"time"
)

func main() {
	//fmt.Println(time.Now())
	//fmt.Println(time.Now().Format("15:04:05"))
	//fmt.Println(time.Now().Format("20060102150405"))
	//fmt.Println(time.Now().Format("2006-01-02"))
	//fmt.Println(time.Now().Format("2006/01/02"))
	//fmt.Println(time.Now().Format("2006-01-02 15:04:05.000"))
	//fmt.Println(time.Now().Format("2006-01-02 15:04:05.000000"))
	//fmt.Println(time.Now().Format("2006-01-02 15:04:05.0000000"))
	//fmt.Println(time.Now().Format("2006-01-02 15:04:05.000000000"))
	//fmt.Println(DateUtil.Now().Format("yyyy/MM/dd HH:mm:ss"))
	//fmt.Println(DateUtil.Now().Format("yyyyMMdd HH:mm:ss"))
	//fmt.Println(DateUtil.Now().Format("yyyyMMdd HHmmss"))
	//fmt.Println(DateUtil.Now().Format("yyyy"))
	//fmt.Println(DateUtil.Now().Format("yyyy-MM-dd HH:mm:ss"))
	//fmt.Println(DateUtil.Now().Format("yyyy-MM-dd HH:mm:ss.SSS"))
	//fmt.Println(DateUtil.Now().Format("yyyy-MM-dd HH:mm:ss.SSSSSS"))
	//fmt.Println(DateUtil.Now().Format("yyyy-MM-dd HH:mm:ss.SSSSSSS"))
	//fmt.Println(DateUtil.Now().Format("yyyy-MM-dd HH:mm:ss.SSSSSSSSS"))
	//fmt.Println(DateUtil.FormatByTime(time.Now(), "yyyy-MM-dd HH:mm:ss.SSSSSSSSS"))
	//fmt.Println(time.Parse("2006-01-02 15:04:05", "2023-07-30 22:26:06"))
	//d, _ := DateUtil.Parse("yyyy-MM-dd HH:mm:ss", "2023-05-10 22:26:06")
	//fmt.Println(d.Day())
	now := time.Now()
	then := time.Date(2021, 10, 1, 0, 0, 0, 0, time.UTC)
	days := now.Sub(then).Hours() / 24
	fmt.Println("两个时间之间有", days, "天")

}
