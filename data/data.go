package data

import "fmt"

//变量交换 不用借助第三个变量
func Exchange() {
	a := 1
	b := 2
	a, b = b, a
	fmt.Println(a)
	fmt.Println(b)
}

//打印
func Print() {
	i := 10
	f := 3.1415926
	s := "你是不是傻"
	//%d 表示整型
	fmt.Printf("%d \n", i)
	//%T 表示数据的类型
	fmt.Printf("%T \n", i)
	//%f 表示浮点型
	//%.2f 保留小数位数为两位  会对第三位小数进行四舍五入
	fmt.Printf("%f \n", f)
	fmt.Printf("%.2f", f)
	//%s 表示字符串型
	fmt.Printf("%s \n", s)

}
