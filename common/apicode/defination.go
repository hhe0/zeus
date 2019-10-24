package apicode

type APICode struct {
	Code    int
	Message string
}

var (
	APISuccessCode = APICode{999999, "Success"}

	InputError = APICode{100001, "输入参数有误"}
)
