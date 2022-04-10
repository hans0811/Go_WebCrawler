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

// Try overwriting
func (myNode *myTreeNode) Traverse() {
	fmt.Println("This method is shadowed")

}

func main() {
	//var root tree.Node
	root := myTreeNode{&tree.Node{Value: 3}}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)

	root.Print()
	root.Right.Left.SetValue(4)
	root.Right.Left.Print()
	fmt.Println(root)

	fmt.Println()
	fmt.Println("In-order traversal: ")
	root.Traverse()
	fmt.Println("root.Node.Traverse(): ")
	root.Node.Traverse()

	fmt.Println()
	fmt.Println("post-order ")
	root.postOrder()
	fmt.Println()

	nodeCount := 0
	root.TraversFunc(func(node *tree.Node) {
		nodeCount++
	})

	fmt.Println(nodeCount)
	//var baseRoot *tree.Node
	//cannot use &root (type *myTreeNode) as type *tree.Node in assignment
	//baseRoot := &root

	//root.Print()
	//root.SetValue(100)
	//
	//pRoot := &root
	//pRoot.Print()
	//pRoot.SetValue(200)
	//pRoot.Print()

	//var nRoot *tree.Node
	//nRoot.SetValue(200)
	//nRoot = &root
	//nRoot.SetValue(300)
	//nRoot.Print()
	//
	//nodes := []tree.Node{
	//	{Value: 3},
	//	{},
	//	{6, nil, &root},
	//}
	//fmt.Println(nodes)

	fmt.Println("Node count: ", nodeCount)

	c := root.TraverseWithChannel()
	maxNode := 0
	for node := range c {
		if node.Value > maxNode {
			maxNode = node.Value
		}
	}

	fmt.Println("Max node value:", maxNode)
}
