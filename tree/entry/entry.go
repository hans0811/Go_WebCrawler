package main

import (
	"001_go_env/tree"
	"fmt"
)

type myTreeNode struct {
	//node *tree.Node
	*tree.Node // embedding
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.Node == nil {
		return
	}

	left := myTreeNode{myNode.Left}
	left.postOrder()
	right := myTreeNode{myNode.Right}
	right.postOrder()
	myNode.Print()

}

func main() {
	//var root tree.Node
	root := tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)

	root.Print()
	root.Right.Left.SetValue(4)
	root.Right.Left.Print()
	fmt.Println(root)

	fmt.Println()
	root.Traverse()

	myRoot := myTreeNode{&root}
	myRoot.postOrder()
	fmt.Println()

	root.Print()
	root.SetValue(100)

	pRoot := &root
	pRoot.Print()
	pRoot.SetValue(200)
	pRoot.Print()

	var nRoot *tree.Node
	nRoot.SetValue(200)
	nRoot = &root
	nRoot.SetValue(300)
	nRoot.Print()

	nodes := []tree.Node{
		{Value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes)
}
