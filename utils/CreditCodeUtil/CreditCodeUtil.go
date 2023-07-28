package CreditCodeUtil

import (
	"math/rand"
	"regexp"
	"strings"
	"time"
)

// CheckSocietyId your society id
func CheckSocietyId(str string) bool {
	return checkSocialCreditCode(str)
}

func GenerateRandomSocialCreditCode() string {
	rand.Seed(time.Now().UnixNano())
	code := ""
	// 第一部分：登记管理部门代码
	departmentCode := generateRandomCharacter()
	code += departmentCode
	// 第二部分：机构类别代码
	categoryCode := generateRandomCharacter()
	code += categoryCode
	// 第三部分：登记管理机关行政区划码
	adminDivisionCode := generateRandomNumberString(6)
	code += adminDivisionCode
	// 第四部分：主体标识码（组织机构代码）
	organizationCode := generateRandomAlphanumericString(9)
	code += organizationCode
	// 第五部分：校验码
	checkCode := generateRandomCharacter()
	code += checkCode
	return code
}
func generateRandomCharacter() string {
	characters := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	index := rand.Intn(len(characters))
	return string(characters[index])
}
func generateRandomNumberString(length int) string {
	numberString := ""
	for i := 0; i < length; i++ {
		numberString += generateRandomCharacter()
	}
	return numberString
}
func generateRandomAlphanumericString(length int) string {
	alphanumericString := ""
	for i := 0; i < length; i++ {
		// 70%的概率生成数字，30%的概率生成大写字母
		if rand.Float32() < 0.7 {
			alphanumericString += generateRandomCharacter()
		} else {
			alphanumericString += generateRandomNumberString(1)
		}
	}
	return alphanumericString
}
func checkSocialCreditCode(code string) bool {
	// 去除空格和连字符
	code = strings.ReplaceAll(code, " ", "")
	code = strings.ReplaceAll(code, "-", "")
	// 校验长度
	if len(code) != 18 {
		return false
	}
	// 校验格式
	codeRegex := regexp.MustCompile(`^[0-9A-Z]+$`)
	if !codeRegex.MatchString(code) {
		return false
	}
	// 校验第一部分
	if !isValidDepartmentCode(code[:1]) {
		return false
	}
	// 校验第二部分
	if !isValidCategoryCode(code[1:2]) {
		return false
	}
	// 校验第三部分
	if !isValidAdminDivisionCode(code[2:8]) {
		return false
	}
	// 校验第四部分
	if !isValidOrganizationCode(code[8:17]) {
		return false
	}
	// 校验第五部分
	if !isValidCheckCode(code[17:]) {
		return false
	}
	return true
}

// 校验登记管理部门代码
func isValidDepartmentCode(code string) bool {
	departmentRegex := regexp.MustCompile(`^[0-9A-Z]$`)
	return departmentRegex.MatchString(code)
}

// 校验机构类别代码
func isValidCategoryCode(code string) bool {
	categoryRegex := regexp.MustCompile(`^[0-9A-Z]$`)
	return categoryRegex.MatchString(code)
}

// 校验登记管理机关行政区划码
func isValidAdminDivisionCode(code string) bool {
	adminDivisionRegex := regexp.MustCompile(`^[0-9]{6}$`)
	return adminDivisionRegex.MatchString(code)
}

// 校验主体标识码（组织机构代码）
func isValidOrganizationCode(code string) bool {
	organizationRegex := regexp.MustCompile(`^[0-9A-Z]{9}$`)
	return organizationRegex.MatchString(code)
}

// 校验校验码
func isValidCheckCode(code string) bool {
	checkCodeRegex := regexp.MustCompile(`^[0-9A-Z]$`)
	return checkCodeRegex.MatchString(code)
}
