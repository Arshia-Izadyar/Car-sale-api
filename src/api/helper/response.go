package helper

import "github.com/Arshia-Izadyar/Car-sale-api/src/api/validators"

type Response struct {
	Result           any                           `json:"result"`
	Success          bool                          `json:"success"`
	ResultCode       int                           `json:"result_code"`
	ValidationErrors *[]validators.ValidationError `json:"validation_error"`
	Error            any                           `json:"error"`
}

func GenerateBaseResponse(result any, resultCode int, success bool) *Response {
	return &Response{
		Result:     result,
		Success:    success,
		ResultCode: resultCode,
	}
}
func GenerateBaseResponseWithError(result any, success bool, resultCode int, err string) *Response {
	return &Response{
		Result:     result,
		Success:    success,
		ResultCode: resultCode,
		Error:      err,
	}
}

func GenerateBaseResponseWithValidationError(result any, success bool, resultCode int, err error) *Response {
	ve := validators.GetValidationError(err)
	return &Response{
		Result:           result,
		Success:          success,
		ResultCode:       resultCode,
		ValidationErrors: ve,
	}
}
