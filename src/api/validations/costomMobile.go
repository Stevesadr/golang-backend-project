package validations

import (
	"log"
	"regexp"
	"github.com/go-playground/validator/v10"
)

const regex = `^09(1[0-9]|2[0-2]|3[0-9]|9[0-9])[0-9]{7}$`

func IranianMobileNumberValidation(fl validator.FieldLevel) bool{
	mobile, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}

	res, err:= regexp.MatchString(regex, mobile)
	if err != nil{
		log.Print(err.Error())
	}
	return res
}