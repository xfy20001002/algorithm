package main

import "fmt"

type TreeNode struct {
	Val   byte
	left  *TreeNode
	right *TreeNode
}

//根据前序遍历的顺序创建一颗树
/*
	1
  2   3
 5 6 # 8
*/
//根据字符序列创建二叉树
func createTree(arr []byte, index *int) *TreeNode {
	var root *TreeNode
	if *index == len(arr) {
		return root
	} else if arr[*index] == '#' {
		root = nil
		*index = *index + 1
	} else {
		root = &TreeNode{Val: arr[*index]}
		//fmt.Printf("%c ", root.Val)
		*index = *index + 1
		root.left = createTree(arr, index)
		root.right = createTree(arr, index)
	}
	return root
}

func CreateTree(arr []byte) *TreeNode {
	zero := 0
	return createTree(arr, &zero)
}

func preOrder(root *TreeNode) {
	if root == nil {
		return
	} else {
		fmt.Printf("%c ", root.Val)
		preOrder(root.left)
		preOrder(root.right)
	}
}

func main() {
	arr := []byte{'1', '2', '5', '#', '#', '6', '#', '#', '3', '#', '8'}
	root := CreateTree(arr)
	preOrder(root)
}
