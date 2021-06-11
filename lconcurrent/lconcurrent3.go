package lconcurrent

import (
	"context"
	"fmt"
	"time"
)

//context 控制

func DoSomeThing(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("ctx done")
			return
		default:
			fmt.Println("DoSomeThing")
			time.Sleep(time.Second)
		}
	}
}

func Test3() {
	ctx, cancel := context.WithCancel(context.Background())

	for i := 0; i < 3; i++ {
		go DoSomeThing(ctx)
	}
	time.Sleep(10 * time.Second)
	cancel()
	time.Sleep(5 * time.Second)
}

func Test4() {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	for i := 0; i < 3; i++ {
		go DoSomeThing(ctx)
	}
	time.Sleep(10 * time.Second)
}
