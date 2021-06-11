package lconcurrent

import (
	"fmt"
	"math"
	"runtime"
	"sync"
)

func product(wg *sync.WaitGroup, ch chan int, num int) {
	wg.Add(1)
	ch <- num
}

func consumer(wg *sync.WaitGroup, ch chan int) {
	for num := range ch {
		//todo...
		fmt.Println("num :", num, "NumGoroutine :", runtime.NumGoroutine())
		wg.Done()
	}
}

func Test2() {
	//无缓冲的chan
	ch := make(chan int)

	wg := sync.WaitGroup{}
	//启动3个消费者
	for i := 0; i < 2; i++ {
		go consumer(&wg, ch)
	}

	taskCnt := math.MaxInt64

	for i := 0; i < taskCnt; i++ {
		product(&wg, ch, i)
	}
	wg.Wait()
}
