package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// 合并二叉树
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	var merge func(node1 *TreeNode, node2 *TreeNode) *TreeNode
	merge = func(node1 *TreeNode, node2 *TreeNode) *TreeNode {
		if node1 == nil {
			return node2
		}
		if node2 == nil {
			return node1
		}
		node1.Val += node2.Val
		node1.Left = merge(node1.Left, node2.Left)
		node1.Right = merge(node1.Right, node2.Right)
		return node1
	}
	return merge(root1, root2)
}
