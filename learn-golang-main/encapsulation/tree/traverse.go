package tree

func (node *Node) TraverseInorder() {
	if node == nil {
		return
	}
	node.Left.TraverseInorder()
	println(node.Value)
	node.Right.TraverseInorder()
}

func TraverseInorder(root *Node) {
	if root == nil {
		return
	}
	TraverseInorder(root.Left)
	println(root.Value)
	TraverseInorder(root.Right)
}

func (node *Node) TraverseFunc(f func(*Node)) {
	if node == nil {
		return
	}
	node.Left.TraverseFunc(f)
	f(node) // do sth in function
	node.Right.TraverseFunc(f)
}

func (node *Node) Traverse() {
	node.TraverseFunc(func(n *Node) {
		n.PrintTreeNodeWithPrefix("")
	})
}

func countNode(root *Node) {
	count := 0
	root.TraverseFunc(func(n *Node) {
		count++
	})
	println("count = ", count)
}
