package config

const (
	DefaultError = 0
	Success = 10000
	UnLogin = 10001
	NoEnoughAuth = 10002

	ParamError = 10103

	SysError = 10201
)

var statusTitle = map[int] string{
	Success: "success",
	DefaultError: "unknown",
	UnLogin: "need login",
	ParamError: "参数错误或缺失",
	NoEnoughAuth: "权限不足",
	SysError: "系统内部错误，请稍候重试",
}

func GetErrorCode(code int) int {
	_, ok := statusTitle[code]
	if ok {
		return code
	}
	return DefaultError
}

func GetErrorTitle(code int) string {
	return statusTitle[code]
}
/*func GetErrorTitle2(code int) string {
	str, ok := statusTitle[code]
	if ok {
		return str
	}
	return statusTitle[code]
}*/