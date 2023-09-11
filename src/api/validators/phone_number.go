package validators

import (
	"github.com/Arshia-Izadyar/Car-sale-api/src/common"
	"github.com/go-playground/validator/v10"
)

func IranPhoneNumberValidator(fld validator.FieldLevel) bool {
	value, ok := fld.Field().Interface().(string)
	if !ok {
		return false
	}
	return common.ValidateNumber(value)
}
