package helper

import "github.com/Stevesadr/golang-backend-project/api/validations"

type BaseResponse struct{
	Result any `json:"result"`
	Success bool `json:"success"`
	ResultCode int `json:"resultCode"`
	ValidationError *[]validations.ValidationError
	Error any `json:"error"`
}

func GenerateResponse(res any, success bool, resCod int) *BaseResponse {
	return &BaseResponse{
		Result: res,
		Success: success,
		ResultCode: resCod,
	}
}

func GenerateResponseWithError(res any, success bool, resCod int, err error) *BaseResponse {
	return &BaseResponse{
		Result: res,
		Success: success,
		ResultCode: resCod,
		Error: err.Error(),
	}
}
func GenerateResponseWithValidation(res any, success bool, resCod int, err error) *BaseResponse {
	return &BaseResponse{
		Result: res,
		Success: success,
		ResultCode: resCod,
		Error: validations.GetValidationError(err) ,
	}
}