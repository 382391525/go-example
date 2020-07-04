package basics

import "fmt"

//变量交换 不用借助第三个变量
func Exchange() {
	a := 1
	b := 2
	a, b = b, a
	fmt.Println(a)
	fmt.Println(b)
}
