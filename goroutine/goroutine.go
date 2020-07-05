package goroutine

import (
	"fmt"
	"runtime"
	"time"
)

func GoExit() {
	go func() {
		defer fmt.Println("def a")

		func() {
			defer fmt.Println("def b")
			runtime.Goexit()
			fmt.Println("b")
		}()
		fmt.Println("a")
	}()
	time.Sleep(10 * time.Second)
}
