package lruntime

import (
	"fmt"
	"runtime"
)

func Test() {
	go func(s string) {
		for i := 0; i < 2; i++ {
			fmt.Println(s)
		}
	}("world")
	for i := 0; i < 2; i++ {
		//出让当前cpu的时间片
		runtime.Gosched()
		fmt.Println("hello")
	}
}

func run(s string) {
	for i := 0; i < 10; i++ {
		fmt.Println(s, ":", i)
	}
}

func GoMaxProcess() {
	runtime.GOMAXPROCS(1)
	go run("a")
	go run("b")
	for  {

	}
}
