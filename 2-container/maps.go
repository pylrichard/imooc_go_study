package main

import "fmt"

func main() {
	/*
		创建map
	 */
	fmt.Println("--Create map")
	map1 := map[string]string{
		"course":  "golang",
		"site":    "imooc",
	}
	// map2为empty map
	map2 := make(map[string]int)
	// map3为nil，nil不同于Java的null，可参与运算
	var map3 map[string]int
	fmt.Println(map1, map2, map3)

	/*
		遍历map
	 */
	fmt.Println("--Traverse map")
	for k, v := range map1 {
		fmt.Println(k, ":", v, "...")
	}

	/*
		获取map值，并判断是否存在
	 */
	fmt.Println("--Get values")
	site := map1["site"]
	fmt.Println(`map1["site"] =`, site)
	site, exist := map1["sit"]
	if exist {
		fmt.Println(site)
	} else {
		fmt.Println("key 'sit' does not exist")
	}

	/*
		删除map元素
	 */
	fmt.Println("--Delete values")
	course, exist := map1["course"]
	fmt.Printf("map1[%q] before delete: %q, %v\n",
		"course", course, exist)

	delete(map1, "course")
	course, exist = map1["course"]
	fmt.Printf("map1[%q] after delete: %q, %v\n",
		"course", course, exist)
}
