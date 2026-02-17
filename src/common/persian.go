package common

import (
	"regexp"

	"github.com/Stevesadr/golang-backend-project/config"
	"github.com/Stevesadr/golang-backend-project/pkg/logging"
)

var logger = logging.NewLogger(config.GetConfig())

const IranianMobilePatternString = `^09(1[0-9]|2[0-2]|3[0-9]|9[0-9])[0-9]{7}$`

func IranianMobileNumberValidation(mobileNumber string) bool {
	res, err := regexp.MatchString(IranianMobilePatternString, mobileNumber)
	if err != nil{
		logger.Error(logging.Validation, logging.MobileValidation, err.Error(), nil)
	}
	return res
}