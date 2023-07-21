package NumberUtil

import (
	"math"
	"math/big"
	"strconv"
)

// Add adds two floating-point numbers and returns the result and accuracy information.
// Parameters v1 and v2 are the two numbers to be added.
// The return value float64 is the floating-point representation of the sum result
func Add(v1, v2 float64) float64 {
	precision := math.Pow10(10)
	num1 := int64(v1 * precision)
	num2 := int64(v2 * precision)
	sum := num1 + num2
	result := float64(sum) / precision
	return result
}

// Sub subtracts two floating-point numbers and returns the result and accuracy information.
// Parameters v1 and v2 are the two numbers to be subtracted.
// The return value float64 is the floating-point representation of the difference result
func Sub(v1, v2 float64) float64 {
	precision := math.Pow10(10)
	num1 := int64(v1 * precision)
	num2 := int64(v2 * precision)
	sum := num1 - num2
	result := float64(sum) / precision
	return result
}

// Mul multiplies two floating-point numbers and returns the result and accuracy information.
// Parameters v1 and v2 are the two numbers to be multiplied.
// The return value float64 is the floating-point representation of the product result
func Mul(v1, v2 float64) float64 {
	prec := 64 // Set the desired precision here
	f1 := big.NewFloat(v1)
	f2 := big.NewFloat(v2)
	result := new(big.Float).Mul(f1, f2).SetPrec(uint(prec))
	f, _ := result.Float64()
	return f
}

// Div divides two floating-point numbers and returns the result and accuracy information.
// Parameters v1 and v2 are the two numbers to be divided.
// The return value float64 is the floating-point representation of the division result
func Div(v1, v2 float64) float64 {
	precision := 10 // Set the desired precision here
	scale := math.Pow(10, float64(precision))
	result := math.Round(v1/v2*scale) / scale
	return result
}

// RoundFloat rounds a float64 number to a specific number of decimal places.
// It takes a float64 number 'num' and an integer 'decimalPlaces' as parameters.
// The function returns the rounded float64 value.
func RoundFloat(num float64, decimalPlaces int) float64 {
	shift := math.Pow(10, float64(decimalPlaces))
	rounded := math.Round(num*shift) / shift
	return rounded
}

// TruncateFloat truncates a float64 number to a specific number of decimal places using the truncation mode.
// It takes a float64 number 'num' and an integer 'decimalPlaces' as parameters.
// The function returns the truncated float64 value.
func TruncateFloat(num float64, decimalPlaces int) float64 {
	shift := math.Pow(10, float64(decimalPlaces))
	truncated := math.Trunc(num*shift) / shift
	return truncated
}

// RoundFloatString rounds a numeric string to a specified number of decimal places (rounding half up).
// It takes a string representation of a numeric value 'numString' and an integer 'decimalPlaces' as parameters.
// The function returns a float64 value with the specified number of decimal places.
func RoundFloatString(numString string, decimalPlaces int) (float64, error) {
	num, err := strconv.ParseFloat(numString, 64)
	if err != nil {
		return 0, err
	}
	shift := math.Pow(10, float64(decimalPlaces))
	rounded := math.Round(num*shift) / shift
	return rounded, nil
}

// TruncateFloatString truncates a numeric string to a specified number of decimal places.
// It takes a string representation of a numeric value 'numString' and an integer 'decimalPlaces' as parameters.
// The function returns a float64 value with the specified number of decimal places.
func TruncateFloatString(numString string, decimalPlaces int) (float64, error) {
	num, err := strconv.ParseFloat(numString, 64)
	if err != nil {
		return 0, err
	}
	shift := math.Pow(10, float64(decimalPlaces))
	truncated := math.Trunc(num*shift) / shift
	return truncated, nil
}
