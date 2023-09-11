package helper

import "github.com/Arshia-Izadyar/Car-sale-api/src/api/validators"

type Response struct {
	Err              error
	Result           interface{}
	StatusCode       int
	Success          bool
	ValidationErrors *[]validators.ValidationError
}

func GenerateBaseResponse(r interface{}, c int, s bool) *Response {
	return &Response{Result: r, StatusCode: c, Success: s}
}
func GenerateBaseResponseWithError(result any, success bool, resultCode int, err error) *Response {
	return &Response{
		Result:     result,
		Success:    success,
		StatusCode: resultCode,
		Err:        err,
	}
}

func GenerateBaseResponseWithValidationError(result any, success bool, resultCode int, err error) *Response {
	ve := validators.GetValidationError(err)
	return &Response{
		Result:           result,
		Success:          success,
		StatusCode:       resultCode,
		ValidationErrors: ve,
	}
}
