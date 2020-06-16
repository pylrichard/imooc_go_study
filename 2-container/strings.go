package main

import (
	"fmt"
	"unicode/utf8"
	"strings"
)

func main() {
	// UTF-8编码
	str := "Yes在慕课网学习!"
	fmt.Printf("%s, len: %d", str, len(str))
	fmt.Println()

	for _, b := range []byte(str) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()

	// 打印rune十六进制(int32)
	for i, char := range str {
		fmt.Printf("(%d %X) ", i, char)
	}
	fmt.Println()

	fmt.Println("Rune count:",
		utf8.RuneCountInString(str))

	// 解码byte成一个个rune
	bytes := []byte(str)
	for len(bytes) > 0 {
		char, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", char)
	}
	fmt.Println()

	for i, char := range []rune(str) {
		fmt.Printf("(%d %c) ", i, char)
	}
	fmt.Println()

	// 修改rune
	runes := []rune(str)
	runes[3] = '上'
	fmt.Println(string(runes))
	fmt.Println(runes)

	fmt.Println(strings.Split(str, "在"))
	fmt.Println(strings.Contains(str, "学"))
	fmt.Println(strings.Replace(str, "习", "课", -1))
}
