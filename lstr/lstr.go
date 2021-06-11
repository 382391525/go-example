package lstr

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func Test() {
	a := "go"
	b := "hello"
	//1.strings.Compare 对比两个字符串的大小
	fmt.Println(strings.Compare(a, b))
	fmt.Println(strings.Compare(a, a))
	fmt.Println(strings.Compare(b, a))
	//2.strings.EqualFold 比较两个字符,忽略大小写后是否相等
	fmt.Println(strings.EqualFold("go", "Go"))
	fmt.Println(strings.EqualFold("哈哈", "呵呵"))
	//3. 子串是否在字符串中
	fmt.Println(strings.Contains("abcd", "a"))
	//4. Fields 用一个或多个连续的空格分隔字符串 s，返回子字符串的数组 slice
	fmt.Println(strings.Fields(" a  b   c    d   "))
	//5.
	fmt.Println(strings.FieldsFunc(" a  b   c    d   ", unicode.IsSpace))

	//5.字符串分割
	fmt.Println(strings.Split(" a,b,c,d, ,", ","))
	fmt.Println(strings.SplitAfter(" a,b,c,d, ,", ","))
	//6.字符串重复几次
	fmt.Println(strings.Repeat("apple", 5))
	//字符串替换
	fmt.Println(strings.Map(func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return r + 32
		case r >= 'a' && r <= 'z':
			return r
		case unicode.Is(unicode.Han, r):
			return '\n'
		}
		return -1
	}, "Hello你#￥%……\\n（'World\\n,好Hello^(&(*界gopher..."))
}

func Join(s []string, sep string) string {
	if len(s) == 0 {
		return ""
	}
	if len(s) == 1 {
		return s[0]
	}
	n := len(sep) * (len(s) - 1)
	for i := 0; i < len(s); i++ {
		n += len(s[i])
	}
	b := make([]byte, n)
	bp := copy(b, s[0])
	for _, s := range s[1:] {
		bp += copy(b[bp:], sep)
		bp += copy(b[bp:], s)
	}
	return string(b)
}

func StrconvTest() {
	//1.格式化有符号整型
	fmt.Println(strconv.FormatInt(100, 10))
	fmt.Println(strconv.Atoi("100"))
	fmt.Println(strconv.Itoa(1000))

}
