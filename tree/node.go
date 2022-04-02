package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

func CreateNode(value int) *Node {
	return &Node{Value: value}
}

//pass value
func (node Node) Print() {
	fmt.Print(node.Value, " ")
}

//func (node *Node) Traverse() {
//	if node == nil {
//		return
//	}
//	node.Left.Traverse()
//	node.Print()
//	node.Right.Traverse()
//}

// need *
func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("Setting value to nil node, Ignored.")
		return
	}
	// cant get value
	node.Value = value
}
