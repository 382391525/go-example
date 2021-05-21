package t

import (
	"fmt"
	"github.com/golang-module/carbon"
	"log"
)

func Constellation() {
	c := carbon.Now().AddHours(1).SetLocale("zh-CN")
	if c.Error != nil {
		// 错误处理
		log.Fatal(c.Error)
	}
	// 狮子座
	fmt.Println(c.Constellation())
}
func IotaTest() {
	//1.iota 枚举
	//2.每遇到const iota重置为0
	//3.每出现一个iota 值加1
	const (
		c0 = iota //0
		c1 = iota //1  相当于沿用上一个变量的赋值
		c2 = 5    //5
		c3 = iota //3  //这里不是 5+1
	)
	fmt.Println(c0, c1, c2, c3)
	const (
		a1 = 5    //5
		a2 = iota //1
		a3        //2
		a4        //3
	)
	fmt.Println(a1, a2, a3, a4)
}


type People interface {
	Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "你是个大帅比"
	} else {
		talk = "您好"
	}
	return
}

func Test() {
	var peo People = &Student{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
}