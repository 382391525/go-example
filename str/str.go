package str

import "strings"

//判断字符串 是否包含子串
func Contains(str, substr string) bool {
	return strings.Contains(str, substr)
}

//字符串拼接
func Join(str []string, sep string) string {
	return strings.Join(str, sep)
}

//字符串切割
func Explode(str, sep string) []string {
	return strings.Split(str, sep)
}
