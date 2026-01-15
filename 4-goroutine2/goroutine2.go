package main

import (
	"fmt"
	"sync"
	"time"
)

// 开启多个协程,每个协程打印自己的编号和循环次数,主线程要等待所有协程结束后再退出
var wg sync.WaitGroup

func test(num int) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("协程编号:%d,循环次数:%d\n", num, i)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go test(i) //开启多个协程
	}
	wg.Wait()
	fmt.Println("main线程等待所有协程结束...")

}
