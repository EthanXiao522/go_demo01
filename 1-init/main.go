package main

import (
	"fmt"
	"go_demo01/1-init/lib1"
	"go_demo01/1-init/lib2"

	"github.com/shopspring/decimal"
	"github.com/tidwall/gjson"
)

func init() {
	fmt.Println("main---init--")
}

func main() {
	fmt.Println("main---main--")
	lib1.Lib1Test()
	lib2.Lib2Test()
	fmt.Println("-------------------")

	n, err := decimal.NewFromString("-123.4567")
	if err != nil {
		panic(err)
	}
	fmt.Println("n=", n) // output: "-123.4567"

	const json = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`
	value := gjson.Get(json, "name.last")
	fmt.Println(value)
}
