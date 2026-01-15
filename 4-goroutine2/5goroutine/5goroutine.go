package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// 统计1-1000之间素数的个数和使用时间
func main() {

	// wg.Add(1)
	// m01(1, 100000)

	start := time.Now()
	//开启5个协程,分别统计不同范围内素数的个数
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go m01((i-1)*20000+1, i*20000)
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("1-100000之间素数,耗时:%v\n", elapsed)
}

func m01(m int, n int) {
	defer wg.Done()
	fmt.Printf("m01:统计%d到%d之间素数的个数\n", m, n)
	start := time.Now()

	count := 0
	for i := m; i <= n; i++ {
		flag := true //是素数
		for j := 2; j < i; j++ {
			if i%j == 0 { //能被除了1和自身以外的数整除,不是素数
				flag = false
				break
			}
		}
		if flag && i > 1 {
			count++
			// fmt.Printf("素数:%d\n", i)
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("%d-%d之间素数的个数:%d,耗时:%v\n", m, n, count, elapsed)

}
