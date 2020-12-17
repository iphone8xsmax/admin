package util


//检查切片中是否包含字符串，用于去重
func IsContains(a string, b []string) bool {
	for _, v := range b{
		if a == v{
			return true
		}
	}
	return false
}