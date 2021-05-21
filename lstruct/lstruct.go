package lstruct

import "fmt"

type UserInfo struct {
	Name   string
	Age    int
	Height int
}

func Test() {
	u := UserInfo{
		Name:   "fitz",
		Age:    18,
		Height: 180,
	}
	pu := &u
	pu.Name = "fzf"
	fmt.Println(u)
	fmt.Println(pu)
}
