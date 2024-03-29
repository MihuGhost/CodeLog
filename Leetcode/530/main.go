package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func getMinimumDifference1(root *TreeNode) int {
	var res []int
	traversal1(root, &res)
	num := math.MaxInt64
	for i := 1; i < len(res); i++ {
		if res[i]-res[i-1] < num {
			num = res[i] - res[i-1]
		}
	}
	return num
}

// 递归
// 转为有序数组
func traversal1(root *TreeNode, resPtr *[]int) {
	if root == nil {
		return
	}
	if root.Left != nil {
		traversal1(root.Left, resPtr)
	}
	*resPtr = append(*resPtr, root.Val)
	if root.Right != nil {
		traversal1(root.Right, resPtr)
	}
}

// 递归
// 保存上一节点的值
func getMinimumDifference(root *TreeNode) int {
	var preNode *TreeNode
	min := math.MaxInt64
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left != nil {
			traversal(node.Left)
		}

		if preNode != nil && node.Val-preNode.Val < min {
			min = node.Val - preNode.Val
		}
		preNode = node
		if node.Right != nil {
			traversal(node.Right)
		}
	}
	traversal(root)
	return min
}
