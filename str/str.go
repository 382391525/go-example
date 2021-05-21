package str

import "strings"

// Contains 判断字符串 是否包含子串
func Contains(str, substr string) bool {
	return strings.Contains(str, substr)
}

// Join 字符串拼接
func Join(str []string, sep string) string {
	return strings.Join(str, sep)
}

// Explode 字符串切割
func Explode(str, sep string) []string {
	return strings.Split(str, sep)
}

// Splicing 字符串拼接
func Splicing(s ...string) string {
	var sb strings.Builder
	for _, v := range s {
		sb.WriteString(v)
	}
	return sb.String()
}
