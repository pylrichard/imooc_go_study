package main

import (
	"fmt"
	"strconv"
)

// printArray arr []int代表切片，arr [5]int代表数组
// 数组作为参数传递时会拷贝数组(函数内部使用的是副本)
func printArray(arr [5]int) {
	// 数组是值类型，此处修改不影响实参
	arr[0] = 100
	for i, val := range arr {
		fmt.Println("arr[" + strconv.Itoa(i) + "]=" + strconv.Itoa(val))
	}
}

// 指针修改数组
func modifyArray(arr *[5]int) {
	arr[0] = 100
}

func main() {
	//使用var定义数组，无需设置初始值
	var arr1 [5]int
	//使用:=定义数组，必须指定设置初始值
	arr2 := [3]int{1, 3, 5}
	//省略个数，有编译器根据值自动设定长度
	arr3 := [...]int{2, 4, 6, 8, 10}
	//二维数组
	var grid [4][5]int
	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)
	fmt.Println()

	printArray(arr3)
	fmt.Println(arr3)
	//和C语言不同，数组名本身不可作为地址，需要加&
	modifyArray(&arr3)
	printArray(arr3)
}
