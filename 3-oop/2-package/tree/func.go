package tree

import "fmt"

// PrintNodeVal 打印节点值
func (node Node) PrintNodeVal() {
	fmt.Println(node.Val)
}

// SetNodeVal 函数参数为传值，需要改值要传引用即指针
func (node *Node) SetNodeVal(val int) {
	if (node == nil) {
		fmt.Println("nil")
		return
	}
	node.Val = val
}

// InOrder 中序遍历方法
func (node *Node) InOrder() {
	if (node == nil) {
		return
	}
	node.Left.InOrder()
	fmt.Print(node.Val, " ")
	node.Right.InOrder()
}

// CreateNode 没有结构构造函数可用工厂模式替代初始化
func CreateNode(val int) *Node {
	return &Node{Val: val}
}