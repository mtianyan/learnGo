package tree

import "fmt"

// Node
// 首字母大写：包外可见
// 首字母小写：包外不可见
type Node struct {
	Value       int
	Left, Right *Node
}

func (node *Node) SetNodeValue(value int) {
	node.Value = value
}

func (node Node) PrintTreeNodeWithPrefix(prefix string) {
	fmt.Printf("%v, Value = %v \n", prefix, node.Value)
}

func NodeFactory(value int) *Node {
	return &Node{Value: value} // 局部创建，GC 会控制变量回收，类似java
}

func CreateTreeNode() {
	root := Node{Value: 3}
	root.Left = &Node{}
	root.Right = &Node{5, nil, nil}
	root.Right.Left = new(Node) // new() 创建一个新的TreeNode 并返回指针

	nodes := []Node{
		{Value: 3},
		{Value: 5},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes)
}
