package main

import (
	"fmt"
	"runtime"
	"time"
	"unsafe"
)

func test1() {
	a := make([]*int, 1e9)

	for i := 0; i < 10; i++ {
		start := time.Now()
		runtime.GC()
		fmt.Printf("GC took %s\n", time.Since(start))
	}

	runtime.KeepAlive(a)
}

func test2() {
	a := make([]int, 1e9)

	for i := 0; i < 10; i++ {
		start := time.Now()
		runtime.GC()
		fmt.Printf("GC took %s\n", time.Since(start))
	}

	runtime.KeepAlive(a)
}

type Node struct {
	uint64
	a int32
}
func main() {
	// test1()
	// test2()
	fmt.Println(unsafe.Sizeof(Node{}))
}
