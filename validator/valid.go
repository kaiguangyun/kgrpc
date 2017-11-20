package validator

import (
	"github.com/asaskevich/govalidator"
	"regexp"
)

// IsNull check if the string is null.
func IsNull(str string) bool {
	return len(str) == 0
}

// mobile phone
func IsMobilePhone(value string) bool {
	reg := `^1([3-9][0-9]|14[57]|5[^4])\d{8}$`

	return regexp.MustCompile(reg).MatchString(value)
}

// IsFixedPhone
func IsFixedPhone(value string) bool {
	reg := `^(0[0-9]{2,3}-)?([2-9][0-9]{6,7})+(-[0-9]{1,4})?$`

	return regexp.MustCompile(reg).MatchString(value)
}

// bank account
func IsBankAccount(value string) bool {
	reg := `^([1-9]{1})(\d{14}|\d{18})$`

	return regexp.MustCompile(reg).MatchString(value)
}

// idcard
func IsIdcard(value string) bool {
	if len(value) == 15 {
		reg := `^[1-9]\d{5}\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{2}$`

		return regexp.MustCompile(reg).MatchString(value)
	} else {
		reg := `^[1-9]\d{5}(18|19|([23]\d))\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$`

		return regexp.MustCompile(reg).MatchString(value)
	}
	return false
}

// chinese
func IsChinese(value string) bool {
	reg := `^[\p{Han}]+$`

	return regexp.MustCompile(reg).MatchString(value)
}

// postcode
func IsPostcode(value string) bool {
	reg := `^[1-9][0-9]{5}$`

	return regexp.MustCompile(reg).MatchString(value)
}

// numbers
func IsNumeric(value string) bool {
	reg := `^[0-9]+$`

	return regexp.MustCompile(reg).MatchString(value)
}

// letters
func IsAlpha(value string) bool {
	reg := `^[a-zA-Z]+$`

	return regexp.MustCompile(reg).MatchString(value)
}

// letters and numbers
func IsAlphaNumeric(value string) bool {
	reg := `^[a-zA-Z0-9]+$`

	return regexp.MustCompile(reg).MatchString(value)
}

// letters and numbers and - and  _
func IsAlphaDash(value string) bool {
	reg := `^[a-zA-Z0-9_-]+$`

	return regexp.MustCompile(reg).MatchString(value)
}

// url
func IsUrl(value string) bool {
	return govalidator.IsURL(value)
}

// email
func IsEmail(value string) bool {
	return govalidator.IsEmail(value)
}
