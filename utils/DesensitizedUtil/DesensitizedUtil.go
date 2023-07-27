package DesensitizedUtil

import (
	"errors"
)

func UserId(str string) string {
	data, _ := hide(str, 1, 0)
	return data
}

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

func IdCard(str string, start, end int) (string, error) {
	if len(str) != 18 {
		return str, errors.New("The card number is not legal")
	}
	data, err := hide(str, start, end)
	return data, err
}

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

func generateAsterisk(num int) string {
	strS := ""
	for i := 0; i < num; i++ {
		strS = strS + "*"
	}
	return strS
}
