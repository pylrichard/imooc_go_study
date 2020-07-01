package main

// go fmt自动格式化
// go imports自动格式化导入包
// go get -v github.com/gpmgo/gopm
// gopm get -g -u -v golang.org/x/tools/cmd/goimports
// go install golang.org/x/tools/cmd/goimports
// 构建并安装当前目录下所有可执行文件
// go install ./...
import (
	"go/imooc_go_study/3-oop/4-extend/tree"
	"fmt"
)

// 组合封装结构对该结构进行扩展
type myTreeNode struct {
	node *tree.Node;
}

// 扩展封装结构的方法
func (myNode *myTreeNode) postOreder() {
	// 子树有可能为空
	if myNode == nil || myNode.node == nil {
		return
	}
	// 本方法需要指针接收者，单独提出来作为变量，可以获取变量地址
	left := myTreeNode{myNode.node.Left}
	left.postOreder()
	right := myTreeNode{myNode.node.Right}
	right.postOreder()
	fmt.Print(myNode.node.Val, " ")
}

func main() {
	// 创建结构初始化
	var root tree.Node
	left := tree.Node{Val: 1}
	right := tree.Node{2, nil, nil}
	root.Left = &left
	root.Right = &right
	fmt.Println(root)

	//  创建结构切片初始化
	nodes := []tree.Node{
		{Val: 1},
		{Val: 2},
		{3, nil, nil},
	}
	fmt.Println(nodes)

	// 使用工厂函数替代
	node1 := tree.CreateNode(3)
	node2 := tree.CreateNode(4)
	left.Right = node1
	right.Left = node2
	fmt.Println(*node1, node2)

	// 使用结构方法
	root.InOrder()
	fmt.Println()
	myRoot := myTreeNode{&root}
	myRoot.postOreder()
}
