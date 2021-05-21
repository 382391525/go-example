package ltrace

import (
	"os"
	"runtime/trace"
)

func Test() {
	trace.Start(os.Stderr)
	defer trace.Stop()
	ch := make(chan string)
	go func() {
		ch <- "hello"
	}()
	<-ch
}
