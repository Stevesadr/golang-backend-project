package common

import (
	"log"
	"regexp"
)

const IranianMobilePatternString = `^09(1[0-9]|2[0-2]|3[0-9]|9[0-9])[0-9]{7}$`

func IranianMobileNumberValidation(mobileNumber string) bool {
	res, err := regexp.MatchString(IranianMobilePatternString, mobileNumber)
	if err != nil{
		log.Print(err.Error())
	}
	return res
}