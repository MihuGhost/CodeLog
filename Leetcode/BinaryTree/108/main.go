package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func sortedArrayToBST(nums []int) *TreeNode {
	return convert(nums, 0, len(nums)-1)
}

func convert(nums []int, l, r int) *TreeNode {
	if l > r {
		return nil
	}
	mid := l + (r-l)/2
	root := &TreeNode{Val: nums[mid]}
	root.Left = convert(nums, l, mid-1)
	root.Right = convert(nums, mid+1, r)
	return root
}
