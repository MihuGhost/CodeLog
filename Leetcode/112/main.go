package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

// 递归 修改前
func hasPathSum1(root *TreeNode, targetSum int) bool {
	var getSum func(node *TreeNode, sum int)
	var res []int
	getSum = func(node *TreeNode, sum int) {
		if node.Left == nil && node.Right == nil {
			sum += node.Val
			res = append(res, sum)
		}
		sum += node.Val
		if node.Left != nil {
			getSum(node.Left, sum)
		}
		if node.Right != nil {
			getSum(node.Right, sum)
		}
	}
	if root == nil {
		return false
	}
	getSum(root, 0)
	return contains(res, targetSum)
}

func contains(numbers []int, num int) bool {
	for _, number := range numbers {
		if num == number {
			return true
		}
	}
	return false
}

//修改后

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	var getSum func(node *TreeNode, sum int) bool
	getSum = func(node *TreeNode, sum int) bool {
		if node.Left == nil && node.Right == nil {
			//sum -= node.Val
			return sum-node.Val == 0
		}

		if node.Left != nil {
			if getSum(node.Left, sum-node.Val) {
				return true
			}
		}
		if node.Right != nil {
			if getSum(node.Right, sum-node.Val) {
				return true
			}
		}
		return false
	}
	return getSum(root, targetSum)
}
