package main

import (
	"fmt"
	"math"
)

func triangle(a int, b int) int {
	// 强制类型转换
	c := int(math.Sqrt(float64(a * a + b * b)))
	fmt.Println(c)

	return c
}