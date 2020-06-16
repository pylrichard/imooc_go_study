package main

import "fmt"

func modifySlice(slice []int) {
	slice[0] = 100
}

func printSlice(slice []int) {
	fmt.Printf("len=%d, cap=%d\n",
		len(slice), cap(slice))
	for i, val := range slice {
		fmt.Println("slice[",i,"] = ", val)
	}
	fmt.Println()
}

func createSlice() {
	fmt.Println("--Create slice")
	array := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	// 切片左闭右开
	slice1 := array[2:6]
	fmt.Println("array[2:6]=", slice1)
	slice2 := array[2:]
	fmt.Println("array[2:]=", slice2)
	slice3 := array[:6]
	fmt.Println("array[:6]=", slice3)
	slice4 := array[:]
	fmt.Println("array[:]=", slice4)
	fmt.Println()

	fmt.Println("--After modifySlice(array[2:6])")
	modifySlice(slice1)
	fmt.Println(array)
	fmt.Println("--After modifySlice(array[:])")
	modifySlice(slice4)
	fmt.Println(array)

	// 多次进行切片操作
	modifySlice(array[:][3:])
	fmt.Println(array)
	fmt.Println()

	printSlice(array[:])

	// 切片的零值是nil
	var slice5 []int

	for i := 0; i < 6; i++ {
		printSlice(slice5)
		slice5 = append(slice5, 2 * i + 1)
	}
	fmt.Println(slice5)
	
	slice6 := []int{2, 4, 6, 8}
	printSlice(slice6)
	slice7 := make([]int, 16)
	printSlice(slice7)
	slice8 := make([]int, 10, 32)
	printSlice(slice8)
}

func copySlice() {
	fmt.Println("--Copy slice")
	slice1 := []int{2, 4, 6, 8}
	// 需要预先分配空间
	slice2 := make([]int, len(slice1))
	copy(slice2, slice1)
	fmt.Println(slice1)
	printSlice(slice1)
	fmt.Println(slice2)
	printSlice(slice2)
}

func appendToSlice() {
	fmt.Println("--Append elements to slice")
	array := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	slice1 := array[2:6]
	slice2 := slice1[3:5]
	// 因为参数是传递，必须接收append()返回值
	slice3 := append(slice2, 10)
	slice4 := append(slice3, 11)
	// slice3/slice4底层数组不再是array，是一个拷贝了array元素的新数组
	fmt.Println("slice3, slice4 = ", slice3, slice4)
}

func deleteFromSlice() {
	fmt.Println("--Delete elements from slice")
	slice1 := []int{2, 4, 6, 8}
	// 删除slice1[1]即在slice1[0:1]添加slice1[2:]以及可变元素...
	slice2 := append(slice1[0:1], slice1[2:]...)
	printSlice(slice2)

	fmt.Println("--Popping from front")
	front := slice2[0]
	slice2 = slice2[1:]
	fmt.Println(front)
	printSlice(slice2)
	
	fmt.Println("--Popping from back")
	tail := slice2[len(slice2)-1]
	slice2 = slice2[:len(slice2)-1]
	fmt.Println(tail)
	printSlice(slice2)
}

func extendSlice() {
	// 切片可以扩展访问底层数组，切片 = 指针ptr + 长度len + 容量cap
	// 但只能向后扩展不能向前扩展，slice1[i]不可以超越len(slice1)，向后扩展不可以超越cap(slice1)
	fmt.Println("--Extend slice")
	array := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	slice1 := array[2:6]
	fmt.Println("array[2:6]=", slice1)
	fmt.Println("len(array[2:6])=", len(slice1))
	fmt.Println("cap(array[2:6])=", cap(slice1))
	fmt.Println()
	// 切片值为[slice1[3], slice1[4]]
	slice1 = slice1[3:5]
	fmt.Println("array[2:6][3:5]=", slice1)
	fmt.Println("len(array[2:6][3:5])=", len(slice1))
	fmt.Println("cap(array[2:6][3:5])=", cap(slice1))
}

func reslice() {
	fmt.Println("--Reslice")
	array := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	slice1 := array[2:6]
	slice2 := array[2:]
	fmt.Println(slice1)
	slice1 = slice1[:5]
	fmt.Println(slice1)
	slice2 = slice2[2:]
	fmt.Println(slice2)
}

func main() {
	createSlice()
	reslice()
	extendSlice()
	copySlice()
	appendToSlice()
	deleteFromSlice()
}