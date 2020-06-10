package main

import "fmt"

// 包级变量定义不能使用:=
var (
	str     = "hello go!"
	packInt = 1
)

func variableZeroValue() {
	// 默认值为0
	var a int
	// 默认值为空字符串
	var s string
	// %q表示双引号字符串，可用于标识空字符串
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	// Go语言要求所有变量都必须使用，否则报错
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Printf("%d %d %q\n", a, b, s)
}

func variableTypeDeduction() {
	// 类型自动推断
	var a, b, s = 3, 4, "def"
	fmt.Printf("%d %d %q\n", a, b, s)
}

func variableLocal() {
	// 定义变量
	a, b, s := 3, 4, "def"
	// 赋值
	b = 5
	fmt.Printf("%d %d %q\n", a, b, s)
}

func main() {
	fmt.Printf("%d\n", packInt)
	fmt.Print(str + "\n")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableLocal()
}