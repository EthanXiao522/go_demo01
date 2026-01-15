package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// 2个协程,一个协程向chan中写数据,另一个协程从chan中读数据
func main() {

	chan01 := make(chan int, 5)

	wg.Add(2)
	go write(chan01)

	go read(chan01)

	wg.Add(1)
	go func() {
		defer func() {
			//捕获panic
			if err := recover(); err != nil {
				fmt.Print("test捕获到panic:", err)
			}
		}()
		defer wg.Done()
		//定义一个map
		var myMap map[int]string
		myMap[0] = "hello" //panic: assignment to entry in nil map
	}()

	wg.Wait()
}

func write(ch chan<- int) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		ch <- i
		println("写入数据:", i)
		time.Sleep(5 * time.Millisecond)
	}
	close(ch)
}

func read(ch <-chan int) {
	defer wg.Done()
	for {
		num, ok := <-ch
		if !ok {
			println("chan已经关闭,读取数据结束")
			break
		}
		println("--读取数据:", num)
		time.Sleep(200 * time.Millisecond)
	}
}
