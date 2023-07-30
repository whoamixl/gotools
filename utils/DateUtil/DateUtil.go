package DateUtil

import (
	"strings"
	"time"
)

const (
	FIRST_MEADOW  = "上旬"
	SECOND_MEADOW = "中旬"
	THIRD_MEADOW  = "下旬"
)

// Time is a struct that represents a time
type Time struct {
	time.Time
}

// Now returns the current time
func Now() Time {
	return Time{time.Now()}
}

// NowUnix returns the current time in Unix format
func NowUnix() int64 {
	return time.Now().Unix()
}

// NowUnixNano returns the current time in UnixNano format
func NowUnixNano() int64 {
	return time.Now().UnixNano()
}

// NowString returns the current time in a string format
func NowString() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// Format formats the time using the given string
func (t Time) Format(str string) string {
	layout := build(str)
	const bufSize = 64
	var b []byte
	max := len(layout) + 10
	if max < bufSize {
		var buf [bufSize]byte
		b = buf[:0]
	} else {
		b = make([]byte, 0, max)
	}
	b = t.AppendFormat(b, layout)
	return string(b)
}

// FormatByTime formats the time using the given time and string
func FormatByTime(t time.Time, str string) string {
	layout := build(str)
	return t.Format(layout)
}

// GetMeadow  returns the current time in the meadow format
func (t Time) GetMeadow() string {
	d := t.Day()
	return buildMeadow(d)
}

// Parse parses the given string using the given layout
func Parse(layout, value string) (time.Time, error) {
	layout = build(layout)
	return time.Parse(layout, value)
}

// Day Return the day of the month of the given time
func Day(t time.Time) int {
	return t.Day()
}

// Weekday Return the day of the week of the given time
func Weekday(t time.Time) time.Weekday {
	return t.Weekday()
}

// YearDay Return the day of the year of the given time
func YearDay(t time.Time) int {
	return t.YearDay()
}

// Meadow Return the day of the year of the given time
func Meadow(t time.Time) string {
	return buildMeadow(t.Day())
}

// SubDay calculate the difference between two times in hours
func SubDay(t1, t2 time.Time) float64 {
	return t1.Sub(t2).Hours() / 24
}

// build formats the given string
func build(str string) string {
	str = strings.Replace(str, "yyyy", "2006", -1)
	str = strings.Replace(str, "yy", "06", -1)
	str = strings.Replace(str, "MM", "01", -1)
	str = strings.Replace(str, "dd", "02", -1)
	str = strings.Replace(str, "HH", "15", -1)
	str = strings.Replace(str, "mm", "04", -1)
	str = strings.Replace(str, "ss", "05", -1)
	str = strings.Replace(str, "SSS", "000", -1)
	str = strings.Replace(str, "S", "0", -1)
	str = strings.Replace(str, "ZZZ", "0700", -1)
	return str
}

// buildMeadow formats the given string
func buildMeadow(day int) string {
	if day <= 10 {
		return FIRST_MEADOW
	} else if day <= 20 {
		return SECOND_MEADOW
	} else {
		return THIRD_MEADOW
	}

}
