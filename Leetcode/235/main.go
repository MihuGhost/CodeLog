package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor1(root.Right, p, q)
	} else if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor1(root.Left, p, q)
	} else {
		return root
	}
}

// 迭代
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	for root != nil {
		if root.Val > q.Val && root.Val > p.Val {
			root = root.Left
		} else if root.Val < q.Val && root.Val < p.Val {
			root = root.Right
		} else {
			return root
		}
	}
	return nil
}
