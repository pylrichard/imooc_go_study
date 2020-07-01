package queue

// MyQueue 作为切片[]int的别名
type MyQueue []int

// Push 从队尾添加
func (q *MyQueue) Push(n int) {
	*q = append(*q, n)
}

// Pop 从队首弹出
func (q *MyQueue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]

	return head
}

// IsEmpty 判断是否为空
func (q *MyQueue) IsEmpty() bool {
	return len(*q) == 0
}