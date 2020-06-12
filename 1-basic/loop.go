package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func convertDecToBin(n int) string {
	result := ""
	// for代替while，省略初始条件
	for ; n > 0; n /= 2 {
		num := n % 2
		result = strconv.Itoa(num) + result
	}

	return result
}

func printFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	printFileContents(file)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	// for代替while，省略结束条件和递增条件
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func forever() {
	// for代替while，死循环
	for {
		fmt.Println("abc")
	}
}

func main() {
	fmt.Println("convertDecToBin results:")
	fmt.Println(
		convertDecToBin(5),
		convertDecToBin(13),
		// 最后一行要跟逗号
		convertDecToBin(0),
	)

	fmt.Println("\nfile contents:")
	printFile("loop.go")

	fmt.Println("\nprint a string:")
	s := `abc"d"
	123

	p`
	printFileContents(strings.NewReader(s))

	// forever()
}
