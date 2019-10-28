package apicode

type APICode struct {
	Code    int
	Message string
}

// 请求成功返回码
var APISuccessCode = APICode{999999, "Success"}

// 公共错误码
var (
	InputError = APICode{100001, "输入参数有误"}
)
