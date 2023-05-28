package utils

const (
	SUCCESS = 200
	ERROR   = 500
	//code以100开头 用户模块错误
	ERROR_USERNAME_USED      = 1001
	ERROR_USERNAME_WRONG     = 1002
	ERROR_PASSWORD_WRONG     = 1003
	ERROR_TEL_WRONG          = 1004
	ERROR_TEL_USED           = 1005
	ERROR_CREAT_WRONG        = 1006
	ERROR_USERNAME_NOT_EXIST = 1007
	ERROR_EMAIL_WRONG        = 1008
	ERROR_EMAIL_CHECK        = 1009
	ERROR_EMAIL_NOT_EXIST    = 1010
	ERROR_CHANGE_WRONG       = 1011
	ERROR_UPLOAD_WRONG       = 1012

	//code以400开头 权限错误
	ERROR_USER_AUTH_NOT_ENOUGH = 4001
	ERROR_WRONG_CHANGE         = 4002
	ERROR_USER_PASSWORD_WRONG  = 4003
)

var codemsg = map[int]string{
	SUCCESS:                  "OK",
	ERROR:                    "FAIL",
	ERROR_USERNAME_USED:      "用户名已存在",
	ERROR_USERNAME_WRONG:     "用户名不符合",
	ERROR_PASSWORD_WRONG:     "密码错误",
	ERROR_TEL_WRONG:          "电话号码不符",
	ERROR_TEL_USED:           "该邮箱已被注册",
	ERROR_CREAT_WRONG:        "注册失败,请检查网络",
	ERROR_USERNAME_NOT_EXIST: "该用户名不存在",
	ERROR_EMAIL_WRONG:        "发送验证码错误",
	ERROR_EMAIL_CHECK:        "验证码错误",
	ERROR_EMAIL_NOT_EXIST:    "该邮箱未注册",
	ERROR_CHANGE_WRONG:       "修改错误",
	ERROR_UPLOAD_WRONG:       "上传失败",

	ERROR_USER_AUTH_NOT_ENOUGH: "当前用户权限不足",

	ERROR_WRONG_CHANGE: "当前无法删除业务元数据类型信息",

	ERROR_USER_PASSWORD_WRONG: "token密码错误，请重新登录",
}

func GetErrMsg(code int) string {
	return codemsg[code]
}
