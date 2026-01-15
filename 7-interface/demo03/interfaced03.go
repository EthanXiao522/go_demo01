package main

import "fmt"

type Animaler interface {
	SetName(name string)
	GetName() string
}

type Dog struct {
	name string
}

func (d *Dog) SetName(name string) { //指针接收者实现接口
	d.name = name
}

func (d Dog) GetName() string { //值接收者实现接口
	return d.name
}

type Address struct {
	Name string
	Age  int
}

func main() {
	d1 := &Dog{
		name: "小黄",
	}

	var a Animaler = d1
	fmt.Println("name=", a.GetName())
	a.SetName("旺财")
	fmt.Println("name=", a.GetName())

	fmt.Println("===========")

	//空接口和类型断言的使用细节
	var v = make(map[string]interface{})
	v["name"] = "golang"
	v["age"] = 18
	v["married"] = false
	v["hobby"] = []string{"唱", "跳", "rap", "篮球"}
	v["address"] = Address{
		Name: "张三",
		Age:  10,
	}

	fmt.Println(v["name"])
	fmt.Println(v["age"])
	fmt.Println(v["married"])
	fmt.Println(v["hobby"])
	fmt.Println(v["address"])

	if addr, ok := v["address"].(Address); ok {
		fmt.Println("地址Name:", addr.Name)
	}

	if hobbies, ok := v["hobby"].([]string); ok {
		fmt.Println("hobby的第一个爱好:", hobbies[0])
	}
}
