package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

// 递归
func searchBST1(root *TreeNode, val int) *TreeNode {
	var search func(node *TreeNode) *TreeNode
	search = func(node *TreeNode) *TreeNode {
		if node.Val == val {
			return node
		} else if node.Val > val && node.Left != nil {
			return search(node.Left)
		} else if node.Val < val && node.Right != nil {
			return search(node.Right)
		}
		return nil
	}
	return search(root)
}

// 迭代
func searchBST(root *TreeNode, val int) *TreeNode {
	for root != nil {
		if root.Val > val {
			root = root.Left
		} else if root.Val < val {
			root = root.Right
		} else {
			return root
		}
	}
	return nil
}
