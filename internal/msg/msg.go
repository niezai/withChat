package msg

var Err_Map_Msg = map[int]string{
	ERR_PARAMETER:  "参数错误",
	ERR_LOGIN_FAIL: "登录失败,请检查你的用户名信息是否正确,或者请稍后登录",
	FAIL:           "请求失败",
	SUCCESS:        "请求成功",
}

func GetMsg(code int) string {
	if msg, ok := Err_Map_Msg[code]; ok {
		return msg
	}
	return "-1\n"
}
