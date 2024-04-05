package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}
func convertBST(root *TreeNode) *TreeNode {
	pre := 0
	var trversal func(node *TreeNode)
	trversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		trversal(node.Right)
		node.Val += pre
		pre = node.Val
		trversal(node.Left)
	}
	trversal(root)
	return root
}
