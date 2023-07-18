package Constants

const (
	ParamParseErrorCode int    = 1001
	ParamParseErrorMsg  string = "未获取到参数!"

	ParamElementEmptyCode int    = 1002
	ParamUsernameEmptyMsg string = "用户名为空, 请重新输入!"
	ParamPasswordEmptyMsg string = "密码为空, 请重新输入!"

	UserNotFoundCode int    = 1003
	UserNotFoundMsg  string = "用户名或密码错误, 请重试!"

	UserDisableCode int    = 1004
	UserDisableMsg  string = "该用户被禁止使用!"
)
