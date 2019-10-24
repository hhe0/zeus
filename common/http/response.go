package http

import "zeus/common/apicode"

type APIResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(data ...interface{}) APIResponse {
	return APIResponse{
		Code:    apicode.APISuccessCode.Code,
		Message: apicode.APISuccessCode.Message,
		Data:    data,
	}
}

func Error(code apicode.APICode, data ...interface{}) APIResponse {
	return APIResponse{
		Code:    code.Code,
		Message: code.Message,
		Data:    data,
	}
}
