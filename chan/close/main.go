package main

import (
	"fmt"
	"time"
)

/*
结论：
1、关闭chan会导致正在写的goroutine panic
2、关闭chan会导致正在读的goroutine一直获取零值
*/

func readch(ch chan int) {
	for {
		select {
		case x := <- ch:
			fmt.Println(x)
			time.Sleep(time.Second)
		}
	}
}

func writech(ch chan int) {
	for {
		select {
		case ch <- 100:
			fmt.Println("write success!")
			time.Sleep(time.Second)
		}
	}
}

func main() {
	ch := make(chan int)
	go readch(ch)
	go writech(ch)
	time.Sleep(1)
	close(ch)
	time.Sleep(10 * time.Second)
}
