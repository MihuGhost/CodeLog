package main

// 树的最大高度
func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 函数 返回形式
// 终止条件
// 单层逻辑
// 递归
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

// 层序遍历
func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0
	curLevel := []*TreeNode{root}
	for len(curLevel) > 0 {
		nextLevel := []*TreeNode{}
		for _, node := range curLevel {
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
		curLevel = nextLevel
		res++
	}
	return res
}

/*
	node := queue[len(queue)-1]
	queue = queue[:len(queue)-1]
	if node.Left != nil {
		queue = append(queue, node.Left)
	}
	if node.Right != nil {
		queue = append(queue, node.Right)
	}
*/
