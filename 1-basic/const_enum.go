package main

import "fmt"

func consts() {
	const (
		fileName = "abc.txt"
		a, b     = 3, 4
	)
	var c int
	// const可以作为各种类型使用，不需要强制转换
	c = int(math.Sqrt(a * a + b * b))
	fmt.Println(fileName, c)
}

func enums() {
	// 普通枚举
	const (
		lua = 0
		java = 1
	)
	// 自增值枚举
	const (
		// 自增值
		cpp = iota
		// 1跳过
		_
		python
		golang
		// 4
		javascript
	)

	fmt.Println(cpp, javascript, python, golang)

	const (
		// 自增表达式
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	consts()
	enums()
}