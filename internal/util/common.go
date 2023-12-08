package util

// 判断一个字符chuan
func CheckExistsArr(target string, arr []string) bool {
	for _, str := range arr {
		if str == target {
			return true
		}
	}
	return false
}
