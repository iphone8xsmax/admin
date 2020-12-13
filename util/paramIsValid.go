package util

import "regexp"

//正则验证电阻邮箱是否合法 ***@***.***
func CheckEamil(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}


