package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < low {
		right := trimBST(root.Right, low, high)
		return right
	}
	if root.Val > high {
		left := trimBST(root.Left, low, high)
		return left
	}

	root.Right = trimBST(root.Right, low, high)
	root.Left = trimBST(root.Left, low, high)
	return root
}
