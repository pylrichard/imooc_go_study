// 定义别名扩展结构
package main

import (
	"go/imooc_go_study/3-oop/3-extend/queue"
	"fmt"
)

func main() {
	q := queue.MyQueue{1}
	fmt.Println(q)
	q.Push(2)
	q.Push(3)
	fmt.Println(q)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
}
