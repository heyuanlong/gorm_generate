package code

// 不可变参数

const (
	KEY    = "xxxxxxxxxxxxxxx"
	SECRET = "xxxxxxxxxxxxxxxxxxxxxxxxxxxx"

	//token

)

//-----------------------------------------------------

var codeToMsg map[int]string
var codeToChnMsg map[int]string

const (
	SUCCESS_STATUS = 200

	OPERATION_WRONG       = 20001
	ACCESS_TOKEN_FAIL     = 20002
	ACCESS_TOKEN_EXPIRE   = 20003
	PARAM_WRONG           = 20004
	WRONG_PASSWORD        = 20005
	USER_NOT_NORMAL       = 20006
	MOBILE_EXIST          = 20007
	PASSWD_TOO_SHORT      = 20008
	INVITE_CODE_NOT_EXIST = 20009
	MOBILE_NOT_OK         = 20010
	VERIFYCODE_EXPIRE     = 20011
	VERIFYCODE_ERROR      = 20012
	SMS_SEND_FAIL         = 20013
	SMSCODE_EXPIRE        = 20014
	SMSCODE_ERROR         = 20015
	USER_NOT_EXIST        = 20016
)

func init() {

	codeToChnMsg = make(map[int]string)
	codeToChnMsg[SUCCESS_STATUS] = "成功"
	codeToChnMsg[OPERATION_WRONG] = "操作错误"
	codeToChnMsg[ACCESS_TOKEN_FAIL] = "token错误"
	codeToChnMsg[ACCESS_TOKEN_EXPIRE] = "token过期"
	codeToChnMsg[PARAM_WRONG] = "参数有误"
	codeToChnMsg[WRONG_PASSWORD] = "请输入正确的密码"
	codeToChnMsg[USER_NOT_NORMAL] = "该用户非正常状态"
	codeToChnMsg[MOBILE_EXIST] = "此手机号已被注册"
	codeToChnMsg[PASSWD_TOO_SHORT] = "请按要求设置密码"
	codeToChnMsg[INVITE_CODE_NOT_EXIST] = "邀请码不存在"
	codeToChnMsg[MOBILE_NOT_OK] = "手机号码格式错误"
	codeToChnMsg[VERIFYCODE_EXPIRE] = "图片验证码过期"
	codeToChnMsg[VERIFYCODE_ERROR] = "图片验证码错误"
	codeToChnMsg[SMS_SEND_FAIL] = "短信发送失败"
	codeToChnMsg[SMSCODE_EXPIRE] = "短信验证码过期"
	codeToChnMsg[SMSCODE_ERROR] = "短信验证码错误"
	codeToChnMsg[USER_NOT_EXIST] = "账号尚未注册"

}
func GetCodeMsg(code int) string {
	if msg, ok := codeToMsg[code]; ok {
		return msg
	}
	return ""
}
func GetCodeChnMsg(code int) string {
	if msg, ok := codeToChnMsg[code]; ok {
		return msg
	}
	return ""
}

//--------------------------------------------------------------------------------------
var typeToMsg map[int]string
var typeToChnMsg map[int]string

const (
	Type_test  = 1  //
	Type_testx = -1 //
)

func init() {
	typeToMsg = make(map[int]string)
	typeToChnMsg = make(map[int]string)

	typeToChnMsg[Type_test] = "test"
	typeToChnMsg[Type_testx] = "testx"

}

func GetTypeMsg(types int) string {
	if msg, ok := typeToMsg[types]; ok {
		return msg
	}
	return ""
}

func GetTypeChnMsg(types int) string {
	if msg, ok := typeToChnMsg[types]; ok {
		return msg
	}
	return ""
}

//---------------------------------------------------------------------------------------------
