package NumberUtil

import (
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"strconv"
	"strings"
	"time"
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

// IsNumber checks if a string is a valid number.
// It returns true if the string is a number, otherwise false.
func IsNumber(input string) bool {
	_, err := strconv.ParseFloat(input, 64)
	return err == nil
}

// IsInteger checks if a given string value is an integer.
// It returns true if the value is an integer, otherwise false.
func IsInteger(value string) bool {
	_, err := strconv.Atoi(value)
	return err == nil
}

// IsFloat checks if a given string value is a float.
// It returns true if the value is a float, otherwise false.
func IsFloat(value string) bool {
	_, err := strconv.ParseFloat(value, 64)
	return err == nil
}

// IsPrime checks if a given number is prime.
// It returns true if the number is prime, otherwise false.
func IsPrime(num int) bool {
	if num <= 1 {
		return false
	}
	if num == 2 {
		return true
	}
	if num%2 == 0 {
		return false
	}
	sqrt := int(math.Sqrt(float64(num)))
	for i := 3; i <= sqrt; i += 2 {
		if num%i == 0 {
			return false
		}
	}
	return true
}

// GenerateRandomNumber generates a random number between min and max (inclusive).
func GenerateRandomNumber(min, max int) int {
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	return rand.Intn(max-min+1) + min
}

// GenerateRandomNumberSlice generates a slice of random numbers between min and max (inclusive).
func GenerateRandomNumberSlice(min, max, size int) []int {
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	numbers := make([]int, size)
	for i := 0; i < size; i++ {
		numbers[i] = rand.Intn(max-min+1) + min
	}
	return numbers
}

// GenerateRandomFloat generates a random number between min and max (inclusive).
func GenerateRandomFloat(min, max float64) float64 {
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	return min + rand.Float64()*(max-min)
}

// GenerateRandomFloatSlice generates a slice of random numbers between min and max (inclusive).
func GenerateRandomFloatSlice(min, max float64, size int) []float64 {
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	numbers := make([]float64, size)
	for i := 0; i < size; i++ {
		numbers[i] = min + rand.Float64()*(max-min)
	}
	return numbers
}

// Range generates a list of ordered integers based on the range and step.
func Range(start, end, step int) []int {
	if step <= 0 {
		return nil
	}
	var result []int
	for i := start; i <= end; i += step {
		result = append(result, i)
	}
	return result
}

// Factorial calculates the factorial of the given number.
func Factorial(n int) int {
	if n < 0 {
		return 0
	} else if n == 0 {
		return 1
	} else {
		result := 1
		for i := 1; i <= n; i++ {
			result *= i
		}
		return result
	}
}

// Divisor calculates the greatest common divisor (GCD) of two integers.
func Divisor(a, b int) int {
	if b == 0 {
		return a
	}
	return Divisor(b, a%b)
}

// Multiple calculates the least common multiple (LCM) of two numbers.
func Multiple(a, b int) int {
	return (a * b) / Divisor(a, b)
}

// GetBinaryStr converts an integer to its binary representation as a string.
func GetBinaryStr(num int) string {
	return strconv.FormatInt(int64(num), 2)
}

// BinaryToInt converts a binary string to an integer.
func BinaryToInt(binaryStr string) (int64, error) {
	return strconv.ParseInt(binaryStr, 2, 64)
}

// FloatToString converts a float64 number to a string and removes trailing zeros after the decimal point.
func FloatToString(num float64) string {
	str := fmt.Sprintf("%.2f", num)   // Convert the float to string with 2 decimal places
	str = strings.TrimRight(str, "0") // Remove trailing zeros
	str = strings.TrimRight(str, ".") // Remove decimal point if no decimal places left
	return str
}
