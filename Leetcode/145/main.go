package main

// Binary Tree Postorder Traversal
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func levelOrder(root *TreeNode) [][]int {
	var res [][]int

	if root == nil {
		return res
	}

	curLevel := []*TreeNode{root}

	for len(curLevel) > 0 {
		var subres []int
		var nextLevel []*TreeNode
		for _, node := range curLevel {
			subres = append(subres, node.Val)
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
		res = append(res, subres)
		curLevel = nextLevel
	}

	return res
}
