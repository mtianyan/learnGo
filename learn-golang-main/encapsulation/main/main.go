package main

import (
	"GoDemoProj/encapsulation/tree"

	"go.uber.org/zap"
)

func main() {
	//createTreeNode()
	root := NewNode{&tree.Node{Value: 3}}
	root.PrintTreeNodeWithPrefix("init value")
	root.SetNodeValue(0)
	root.PrintTreeNodeWithPrefix("updated value")

	logger, _ := zap.NewProduction()
	logger.Warn("Logger test, warning !!")
}

type NewNode struct {
	*tree.Node
}

func (newNode *NewNode) SetNodeValue(value int) {
	println("lalala, I'm shadowing")
}

type TreeNode2 struct {
	*tree.Node // 内嵌，embedding
}

// tree post order traverse
type myTreeNode struct {
	*tree.Node // 内嵌，embedding
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.Node == nil {
		return
	}
	// 创建local variable 来接新建的myTreeNode
	// 这样在调用postOrder函数的时候会自动取地址
	left := myTreeNode{myNode.Left}
	left.postOrder()

	right := myTreeNode{myNode.Right}
	right.postOrder()

	myNode.PrintTreeNodeWithPrefix("")
}
