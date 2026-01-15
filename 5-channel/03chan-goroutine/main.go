package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// 统计1-10万 之间素数的个数和使用时间
func main() {
	start := time.Now()
	//存放数字的chan
	numChan := make(chan int, 100)

	//存放素数的chan
	primeChan := make(chan int, 100)

	//存放标识
	exitChan := make(chan bool, 16)

	//开启一个协程,向numChan中放入1-100000的数字
	wg.Add(1)
	go putNum(numChan)

	//开启一个协程,从numChan中取出数据,判断是否为素数,如果是,就放入primeChan中
	for i := 0; i < 16; i++ {
		wg.Add(1)
		go primeNum(numChan, primeChan, exitChan)
	}

	//开启一个协程,等待16个primeNum协程结束后,关闭primeChan
	wg.Add(1)
	go func() {
		for i := 0; i < 16; i++ {
			<-exitChan //等待16个协程结束,chan为空时会阻塞
		}
		close(primeChan)
		wg.Done()
	}()

	//打印素数
	wg.Add(1)
	go func() {
		count := 0
		for range primeChan {
			// println("素数:", v)
			count++
		}
		println("-----素数个数:", count)
		wg.Done()
	}()

	wg.Wait()
	elapsed := time.Since(start)

	fmt.Println("main线程等待所有协程结束...")
	fmt.Println("耗时:", elapsed)
}

func putNum(intChan chan int) {
	for i := 1; i <= 100000; i++ {
		intChan <- i
	}
	close(intChan)
	wg.Done()
}

func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	for num := range intChan {
		if num > 1 {
			isPrime := true
			for i := 2; i*i <= num; i++ { //优化:只需判断到sqrt(num)
				if num%i == 0 {
					isPrime = false
					break
				}
			}
			if isPrime {
				primeChan <- num
			}
		}
	}
	exitChan <- true
	wg.Done()
}
