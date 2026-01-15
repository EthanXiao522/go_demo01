package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var metux sync.Mutex
var rw sync.RWMutex

var count int = 0

func test() {
	metux.Lock()
	count++
	fmt.Println("当前count=", count)
	time.Sleep(50 * time.Millisecond)
	metux.Unlock()
	wg.Done()
}

var m = make(map[int]int, 0)

func test2(num int) {
	defer wg.Done()
	metux.Lock()
	defer metux.Unlock()

	sum := 1
	for i := 1; i <= num; i++ {
		sum *= i
	}
	m[num] = sum

	fmt.Printf("m[%d]=%d\n", num, m[num])
	time.Sleep(time.Millisecond)
}

func main() {
	wg.Add(20)
	for i := 1; i <= 20; i++ {
		go test2(i)
	}
	wg.Wait()
}
