// Package StrUtil 字符串工具
package StrUtil

import (
	"github.com/whoamixl/gotools/utils/SliceUtil"
	"strings"
	"unicode"
)

// HasBlank checks if any of the given strings are blank.
// It returns true if any of the strings are blank, otherwise false.
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

// IsBlank checks if a string is blank.
// It returns true if the string is blank, otherwise false.
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

// HasEmpty checks if any of the given strings are empty.
// It returns true if any of the strings are empty, otherwise false.
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

// IsEmpty checks if a string is empty.
// It returns true if the string is empty, otherwise false.
func IsEmpty(str string) bool {
	return len(str) == 0
}

// RemoveSuffix removes the specified suffix from a string.
// If the string or suffix is empty, it returns the original string.
// If the suffix is found at the end of the string, it removes and returns the modified string.
func RemoveSuffix(str, suffix string) string {
	if !IsEmpty(str) && !IsEmpty(suffix) {
		if len(str) >= len(suffix) && str[len(str)-len(suffix):] == suffix {
			return str[:len(str)-len(suffix)]
		}
	}
	return str
}

// RemovePrefix removes the specified prefix from a string.
// If the string or prefix is empty, it returns the original string.
// If the prefix is found at the beginning of the string, it removes and returns the modified string.
func RemovePrefix(str, prefix string) string {
	if !IsEmpty(str) && !IsEmpty(prefix) {
		if len(str) >= len(prefix) && str[0:len(prefix)] == prefix {
			return str[len(prefix):]
		}
	}
	return str
}

// RemoveSuffixIgnoreCase removes the specified suffix from a string (case-insensitive).
// If the string or suffix is empty, it returns the original string.
// If the suffix is found at the end of the string (ignoring case), it removes and returns the modified string.
func RemoveSuffixIgnoreCase(str, suffix string) string {
	if !IsEmpty(str) && !IsEmpty(suffix) {
		if strings.EqualFold(suffix, str[len(str)-len(suffix):]) {
			return str[:len(str)-len(suffix)]
		}
	}
	return str
}

// RemovePrefixIgnoreCase removes the specified prefix from a string (case-insensitive).
// If the string or prefix is empty, it returns the original string.
// If the prefix is found at the beginning of the string (ignoring case), it removes and returns the modified string.
func RemovePrefixIgnoreCase(str, prefix string) string {
	if !IsEmpty(str) && !IsEmpty(prefix) {
		if strings.HasPrefix(strings.ToLower(str), strings.ToLower(prefix)) {
			return str[len(prefix):]
		}
	}
	return str
}

// Sub extracts a substring from a string based on the given indices.
// If the string is empty, it returns the original string.
// If the indices are out of range, it adjusts them accordingly.
// It returns the extracted substring.
func Sub(str string, fromIndex, toIndex int) string {
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

// Format replaces the "{}" placeholders in a string with the provided parameters.
// If the string is empty, it returns "null".
// If the parameters are empty and the string is blank, it returns the original string.
// It returns the formatted string.
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
