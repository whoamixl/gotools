package HexUtil

import "encoding/hex"

// EncodeHexStr converts a string to its hexadecimal representation.
// Parameter str is the string to be converted.
// The return value string is the hexadecimal representation of the input string.
func EncodeHexStr(str string) string {
	bytes := []byte(str)
	hexStr := hex.EncodeToString(bytes)
	return hexStr
}

// DecodeHexStr converts a hexadecimal string to its original string representation.
// Parameter hexStr is the hexadecimal string to be converted.
// The return value string is the original string representation of the input hexadecimal string.
// The return value error represents any errors that occurred during the conversion process. It is nil if the conversion is successful.
func DecodeHexStr(hexStr string) (string, error) {
	bytes, err := hex.DecodeString(hexStr)
	if err != nil {
		return "", err
	}
	str := string(bytes)
	return str, nil
}
