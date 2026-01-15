package main

import "fmt"

// 空接口
func m01(a interface{}) {
	if v, ok := a.(string); ok {
		fmt.Printf("类型:%T，值是:%s\n", v, v)
	} else {
		fmt.Println("类型断言失败")
	}
}

func m02(a interface{}) {
	switch a.(type) {
	case int:
		fmt.Println("类型是int，值是:", a)
	case string:
		fmt.Println("类型是string，值是:", a)
	case bool:
		fmt.Println("类型是bool，值是:", a)
	default:
		fmt.Println("未知类型")
	}
}

func main() {
	var a interface{}
	a = 100
	fmt.Println(a)

	a = "golang"
	fmt.Println(a)

	a = true
	fmt.Println(a)

	var b interface{}
	b = 3.14
	fmt.Println(b)

	m01("hello")
	m01(200)

	m02(100)
	m02("hello")
	m02(true)
	m02(3.14)
}
