package cons

import "fmt"

const (
	i1 = 10
	i2
	i3 = 1
	i4 = iota
	i5
	i6
)

const (
	ii1, ii2 = iota + 1, iota + 2
	ii3, ii4
	ii5, ii6
	ii7, ii8 = iota * 2, iota * 3
	ii9, ii10
)

func Test() {
	fmt.Println(i1)
	fmt.Println(i2)
	fmt.Println(i3)
	fmt.Println(i4)
	fmt.Println(i5)
	fmt.Println(i6)
	//fmt.Println(ii1, ii2)
	//fmt.Println(ii3, ii4)
	//fmt.Println(ii5, ii6)
	//fmt.Println(ii7, ii8)
	//fmt.Println(ii9, ii10)
}
