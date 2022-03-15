package main

//二叉树的层序遍历
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createTree(arr []byte, index *int) *TreeNode {
	var root *TreeNode
	if *index == len(arr) {
		return root
	} else if arr[*index] == '#' {
		root = nil
		*index = *index + 1
	} else {
		root = &TreeNode{Val: int(arr[*index])}
		//fmt.Printf("%c ", root.Val)
		*index = *index + 1
		root.Left = createTree(arr, index)
		root.Right = createTree(arr, index)
	}
	return root
}

func CreateTree(arr []byte) *TreeNode {
	zero := 0
	return createTree(arr, &zero)
}

/*
 * type TreeNode struct {
 *   Val int
 *   Left *TreeNode
 *   Right *TreeNode
 * }
 */

/**
 *
 * @param root TreeNode类
 * @return int整型二维数组
 */
/*
 * type TreeNode struct {
 *   Val int
 *   Left *TreeNode
 *   Right *TreeNode
 * }
 */

/**
 *
 * @param root TreeNode类
 * @return int整型二维数组
 */
func LevelOrder(root *TreeNode) [][]int {
	// write code here
	//队列
	q := make([]*TreeNode, 0)
	//初始化
	if root == nil {
		return [][]int{}
	}
	res := make([][]int, 0)
	q = append(q, root)
	for len(q) > 0 {
		size := len(q)
		temp := make([]int, 0)
		for size > 0 {

			node := q[0]
			q = q[1:]
			temp = append(temp, node.Val)
			if node.Left != nil || node.Right != nil {
				if node.Left != nil {
					q = append(q, node.Left)
				}
				if node.Right != nil {
					q = append(q, node.Right)
				}
			}
			size--
		}
		res = append(res, temp)
	}
	return res
}

/*
func main() {
	root := CreateTree([]byte{'1', '#', '2'})
	res := LevelOrder(root)

	fmt.Println(res)
}
*/
