package errmsg

const (
	SUCCESS = 200
	ERROR   = 500

	//1000 用户模块错误码
	ERROR_USERNAME_USED    = 1001
	ERROR_PASSWORD_WRONG   = 1002
	ERROR_USER_NOT_EXIST   = 1003
	ERROR_TOKEN_NOT_EXIST  = 1004
	ERROT_TOKEN_RUNTIME    = 1005
	ERROT_TOKEN_WRONG      = 1006
	ERROR_TOKEN_TYPE_WRONG = 1007

	//2000 分类模块错误代码
	ERROR_CATENAME_USED = 2001

	//3000 文章模块错误码
)

var codeMsg = map[int]string{
	SUCCESS:                "OK",
	ERROR:                  "FAIL",
	ERROR_USERNAME_USED:    "用户名已存在",
	ERROR_PASSWORD_WRONG:   "密码错误",
	ERROR_USER_NOT_EXIST:   "用户名不存在",
	ERROR_TOKEN_NOT_EXIST:  "token不存在",
	ERROT_TOKEN_RUNTIME:    "token已过期",
	ERROT_TOKEN_WRONG:      "token不正确",
	ERROR_TOKEN_TYPE_WRONG: "token格式错误",
	ERROR_CATENAME_USED:    "分类已经存在",
}

// GetErrMsg 错误代码
func GetErrMsg(code int) string {
	return codeMsg[code]
}
