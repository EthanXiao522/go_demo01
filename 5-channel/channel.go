package channel

import (
	"fmt"
	"time"
)

func main() {
	defer fmt.Println("main goroutine extist")

	/*
		c := make(chan int)

		go func() {
			defer fmt.Println("goroutine extist")
			c <- 666
		}()

		num := <-c
		fmt.Printf("channel value num=%d\n", num)
	*/

	chanNum := make(chan int, 3)

	go func() {
		defer fmt.Println("子go程结束")

		for i := 0; i < 10; i++ {
			chanNum <- i
			fmt.Printf("写入数据：%d, channel len=%d, channel capability=%d\n", i, len(chanNum), cap(chanNum))
			// time.Sleep(1 * time.Second)
		}
	}()

	for i := 0; i < 10; i++ {
		fmt.Printf("查询数据：channel len=%d, channel capability=%d\n", len(chanNum), cap(chanNum))
		j := <-chanNum
		fmt.Printf("接受数据：%d，channel len=%d, channel capability=%d\n", j, len(chanNum), cap(chanNum))
		time.Sleep(1 * time.Second)
	}
}
