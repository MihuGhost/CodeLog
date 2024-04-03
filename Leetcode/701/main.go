package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

// 递归
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		root = &TreeNode{Val: val}
		return root
	}
	if root.Val < val {
		root = insertIntoBST(root.Right, val)
	} else {
		root = insertIntoBST(root.Left, val)
	}
	return root
}

// 迭代
func insertIntoBST1(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	node := root
	for node != nil {
		if val > node.Val {
			if node.Right != nil {
				node = node.Right
			} else {
				node.Right = &TreeNode{Val: val}
				break
			}
		} else {
			if node.Left != nil {
				node = node.Left
			} else {
				node.Left = &TreeNode{Val: val}
				break
			}
		}
	}
	return root
}
