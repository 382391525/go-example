package newAndMake

import "fmt"

func NewAndMake() {
	b := new(bool)
	fmt.Println(b)
	fmt.Println(*b)

	s := make([]int, 10)
	fmt.Println(s)
	p := &s[2]
	*p += 200
	fmt.Println(s)

	s1 := *new([]int)
	fmt.Println(s1)
	s1 = append(s1, 1)
	fmt.Println(s1)
}
