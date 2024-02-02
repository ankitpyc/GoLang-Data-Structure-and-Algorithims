// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type Node struct {
	val   int
	left  *Node
	right *Node
}

func preOrderTraverse(node *Node) {
	if node == nil {
		return
	}
	fmt.Print((*(node)).val)
	preOrderTraverse(node.left)
	preOrderTraverse(node.right)
}

func InOrderTraversal(node *Node) {
	if node == nil {
		return
	}
	InOrderTraversal(node.left)
	fmt.Print((*(node)).val)
	InOrderTraversal(node.right)
	fmt.Println("----------")

}

func main() {
	root := Node{5, nil, nil}
	lnode := Node{1, nil, nil}
	rnode := Node{2, nil, nil}
	root.left = &lnode
	root.right = &rnode
	root.left.left = &Node{7, nil, nil}
	preOrderTraverse(&root)
	InOrderTraversal(&root)
}
