package time

import (
	"fmt"
	"github.com/golang-module/carbon"
)

func LongLi() {
	//c := carbon.SetLocale("zh-CN").SetTimezone(carbon.PRC).Now()
	//d := carbon.Parse("2020-08-05 13:14:15").Constellation() // Leo
	//fmt.Println(d)
	//fmt.Println(c.ToDateTimeString())
	//fmt.Println(c.Constellation())
	//fmt.Println(c.ToAnimalYear())
	//fmt.Println(c.ToLunarYear())
	//fmt.Println("test")
	//fmt.Println(carbon.Parse("2020-08-05 13:14:15").Constellation())
	fmt.Println(carbon.Parse("2020-08-05 13:14:15").SetLocale("zh-CN").Constellation())
}
