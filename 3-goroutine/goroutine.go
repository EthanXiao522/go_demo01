package main

import (
	"fmt"
	"runtime"
	"time"
)

func newTask() {
	j := 0
	for {
		j++
		fmt.Printf("newTask ... j=%d\n", j)
		time.Sleep(1 * time.Second)
	}
}

func main() {

	// go newTask() //go关键字开启一个协程

	// i := 0
	// for {
	// 	i++
	// 	fmt.Printf("main goroutine .. i=%d\n", i)
	// 	time.Sleep(1 * time.Second)
	// }

	cpuNum := runtime.NumCPU()
	fmt.Printf("cpuNum=%d\n", cpuNum)

	runtime.GOMAXPROCS(cpuNum - 1) //设置cpu的核心数
	fmt.Printf("cpuNum=%d\n", runtime.NumCPU())
}
