package ReUtil

import (
	"regexp"
	"strings"
)

// Get returns the matched string based on the given pattern and content
// with the specified group index.
func Get(pattern *regexp.Regexp, content string, groupIndex int) string {
	match := pattern.FindStringSubmatch(content)
	if len(match) > groupIndex && groupIndex >= 0 {
		return match[groupIndex]
	}
	return ""
}

// Find returns a matching string
func Find(regex string, content string, groupIndex int) string {
	re := regexp.MustCompile(regex)
	match := re.FindStringSubmatch(content)
	if len(match) > groupIndex && groupIndex >= 0 {
		return match[groupIndex]
	}
	return ""
}

// FindAll returns the matched strings based on the given pattern and content
// with the specified group index.
func FindAll(regex string, content string, groupIndex int) []string {
	pattern := regexp.MustCompile(regex)
	matches := pattern.FindAllStringSubmatch(content, -1)
	var matchedStrings []string
	for _, match := range matches {
		if len(match) > groupIndex && groupIndex >= 0 {
			matchedStrings = append(matchedStrings, match[groupIndex])
		}
	}
	return matchedStrings
}

// IsMatch Determines whether the given string matches the regular expression pattern
func IsMatch(pattern string, content string) bool {
	match, _ := regexp.MatchString(pattern, content)
	return match
}

// DelFirst Deletes the first matching substring in the string
func DelFirst(str, sub string) string {
	index := strings.Index(str, sub)
	if index == -1 {
		return str
	}
	return str[:index] + str[index+len(sub):]
}
