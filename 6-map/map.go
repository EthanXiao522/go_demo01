package main

import (
	"encoding/json"
	"fmt"
)

func main2() {
	//1.map 声明和初始化
	var map1 map[int]string
	map1 = make(map[int]string)
	map1[1] = "golang"
	map1[2] = "java"
	map1[3] = "python"
	fmt.Println("map1=", map1)

	//2.声明时初始化 空map
	var map2 map[int]string = map[int]string{}
	map2[1] = "c"
	fmt.Println("map2=", map2)

	//3.使用make函数
	var map3 = make(map[int]string)
	map3[1] = "c++"
	fmt.Println("map3=", map3)

	//4.使用map字面量
	var map4 = map[int]string{
		1: "go",
		2: "java",
		3: "python",
	}
	fmt.Println("map4=", map4)

	//5.短变量声明
	map5 := map[int]string{
		1: "c",
		2: "c++",
	}
	fmt.Println("map5=", map5)

	//6.map元素操作
	map6 := make(map[string]string)
	map6["name"] = "golang"
	map6["type"] = "programming language"
	fmt.Println("map6=", map6)
	//6.1根据key取值
	name := map6["name"]
	fmt.Println("name=", name)
	//6.2删除元素
	delete(map6, "type")
	fmt.Println("map6=", map6)
	//6.3判断key是否存在
	value, ok := map6["type"]
	if ok {
		fmt.Println("存在，value=", value)
	} else {
		fmt.Println("不存在")
	}

	if v, ok := map6["name"]; ok {
		fmt.Println("存在，v=", v)
	} else {
		fmt.Println("不存在")
	}

	//7.遍历map
	map7 := map[int]string{
		1: "golang",
		2: "java",
		3: "python",
	}
	for k, v := range map7 {
		fmt.Printf("k=%d v=%s\n", k, v)
	}

	//8.计算map长度
	fmt.Println("map7长度=", len(map7))

	//9.嵌套map
	var map8 = make(map[string]map[string]string)
	map8["name"] = make(map[string]string)
	map8["name"]["first"] = "Janet"
	map8["name"]["last"] = "Prichard"
	fmt.Println("map8=", map8)

	//10.创建复杂结构的map
	response := make(map[string]interface{})
	response["code"] = 200
	response["message"] = "ok"
	response["data"] = map[string]interface{}{
		"user": map[string]interface{}{
			"id":    1,
			"name":  "Janet",
			"email": "janet@example.com",
		},
		"preferences": []string{"golang", "python", "java"},
		"metadata": map[string]interface{}{
			"created_at": "2024-06-01",
			"updated_at": "2024-06-15",
		},
	}
	fmt.Println("原始的map结构：")
	printMap(response, "")

	//序列化为JSON(美化输出)
	jsonData, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		fmt.Println("JSON序列化失败:", err)
		return
	}
	fmt.Println("序列化后的JSON结构：")
	fmt.Println(string(jsonData))

	//反序列化回map
	var parsedMap map[string]interface{}
	if err := json.Unmarshal(jsonData, &parsedMap); err != nil {
		fmt.Println("JSON反序列化失败:", err)
		return
	}
	fmt.Println("反序列化后的map结构：")
	printMap(parsedMap, "")

}

// 辅助函数：递归打印map
func printMap(m map[string]interface{}, indent string) {
	for k, v := range m {
		switch val := v.(type) {
		case map[string]interface{}:
			fmt.Printf("%s%s:\n", indent, k)
			printMap(val, indent+"  ")
		case map[string]string:
			fmt.Printf("%s%s:\n", indent, k)
			for kk, vv := range val {
				fmt.Printf("%s  %s: %v\n", indent, kk, vv)
			}
		case []interface{}:
			fmt.Printf("%s%s: [\n", indent, k)
			for _, elem := range val {
				switch e := elem.(type) {
				case map[string]interface{}:
					printMap(e, indent+"    ")
				default:
					fmt.Printf("%s    %v\n", indent, e)
				}
			}
			fmt.Printf("%s]\n", indent)
		case []string:
			fmt.Printf("%s%s: [", indent, k)
			for i, s := range val {
				if i > 0 {
					fmt.Printf(", ")
				}
				fmt.Printf("%q", s)
			}
			fmt.Printf("]\n")
		default:
			fmt.Printf("%s%s: %v\n", indent, k, val)
		}
	}
}
