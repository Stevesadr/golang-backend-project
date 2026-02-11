package validations

import (
	"github.com/Stevesadr/golang-backend-project/common"
	"github.com/go-playground/validator/v10"
)

func IranianMobileNumberValidation(fld validator.FieldLevel) bool {
	value, ok:= fld.Field().Interface().(string)
	if !ok{
		return false
	}
	return common.IranianMobileNumberValidation(value)
}