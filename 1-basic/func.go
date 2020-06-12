package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

// 函数可以有多个返回值，获取后不需要的可以填_
func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _, _ := div(a, b)
		return q, nil
	default:
		// fmt.Errorf()有时比panic()友好
		return 0, fmt.Errorf(
			"unsupported operation: %s", op)
	}
}

func div(a, b int) (q, r int, err error) {
	if (b == 0) {
		return q, r, fmt.Errorf("can`t divide 0")
	}
	
	return a / b, a % b, nil
}

// 函数式编程改造eval()
func apply(opFunc func(int, int) int, a, b int) int {
	// 通过反射获取函数指针
	p := reflect.ValueOf(opFunc).Pointer()
	// 获取函数名称
	opFuncName := runtime.FuncForPC(p).Name()
	fmt.Printf("calling function %s with args "+
		"(%d, %d)\n", opFuncName, a, b)

	return opFunc(a, b)
}

func pow(a int, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

// 可变参数列表
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func main() {
	result, err := eval(3, 4, "x")
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println(result)
	}

	q, r, _ := div(13, 3)
	fmt.Printf("13 div 3 is %d mod %d\n", q, r)

	fmt.Println("pow(3, 4) is:", apply(pow, 3, 4))

	fmt.Println("1 + 2 + ... + 5 =", sum(1, 2, 3, 4, 5))
}