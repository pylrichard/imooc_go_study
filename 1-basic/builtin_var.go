package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func euler() {
	// 1i表示虚数
	c := 3 + 4i
	fmt.Printf("Abs(3+4i)=%f\n", cmplx.Abs(c))
	// complex64的实部和虚部都是float32
	// complex128的实部和虚部都是float64
	fmt.Println("e^i(Pi)+1=", cmplx.Pow(math.E, 1i * math.Pi) + 1)
	// %.3f保留小数点后3位
	fmt.Printf("e^i(Pi)+1=%.3f\n", cmplx.Exp(1i * math.Pi) + 1)
}

func triangle() {
	var a, b int = 3, 4
	var c int
	// 强制类型转换
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func main() {
	euler();
	triangle();
}