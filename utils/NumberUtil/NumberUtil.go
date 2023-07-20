package NumberUtil

import (
	"math"
)

// Add adds two floating-point numbers and returns the result and accuracy information.
// Parameters v1 and v2 are the two numbers to be added.
// The return value float64 is the floating-point representation of the sum result
func Add(v1 float64, v2 float64) float64 {
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
func Sub(v1 float64, v2 float64) float64 {
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
func Mul(v1 float64, v2 float64) float64 {
	precision := math.Pow10(10)
	num1 := int64(v1 * precision)
	num2 := int64(v2 * precision)
	sum := num1 * num2
	result := float64(sum) / precision
	return result
}

// Div divides two floating-point numbers and returns the result and accuracy information.
// Parameters v1 and v2 are the two numbers to be divided.
// The return value float64 is the floating-point representation of the division result
func Div(v1 float64, v2 float64) float64 {
	precision := math.Pow10(10)
	num1 := int64(v1 * precision)
	num2 := int64(v2 * precision)
	sum := num1 / num2
	result := float64(sum) / precision
	return result
}
