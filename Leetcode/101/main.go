package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归比较二叉树
func compare(left *TreeNode, right *TreeNode) bool {
	//顺序-空指针报错
	/*	if left != nil && right == nil {
			return false
		} else if left == nil && right != nil {
			return false
		} else if left == nil && right == nil {
			return true
		} else if left.Val != right.Val {
			return false
		}*/
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil || left.Val != right.Val {
		return false
	}

	outside := compare(left.Left, right.Right)
	inside := compare(left.Right, right.Left)
	res := outside && inside
	return res
}

// 队列比较二叉树
func isSymmetric1(root *TreeNode) bool {
	var queue []*TreeNode
	if root != nil {
		queue = append(queue, root.Left, root.Right)
	}

	for len(queue) > 0 {
		left := queue[0]
		right := queue[1]
		queue = queue[2:]

		if left == nil && right == nil {
			continue
		}
		if left == nil || right == nil || left.Val != right.Val {
			return false
		}

		queue = append(queue, left.Left, right.Right, left.Right, right.Left)
	}

	return len(queue) == 0
}

func isSymmetric(root *TreeNode) bool {
	return compare(root.Left, root.Right)
}
