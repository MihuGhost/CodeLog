package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

// 递归为有序数组
func isValidBST1(root *TreeNode) bool {
	res := make([]int, 0)
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		traversal(node.Left)
		res = append(res, node.Val)
		traversal(node.Right)
	}
	traversal(root)

	for i := 1; i < len(res); i++ {
		if res[i] <= res[i-1] {
			return false
		}
	}
	return true
}

// 递归遍历
// 左子树所有节点值 < 中间值
// 右子树所有节点值 > 中间值
func isValidBST2(root *TreeNode) bool {
	minVal := math.MinInt64
	var isValid func(node *TreeNode) bool
	isValid = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		l := isValid(node.Left)
		if minVal < node.Val {
			minVal = node.Val
		} else {
			return false
		}
		r := isValid(node.Right)
		return l && r
	}
	return isValid(root)
}
