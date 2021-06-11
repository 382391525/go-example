package lconcurrent

import (
	"fmt"
	"math"
	"runtime"
	"sync"
)

func do(wg *sync.WaitGroup, ch chan struct{}) {
	defer wg.Done()
	//todo...
	fmt.Println(runtime.NumGoroutine())
	<-ch
}

func Test() {
	ch := make(chan struct{}, 3)
	wg := sync.WaitGroup{}
	taskCnt := math.MaxInt64
	for i := 0; i < taskCnt; i++ {
		ch <- struct{}{}
		wg.Add(1)
		go do(&wg, ch)
	}
	wg.Wait()
}
