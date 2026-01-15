package main

import "fmt"

func main() {
	fmt.Println("你好")
	test()
	fmt.Println("\n---------------------")
	/**
	1.变量命名
	var 变量名 类型（初始化值时可省略 类型）
	变量名 := 值	//短变量声明

	*/
	var name string = "名称"
	fmt.Println(name)

	a := 10
	fmt.Printf("a=%d\n", a)
	fmt.Printf("a的类型：%T\n", a)

	a1, b, c := 1, 2, 3
	fmt.Println(a1, b, c)

	//匿名变量
	// var username, age = getUserInfo()
	// fmt.Println(username, age)

	var username, _ = getUserInfo()
	fmt.Println(username)

	//一次声明多个变量
	var (
		m string
		n int
		l string
	)
	m = "string_m"
	n = 11
	l = "string_l"
	fmt.Println(m, n, l)

	//常量
	const pi = 3.1415
	fmt.Println(pi)
	// pi = 4  //常量值不可改变

	//多个常量一起声明，如果省略后面的值，跟前面一样
	const (
		A = "A"
		B
		C = "C"
	)
	fmt.Println(A, B, C)

	//常量计数器 iota
	const (
		A1 = iota
		B1
		C1
	)
	fmt.Println(A1, B1, C1)

}

// 函数返回多个值
func getUserInfo() (string, int) {
	return "张三", 20
}

func test() {
	fmt.Print("ceshi 测试")
	fmt.Printf("好的")
}
