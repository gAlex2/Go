package main

import "fmt"

func main() {
	array := []int{9, 80, 29, 19, 6, 48, 2}
	albero := arr2tree(array, 0)
	stampaAlberoASommario(albero, 0)
}

//////////////////////////////////////////////////////////

type bitreeNode struct {
	left  *bitreeNode
	right *bitreeNode
	val   int
}

type bitree struct {
	root *bitreeNode
}

func newNode(val int) *bitreeNode {
	return &bitreeNode{nil, nil, val}
}

//////////////////////////////////////////////////////////

func arr2tree(a []int, i int) (root *bitreeNode) {
	if i >= len(a) {
		return nil
	}

	root = newNode(a[i])
	root.left = arr2tree(a, i*2+1)
	root.right = arr2tree(a, i*2+2)
	return root
}

//////////////////////////////////////////////////////////

func preorder(node *bitreeNode) {
	if node == nil {
		return
	}
	fmt.Println(node.val)
	preorder(node.left)
	preorder(node.right)
}

func inorder(node *bitreeNode) {
	if node == nil {
		return
	}
	inorder(node.left)
	fmt.Println(node.val)
	inorder(node.right)
}

func postorder(node *bitreeNode) {
	if node == nil {
		return
	}
	postorder(node.left)
	postorder(node.right)
	fmt.Println(node.val)
}

//////////////////////////////////////////////////////////

func stampaAlberoASommario(node *bitreeNode, spaces int) {
	if node == nil {
		printSpaces(spaces)
		fmt.Println("*")
		return
	}
	printSpaces(spaces)
	fmt.Print("*", node.val, "\n")
	if node.left == nil && node.right == nil {
		return
	}
	stampaAlberoASommario(node.left, spaces+1)
	stampaAlberoASommario(node.right, spaces+1)
}

func printSpaces(spaces int) {
	for i := 0; i < spaces; i++ {
		fmt.Print(" ")
	}
}

//////////////////////////////////////////////////////////

func stampaAlbero(node *bitreeNode) {
	if node == nil {
		return
	}
	fmt.Print(node.val, "")
	if node.left == nil && node.right == nil {
		return
	}
	fmt.Print(" [")
	if node.left != nil {
		stampaAlbero(node.left)
	} else {
		fmt.Print("-")
	}
	fmt.Print(", ")
	if node.right != nil {
		stampaAlbero(node.right)
	} else {
		fmt.Print("-")
	}
	fmt.Print("]")
}

//////////////////////////////////////////////////////////
