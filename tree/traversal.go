package tree

import "fmt"

func (node *Node) Traverse() {

	node.TraversFunc(func(n *Node) {
		n.Print()
	})
	fmt.Println()
}

func (node *Node) TraversFunc(
	f func(*Node)) {
	if node == nil {
		return
	}

	node.Left.TraversFunc(f)
	f(node)
	node.Right.TraversFunc(f)

}

func (node *Node) TraverseWithChannel() chan *Node {

	out := make(chan *Node)
	go func() {
		node.TraversFunc(func(node *Node) {
			out <- node
		})
		close(out)
	}()

	return out
}
