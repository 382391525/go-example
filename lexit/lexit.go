package lexit

import (
	"fmt"
)

func GoExit() {
	go func() {
		defer fmt.Println("A.defer")

		func() {
			defer fmt.Println("B.defer")
			//runtime.Goexit() // 终止当前 goroutine, import "runtime"
			//os.Exit(1)   //结束当前整个程序
			return  //结束当前函数,并返回指定值
			fmt.Println("B") // 不会执行
		}()

		fmt.Println("A") // 不会执行
	}() //不要忘记()

	//目的不让主goroutine结束
	for  {
		
	}
}
