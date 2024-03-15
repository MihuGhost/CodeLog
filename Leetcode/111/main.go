package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	//二叉树最小深度
}
func minDepth(root *TreeNode) int {
	return getHeight(root, 1)
}

func getHeight(root *TreeNode, depth int) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return depth
	} else if root.Left == nil && root.Right != nil {
		return getHeight(root.Right, depth+1)
	} else if root.Right == nil && root.Left != nil {
		return getHeight(root.Left, depth+1)
	} else {
		return min(getHeight(root.Left, depth+1), getHeight(root.Right, depth+1))
	}
}
