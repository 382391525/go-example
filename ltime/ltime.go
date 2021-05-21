package ltime

import (
	"fmt"
	"time"
)

func Unix() {
	fmt.Println(time.Now().Unix())
}

func Timer() {
	timer := time.NewTimer(2 * time.Second)
	<-timer.C
	fmt.Println("timer")
}

func Ticker() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	done := make(chan bool)
	go func() {
		time.Sleep(time.Second * 10)
		done <- true
	}()
	for {
		select {
		case <-done:
			fmt.Println("done !")
			return
		case t := <-ticker.C:
			fmt.Println("ticker:", t.Format("2006-01-02 15:04:05"))
		}
	}
}
func Ticker2() {
	ticker := time.NewTicker(time.Second * 1)
	glb := 1
	go func() {
		for {
			<-ticker.C
			fmt.Println("子:", glb)
			glb++
		}
	}()

	for {
		<-ticker.C
		fmt.Println("父:", glb)
		glb++
	}
}
