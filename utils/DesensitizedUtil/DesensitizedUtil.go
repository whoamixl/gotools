package DesensitizedUtil

import (
	"errors"
	"github.com/whoamixl/gotools/utils/NumberUtil"
	"github.com/whoamixl/gotools/utils/SliceUtil"
)

// UserId function takes a string parameter and returns a hidden version of the string.
func UserId(str string) string {
	data, _ := hide(str, 1, 0)
	return data
}

// ChineseName function takes a string parameter and returns a hidden version of the string, based on the length of the string.
func ChineseName(str string) string {
	if len(str) == 2 {
		data, _ := hide(str, 0, 1)
		return data
	}
	if len(str) > 2 {
		data, _ := hide(str, 1, 1)
		return data
	}
	return str
}

// IdCard function takes a string parameter and two integer parameters, start and end, and returns a hidden version of the string based on the specified start and end positions.
func IdCard(str string, start, end int) (string, error) {
	if len(str) == 18 {
		data, err := hide(str, start, end)
		return data, err
	}
	return str, errors.New("The card number is not legal")
}

func FixedPhone(str string) (string, error) {
	length := len(str)
	lengthSlice := []int{10, 11, 12}
	if SliceUtil.Contains(lengthSlice, length) && NumberUtil.IsNumber(str) {
		data, _ := hide(str, 4, 2)
		return data, nil
	}
	return str, errors.New("The fixed phone number is legal")

}

func MobilePhone(str string) (string, error) {
	if len(str) == 11 && NumberUtil.IsNumber(str) || str[:1] == "1" {
		data, _ := hide(str, 3, 4)
		return data, nil
	}
	return str, errors.New("The fixed phone number is legal")
}

// hide function is a generic function for hiding strings. It takes a string parameter and two integer parameters, start and end, and returns a hidden version of the string based on the specified start and end positions.
func hide(str string, start, end int) (string, error) {
	if (start + end) >= len(str) {
		return str, errors.New("The length of the start and end cannot be greater than the length of str")
	}
	iData := 0
	if start == 0 {
		iData = len(str) - end
		return generateAsterisk(iData) + str[len(str)-end:], nil
	} else if end == 0 {
		iData = len(str) - start
		return str[:start] + generateAsterisk(iData), nil
	}
	return str[:start] + generateAsterisk(len(str)-start-end) + str[len(str)-end:], nil
}

// generateAsterisk function takes an integer parameter and returns a string of asterisks based on the specified number.
func generateAsterisk(num int) string {
	strS := ""
	for i := 0; i < num; i++ {
		strS = strS + "*"
	}
	return strS
}
