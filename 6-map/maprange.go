package main

import "fmt"

func main() {
	//Map的遍历操作
	capitals := map[string]string{
		"China":  "Beijing",
		"USA":    "Washington D.C.",
		"France": "Paris",
		"Japan":  "Tokyo",
	}

	fmt.Println(capitals)

	for country, capital := range capitals {
		fmt.Printf("%s的首都是%s\n", country, capital)
	}
}
