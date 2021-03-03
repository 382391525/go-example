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
