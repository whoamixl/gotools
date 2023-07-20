package NumberUtil

import "math/big"

// Add adds two floating-point numbers and returns the result and accuracy information.
// Parameters v1 and v2 are the two numbers to be added.
// The return value float64 is the floating-point representation of the sum result, and the return value big.Accuracy represents the accuracy information of the result.
func Add(v1 float64, v2 float64) (float64, big.Accuracy) {
	num1 := new(big.Float).SetFloat64(v1)
	num2 := new(big.Float).SetFloat64(v2)
	sum := new(big.Float).Add(num1, num2)
	return sum.Float64()
}

// Sub subtracts two floating-point numbers and returns the result and accuracy information.
// Parameters v1 and v2 are the two numbers to be subtracted.
// The return value float64 is the floating-point representation of the difference result, and the return value big.Accuracy represents the accuracy information of the result.
func Sub(v1 float64, v2 float64) (float64, big.Accuracy) {
	num1 := new(big.Float).SetFloat64(v1)
	num2 := new(big.Float).SetFloat64(v2)
	sum := new(big.Float).Sub(num1, num2)
	return sum.Float64()
}

// Mul multiplies two floating-point numbers and returns the result and accuracy information.
// Parameters v1 and v2 are the two numbers to be multiplied.
// The return value float64 is the floating-point representation of the product result, and the return value big.Accuracy represents the accuracy information of the result.
func Mul(v1 float64, v2 float64) (float64, big.Accuracy) {
	num1 := new(big.Float).SetFloat64(v1)
	num2 := new(big.Float).SetFloat64(v2)
	sum := new(big.Float).Mul(num1, num2)
	return sum.Float64()
}

// Div divides two floating-point numbers and returns the result and accuracy information.
// Parameters v1 and v2 are the two numbers to be divided.
// The return value float64 is the floating-point representation of the division result, and the return value big.Accuracy represents the accuracy information of the result.
func Div(v1 float64, v2 float64) (float64, big.Accuracy) {
	num1 := new(big.Float).SetFloat64(v1)
	num2 := new(big.Float).SetFloat64(v2)
	sum := new(big.Float).Quo(num1, num2)
	return sum.Float64()
}
