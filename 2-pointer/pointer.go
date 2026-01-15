package main

import "fmt"

func swap(pa, pb *int) {
	var temp = *pa
	*pa = *pb
	*pb = temp
}
func main() {
	a := 10
	b := 20

	swap(&a, &b)
	fmt.Println("a=", a, " b=", b)

	//二级指针
	// var p *int
	p := &a
	fmt.Println(&a)
	fmt.Println(p)

	var pp **int
	pp = &p
	fmt.Println(&p)
	fmt.Println(pp)
}
