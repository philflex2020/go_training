package main

import "fmt"

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Value int
}

func InsertNodeToTree(tree *TreeNode, node *TreeNode) {
	if tree == nil {
		panic("cannot insert into nil root")
	}

	if node.Value > tree.Value {
		if tree.Right == nil {
			tree.Right = node
		} else {
			InsertNodeToTree(tree.Right, node)
		}
	}
	if node.Value < tree.Value {
		if tree.Left == nil {
			tree.Left = node
		} else {
			InsertNodeToTree(tree.Left, node)
		}
	}
}

func InitTree(values ...int) (root *TreeNode) {
	rootNode := TreeNode{Value: values[0]}
	for _, value := range values {
		node := TreeNode{Value: value}
		InsertNodeToTree(&rootNode, &node)
	}
	return &rootNode
}

func main() {
	treeNode := InitTree(5, 4, 6, 8, 9, 7, 1, 3, 2)
	fmt.Println(treeNode)
}
