package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	c := make(chan bool, 100)
	t := time.Tick(time.Second * 1) // 一秒的心跳

	go func() {
		for {
			select {
			case <-t:
				fmt.Printf("NumGoroutine: %d\n", runtime.NumGoroutine()) // 一秒打印当前goroutine运行个数
			}
		}
	}()

	for i := 0; i < 1000000; i++ {
		c <- true
		go worker(i, c)
	}

	fmt.Println("done")
}

func worker(i int, c chan bool) {
	time.Sleep(time.Millisecond * 100) // 模拟其他代码运行消耗时间
	<-c
}
