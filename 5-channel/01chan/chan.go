package main

import "fmt"

func main() {
	//chan 的基本使用

	//1. 创建一个无缓冲区的chan, int类型
	chan1 := make(chan int)

	defer close(chan1)

	chan2 := make(chan int, 3) //有缓冲区的chan

	chan2 <- 101
	chan2 <- 102

	num01 := <-chan2
	num02 := <-chan2
	// num03 := <-chan2//fatal error: all goroutines are asleep - deadlock! 管道阻塞-死锁
	// fmt.Println(num03)

	// fmt.Println(num01)
	fmt.Println(num01, num02) //101 102

	//查看chan的长度和容量(管道是引用数据类型)
	fmt.Println("chan1 len=", len(chan1), "cap=", cap(chan1))                    //chan1 len= 0 cap= 0
	fmt.Printf("chan2值:%v chan2 len=%d cap=%d\n", chan2, len(chan2), cap(chan2)) //chan2值:0xc000026100 chan2 len=0 cap=3

	chan2 <- 201
	fmt.Printf("chan2值:%v chan2 len=%d cap=%d\n", chan2, len(chan2), cap(chan2)) //chan2值:0xc000026100 chan2 len=1 cap=3

	//close chan
	close(chan2)

	// chan3 := make(<-chan int) //只读chan

	// chan4 := make(chan<- int) //只写chan

	fmt.Println("-----test01-----")
	test01()
}

// test01 演示for range 读取chan(for range 读取chan时,如果chan没有关闭,会导致死锁)
func test01() {
	chan01 := make(chan int, 10)
	for i := 1; i < 11; i++ {
		chan01 <- i
	}
	close(chan01)

	for v := range chan01 { //fatal error: all goroutines are asleep - deadlock! for range 读取chan时,如果chan没有关闭,会导致死锁
		fmt.Println("v=", v)
	}
}
