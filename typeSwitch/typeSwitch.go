package typeSwitch

import "fmt"

func TypeSwitch()  {
	var i interface{} = "fzf"

	str := i.(string)
	fmt.Println(str)
	s,ok := i.(string)
	fmt.Println(s,ok)
}
