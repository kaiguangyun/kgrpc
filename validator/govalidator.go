package validator

import (
	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)

	mobilePhone()
	alphaDash()
}

func mobilePhone() {
	govalidator.CustomTypeTagMap.Set("mobilePhone", govalidator.CustomTypeValidator(func(i interface{}, o interface{}) bool {
		return IsMobilePhone(i.(string))
	}))
}

func alphaDash() {
	govalidator.CustomTypeTagMap.Set("alphaDash", govalidator.CustomTypeValidator(func(i interface{}, o interface{}) bool {
		return IsAlphaDash(i.(string))
	}))
}
