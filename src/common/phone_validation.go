package common

import (
	"regexp"

	"github.com/Arshia-Izadyar/Car-sale-api/src/pkg/logging"
)

var regexString string = `^09(1[0-9]|2[0-2]|3[0-9]|9[0-9])[0-9]{7}$`
var logger = logging.NewLogger(cfg)

func ValidateNumber(pn string) bool {
	matched, err := regexp.Match(regexString, []byte(pn))
	if err != nil {
		logger.Error(err, logging.General, logging.MobileValidation, "cant match regex", nil)
	}
	return matched
}
