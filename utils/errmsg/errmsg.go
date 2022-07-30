package errmsg

const (
	SUCCESS               = 200
	ERROR                 = 500
	ERROR_USERNAME_USED   = 1001
	ERROR_PASSWORD_WRONG  = 1002
	ERROR_USER_NOT_EXIST  = 1003
	ERROR_TOKEN_NOT_EXIST = 1004
	ERROT_TOKEN_RUNTIME   = 1005
	ERROT_TOKEN_WRONG     = 1006
)

var codeMsg = map[int]string{
	SUCCESS:               "OK",
	ERROR:                 "FAIL",
	ERROR_USERNAME_USED:   "用户名已存在",
	ERROR_PASSWORD_WRONG:  "密码错误",
	ERROR_USER_NOT_EXIST:  "用户名不存在",
	ERROR_TOKEN_NOT_EXIST: "token不存在",
	ERROT_TOKEN_RUNTIME:   "token已过期",
	ERROT_TOKEN_WRONG:     "token格式错误",
}

// GetErrMsg 错误代码
func GetErrMsg(code int) string {
	return codeMsg[code]
}
