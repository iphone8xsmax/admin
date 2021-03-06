package e

//api错误码
var MsgFlags = map[int]string{
	SUCCESS : "ok",
	ERROR : "fail",
	INVALID_PARAMS : "请求参数错误",

	ERROR_USER_CHECK_TOKEN_FAIL : "Token鉴权失败",
	ERROR_USER_CHECK_TOKEN_TIMEOUT : "Token已超时",
	ERROR_USER_TOKEN : "Token生成失败",
	ERROR_USER : "Token错误",
}

//根据传入的code查询对于的字符串信息，不存在返回默认错误
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}