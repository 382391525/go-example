package lgodebug

import (
	"fmt"
	"sync"
)

func Test() {
	wg := sync.WaitGroup{}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(num int) {
			fmt.Println(num)
		}(i)
	}
	wg.Wait()
}
