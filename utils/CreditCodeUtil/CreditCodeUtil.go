// Package CreditCodeUtil 社会信用代码工具
package CreditCodeUtil

import (
	"github.com/whoamixl/gotools/utils"
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
	rand.New(rand.NewSource(time.Now().UnixNano()))
	code := ""
	departmentCode := generateRandomCharacter(utils.ALL_TYPE)
	code += departmentCode
	categoryCode := generateRandomCharacter(utils.ALL_TYPE)
	code += categoryCode
	adminDivisionCode := generateRandomNumberString(6, utils.NUMBER)
	code += adminDivisionCode
	organizationCode := generateRandomAlphanumericString(9)
	code += organizationCode
	checkCode := generateRandomCharacter(utils.ALL_TYPE)
	code += checkCode
	return code
}
func generateRandomCharacter(ty int) string {
	characters := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	if utils.NUMBER == ty {
		characters = "0123456789"
	} else if utils.ALL_TYPE == ty {
		characters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	index := rand.Intn(len(characters))
	return string(characters[index])
}
func generateRandomNumberString(length, ty int) string {
	numberString := ""
	for i := 0; i < length; i++ {
		numberString += generateRandomCharacter(ty)
	}
	return numberString
}
func generateRandomAlphanumericString(length int) string {
	alphanumericString := ""
	for i := 0; i < length; i++ {
		if rand.Float32() < 0.7 {
			alphanumericString += generateRandomCharacter(utils.NUMBER)
		} else {
			alphanumericString += generateRandomNumberString(1, utils.ALL_TYPE)
		}
	}
	return alphanumericString
}
func checkSocialCreditCode(code string) bool {
	code = strings.ReplaceAll(code, " ", "")
	code = strings.ReplaceAll(code, "-", "")
	if len(code) != 18 {
		return false
	}
	codeRegex := regexp.MustCompile(`^[0-9A-Z]+$`)
	if !codeRegex.MatchString(code) {
		return false
	}
	if !isValidDepartmentCode(code[:1]) {
		return false
	}
	if !isValidCategoryCode(code[1:2]) {
		return false
	}
	if !isValidAdminDivisionCode(code[2:8]) {
		return false
	}
	if !isValidOrganizationCode(code[8:17]) {
		return false
	}
	if !isValidCheckCode(code[17:]) {
		return false
	}
	return true
}

func isValidDepartmentCode(code string) bool {
	departmentRegex := regexp.MustCompile(`^[0-9A-Z]$`)
	return departmentRegex.MatchString(code)
}

func isValidCategoryCode(code string) bool {
	categoryRegex := regexp.MustCompile(`^[0-9A-Z]$`)
	return categoryRegex.MatchString(code)
}

func isValidAdminDivisionCode(code string) bool {
	adminDivisionRegex := regexp.MustCompile(`^[0-9]{6}$`)
	return adminDivisionRegex.MatchString(code)
}

func isValidOrganizationCode(code string) bool {
	organizationRegex := regexp.MustCompile(`^[0-9A-Z]{9}$`)
	return organizationRegex.MatchString(code)
}

func isValidCheckCode(code string) bool {
	checkCodeRegex := regexp.MustCompile(`^[0-9A-Z]$`)
	return checkCodeRegex.MatchString(code)
}
