package main

import "fmt"

// treeNode 树节点
type treeNode struct {
	val         int
	left, right *treeNode
}

// 定义结构的方法，(node treeNode)表示this
// 定义和printNode(node treeNode)一致，treeNode值或指针都可以调用该方法(自动识别)
func (node treeNode) printNodeVal() {
	fmt.Println(node.val)
}

// Go语言调用函数传参都是传值，需要改值要传引用即指针，treeNode值或指针都可以进行调用该方法(自动识别)
// 指针作为接收者：在改变成员值、结构比较大、保持一致性时考虑使用
func (node *treeNode) setNodeVal(val int) {
	if (node == nil) {
		fmt.Println("nil")
		// nil可以调用方法，但获取成员会报错
		return
	}
	node.val = val
}

// 中序遍历方法
func (node *treeNode) inOrder() {
	// 判空收敛在被调用方
	if (node == nil) {
		return
	}
	// 遍历左子树
	node.left.inOrder()
	// 访问自身
	fmt.Print(node.val, " ")
	// 遍历右子树
	node.right.inOrder()
}

// 没有结构构造函数，可用工厂模式替代初始化
// 由编译器和runtime决定局部变量分配在栈还是堆，返回局部变量指针提示编译器分配局部变量到堆
func createTreeNode(val int) *treeNode {
	return &treeNode{val: val}
}

func main() {
	/*
		创建结构初始化
	 */
	var root treeNode
	left := treeNode{val: 1}
	right := treeNode{2, nil, nil}
	root.left = &left
	root.right = &right
	fmt.Println(root)

	/*
		创建结构切片初始化
	 */
	nodes := []treeNode{
		{val: 1},
		{},
		{3, nil, &root},
	}
	fmt.Println(nodes)

	// 指针和实例都使用.访问成员
	//root.right.left = new(treeNode)
	/*
		使用工厂函数创建
	 */
	node1 := createTreeNode(3)
	node2 := createTreeNode(4)
	left.right = node1
	right.left = node2
	fmt.Println(*node1, node2)

	/*
		使用结构方法
	 */
	root.inOrder()
	fmt.Println()
	// 自动识别接收者是值还是指针
	(&root).printNodeVal()
	root.printNodeVal()
	root.setNodeVal(0)
	root.printNodeVal()
	// nil也可以调用方法
	var ptr *treeNode
	ptr.setNodeVal(1)
	ptr = &root
	ptr.setNodeVal(2)
}
