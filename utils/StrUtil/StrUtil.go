package StrUtil

import (
	"github.com/whoamixl/gotools/utils/SliceUtil"
	"strings"
	"unicode"
)

func HasBlank(str ...string) bool {
	if SliceUtil.IsEmpty(str) {
		return true
	}
	length := len(str)
	for i := 0; i < length; i++ {
		str := str[i]
		if IsBlank(str) {
			return true
		}
	}
	return false
}

func IsBlank(str string) bool {
	length := len(str)

	for i := 0; i < length; i++ {
		char := rune(str[i])
		if !unicode.IsSpace(char) {
			return false
		}
	}
	return true
}

func HasEmpty(str ...string) bool {
	if SliceUtil.IsEmpty(str) {
		return true
	}
	length := len(str)
	for i := 0; i < length; i++ {
		str := str[i]
		if IsEmpty(str) {
			return true
		}
	}
	return false

}

func IsEmpty(str string) bool {
	return len(str) == 0
}

func RemoveSuffix(str string, suffix string) string {
	if !IsEmpty(str) && !IsEmpty(suffix) {
		if len(str) >= len(suffix) && str[len(str)-len(suffix):] == suffix {
			return str[:len(str)-len(suffix)]
		}
	}
	return str
}

func RemovePrefix(str string, prefix string) string {
	if !IsEmpty(str) && !IsEmpty(prefix) {
		if len(str) >= len(prefix) && str[0:len(prefix)] == prefix {
			return str[len(prefix):]
		}
	}
	return str
}

func RemoveSuffixIgnoreCase(str, suffix string) string {
	if !IsEmpty(str) && !IsEmpty(suffix) {
		if strings.EqualFold(suffix, str[len(str)-len(suffix):]) {
			return str[:len(str)-len(suffix)]
		}
	}
	return str
}

func RemovePrefixIgnoreCase(str, prefix string) string {
	if !IsEmpty(str) && !IsEmpty(prefix) {
		if strings.HasPrefix(strings.ToLower(str), strings.ToLower(prefix)) {
			return str[len(prefix):]
		}
	}
	return str
}

func Sub(str string, fromIndex int, toIndex int) string {
	if IsEmpty(str) {
		return str
	}
	len := len(str)
	if fromIndex < 0 {
		fromIndex += len
		if fromIndex < 0 {
			fromIndex = 0
		}
	} else if fromIndex > len {
		fromIndex = len
	}

	if toIndex < 0 {
		toIndex += len
		if toIndex < 0 {
			toIndex = len
		}
	} else if toIndex > len {
		toIndex = len
	}

	if toIndex < fromIndex {
		tmp := fromIndex
		fromIndex = toIndex
		toIndex = tmp
	}
	if fromIndex == toIndex {
		return ""
	} else {
		return str[fromIndex:toIndex]
	}
}

func Format(str string, params ...string) string {
	if len(str) == 0 {
		return "null"
	}
	if SliceUtil.IsEmpty(params) && IsBlank(str) {
		return str
	}
	strSlice := strings.Split(str, "{}")
	strS := ""
	for i, strI := range strSlice {
		if i == len(strSlice)-1 {
			return strS
		}
		strS = strS + strI + params[i]
	}
	return strS
}
