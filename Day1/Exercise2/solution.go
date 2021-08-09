package main

import "fmt"

type Node struct {
	value string
	left *Node
	right *Node
}



func (root *Node) PreOrderTraversal() {
	if root !=nil {
		fmt.Println(root.value)
		root.left.PreOrderTraversal()
		root.right.PreOrderTraversal()
	}

}
func (root *Node) PostOrderTraversal() {
	if root !=nil {
		root.left.PostOrderTraversal()
		root.right.PostOrderTraversal()
		fmt.Println(root.value)
	}

}
func main(){
	root := Node{"-",nil,nil}
	root.left = &Node{"a",nil,nil}

	root.right = &Node{"+",nil,nil}
	root.right.left = &Node{"b",nil,nil}
	root.right.right = &Node{"c",nil,nil}

	fmt.Println("Preorder Traversal")
	root.PreOrderTraversal()
	fmt.Println(root.value)
	fmt.Println("Post order traversal")
	root.PostOrderTraversal()

}