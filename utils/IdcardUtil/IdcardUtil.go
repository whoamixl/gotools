// Package IdCardUtil 身份证解析
package IdCardUtil

import (
	"strconv"
	"strings"
	"time"
)

type idCard struct {
	Id       string
	BirthDay string
	Age      int64
	Year     int64
	Month    int64
	Day      int64
	Gender   bool
	Province string
}

func NewIdCard(id string) (*idCard, bool) {
	idCard := &idCard{Id: id}
	if !idCard.verifyIdCard() {
		return idCard, false
	}
	idCard.birthDay()
	idCard.age()
	idCard.year()
	idCard.month()
	idCard.day()
	idCard.gender()
	idCard.province()
	return idCard, true
}

func (id *idCard) birthDay() {
	birthday := id.Id[6:14]
	t, err := time.Parse("20060102", birthday)
	if err != nil {
		id.BirthDay = ""
	}
	id.BirthDay = t.Format("2006-01-02")
}

func (id *idCard) age() {
	birthday := id.BirthDay
	currentTime := time.Now()
	birthTime, _ := time.Parse("2006-01-02", birthday)
	age := int64(currentTime.Year() - birthTime.Year())
	if currentTime.YearDay() < birthTime.YearDay() {
		age--
	}
	id.Age = age
}
func (id *idCard) year() {
	birthday := id.BirthDay
	year, _ := strconv.Atoi(birthday[0:4])
	id.Year = int64(year)
}
func (id *idCard) month() {
	birthday := id.BirthDay
	month, _ := strconv.Atoi(birthday[5:7])
	id.Month = int64(month)
}
func (id *idCard) day() {
	birthday := id.BirthDay
	day, _ := strconv.Atoi(birthday[8:10])
	id.Day = int64(day)
}
func (id *idCard) gender() {
	genderCode := id.Id[16:17]
	if genderCode == "1" || genderCode == "3" || genderCode == "5" || genderCode == "7" || genderCode == "9" {
		id.Gender = true // 男性
	} else {
		id.Gender = false // 女性
	}
}
func (id *idCard) province() {
	provinceCode := id.Id[0:2]
	provinceMap := map[string]string{
		"11": "北京市",
		"12": "天津市",
		"13": "河北省",
		"14": "山西省",
		"15": "内蒙古自治区",
		"21": "辽宁省",
		"22": "吉林省",
		"23": "黑龙江省",
		"31": "上海市",
		"32": "江苏省",
		"33": "浙江省",
		"34": "安徽省",
		"35": "福建省",
		"36": "江西省",
		"37": "山东省",
		"41": "河南省",
		"42": "湖北省",
		"43": "湖南省",
		"44": "广东省",
		"45": "广西壮族自治区",
		"46": "海南省",
		"50": "重庆市",
		"51": "四川省",
		"52": "贵州省",
		"53": "云南省",
		"54": "西藏自治区",
		"61": "陕西省",
		"62": "甘肃省",
		"63": "青海省",
		"64": "宁夏回族自治区",
		"65": "新疆维吾尔自治区",
		"71": "台湾省",
		"81": "香港特别行政区",
		"82": "澳门特别行政区",
		"91": "国外",
	}
	province, ok := provinceMap[provinceCode]
	if ok {
		id.Province = province
	} else {
		id.Province = "未知"
	}
}

func (id *idCard) verifyIdCard() bool {
	// 去除空格并转换为大写
	id.Id = strings.ToUpper(strings.ReplaceAll(id.Id, " ", ""))
	// 检查身份证号码长度
	if len(id.Id) != 18 {
		return false
	}
	// 检查身份证号码格式
	if !strings.ContainsAny(id.Id[:17], "0123456789") || (id.Id[17] != 'X' && !strings.ContainsAny(id.Id[17:], "0123456789")) {
		return false
	}
	// 计算校验位
	factor := []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	checkSum := 0
	for i := 0; i < 17; i++ {
		num, err := strconv.Atoi(string(id.Id[i]))
		if err != nil {
			return false
		}
		checkSum += num * factor[i]
	}
	checkDigit := "10X98765432"[checkSum%11]
	// 检查校验位
	if id.Id[17] != checkDigit {
		return false
	}
	return true
}
