package channel

import (
	"fmt"
	"time"
)

func Rang() {
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()
	for v := range c {
		fmt.Println(v)
	}
}

func Test() {
	ch1 := make(chan int)
	go func() {
		fmt.Println(<-ch1)
	}()
	ch1 <- 5
	time.Sleep(1 * time.Second)
}

func Chan1() {
	c := make(chan int, 0) //创建无缓冲的通道 c

	//内置函数 len 返回未被读取的缓冲元素数量，cap 返回缓冲区大小
	fmt.Printf("len(c)=%d, cap(c)=%d\n", len(c), cap(c))

	go func() {
		defer fmt.Println("子go程结束")

		for i := 0; i < 3; i++ {
			c <- i
			fmt.Printf("子go程正在运行[%d]: len(c)=%d, cap(c)=%d\n", i, len(c), cap(c))
		}
	}()

	time.Sleep(2 * time.Second) //延时2s

	for i := 0; i < 3; i++ {
		num := <-c //从c中接收数据，并赋值给num
		fmt.Println("num = ", num)
	}

	fmt.Println("main进程结束")
}

func Chan2() {
	c := make(chan int, 3) //带缓冲的通道

	//内置函数 len 返回未被读取的缓冲元素数量， cap 返回缓冲区大小
	fmt.Printf("len(c)=%d, cap(c)=%d\n", len(c), cap(c))

	go func() {
		defer fmt.Println("子go程结束")

		for i := 0; i < 3; i++ {
			c <- i
			fmt.Printf("子go程正在运行[%d]: len(c)=%d, cap(c)=%d\n", i, len(c), cap(c))
		}
	}()

	time.Sleep(2 * time.Second) //延时2s
	for i := 0; i < 3; i++ {
		num := <-c //从c中接收数据，并赋值给num
		fmt.Println("num = ", num)
	}
	fmt.Println("main进程结束")
}

func Practice() {
	c := make(chan int, 1)
	c <- 1
	go func() {
		for {
			glb := <-c
			fmt.Println("子:", glb)
			glb++
			c <- glb
		}
	}()
	for {
		glb := <-c
		fmt.Println("父:", glb)
		glb++
		c <- glb
	}
}
