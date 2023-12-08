package e

var ErrMapMsg = map[int]string{
	LoginFail:         "登录失败,请检查你的用户名信息是否正确,或者请稍后登录",
	RegisterFail:      "注册失败",
	FAIL:              "请求失败",
	SUCCESS:           "请求成功",
	UnLogin:           "未登录",
	CertificationFail: "身份验证失效,请重新登录",
}

func GetMsg(code int) string {
	if msg, ok := ErrMapMsg[code]; ok {
		return msg
	}
	return "-1\n"
}
