package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

// 后序遍历回溯
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == q || root == p || root == nil {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if right != nil && left != nil {
		return root
	}
	if left != nil {
		return left
	}
	return right
}
