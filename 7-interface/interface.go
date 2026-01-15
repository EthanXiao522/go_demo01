package main

import "fmt"

type Usber interface {
	Start()
	Stop()
}

// 规则:如果接口里面有方法的话,必须要通过结构提或者自定义类型实现这个接口里面的所有方法
type Phone struct {
	Name string
}

// 1.手机要实现usb接口的话必须实现usb接口中的所有方法
func (p Phone) Start() {
	fmt.Println(p.Name + " phone start...")
}

func (p Phone) Stop() {
	fmt.Println(p.Name, "phone stop...")
}

// 2.定义一个Camera结构体
type Camera struct {
	Name string
}

// 3.Camera结构体实现Usber接口中的所有方法
func (c Camera) Start() {
	fmt.Println(c.Name, "camera start...")
}

func (c Camera) Stop() {
	fmt.Println(c.Name, "camera stop...")
}

func (c Camera) run() {
	fmt.Println(c.Name, "相机运行...")
}

// 4.定义computer结构体,可以接收任何实现了Usber接口的类型
type Computer struct {
	Name string
}

func (c Computer) work(usb Usber) {
	fmt.Println(c.Name, "computer work start...")
	usb.Start()
	usb.Stop()
}

func main() {
	p := Phone{
		Name: "Apple",
	}
	p.Start()
	p.Stop()

	//golang中的接口就是一个数据类型
	var p1 = Usber(Phone{
		Name: "Samsung",
	})
	p1.Start()
	p1.Stop()

	//定义p2是接口类型,同时赋值为Phone对象(Phone实现了Usber接口)
	var p2 Usber = p

	p2.Start()
	p2.Stop()

	fmt.Println("=========")

	//表示相机实现了Usber接口
	var c Usber = Camera{
		Name: "Nikon",
	}
	c.Start()
	c.Stop()
	// c.run() //错误,接口类型变量不能调用实现类的特有方法
	// 解决方法:类型断言
	if camera, ok := c.(Camera); ok {
		camera.run()
	}

	fmt.Println("=========")
	com := Computer{
		Name: "HUAWEI",
	}
	com.work(p)
	com.work(c)
}
