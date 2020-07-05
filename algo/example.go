package algo

import "fmt"

//水仙花数
func Sxh() {
	for i := 100; i < +999; i++ {
		a := i / 100
		b := i / 10 % 10
		c := i % 10

		if a*a*a+b*b*b+c*c*c == i {
			fmt.Println(i)
		}
	}
}
//冒泡排序
func BubbleSort() {
	s := []int{1, 4, 8, 2, 0, 2, 6, 2, 9}

	for i := 0; i < len(s)-1; i++ {
		for j := 0; j < len(s)-1-i; j++ {
			if s[j]> s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
	fmt.Println(s)
}
