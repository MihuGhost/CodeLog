package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

// 递归 子切片传入
func buildTree1(inorder []int, postorder []int) *TreeNode {
	var traversal func(subIn []int, subPost []int) *TreeNode
	traversal = func(subIn []int, subPost []int) *TreeNode {
		if len(subPost) == 0 {
			return nil
		}

		node := &TreeNode{Val: subPost[len(subPost)-1]}
		index := 0

		for _, num := range subIn {
			if num == node.Val {
				break
			}
			index++
		}

		node.Left = traversal(subIn[0:index], subPost[0:index])
		node.Right = traversal(subIn[index+1:], subPost[index:len(subPost)-1])
		return node
	}
	root := traversal(inorder, postorder)
	return root
}

// 递归 索引传入

func buildTree(inorder []int, postorder []int) *TreeNode {
	rec := make(map[int]int)
	for k, v := range inorder {
		rec[v] = k
	}
	root := traversal(postorder, 0, len(postorder)-1, 0, rec)
	return root
}

func traversal(postorder []int, postStart int, postEnd int, inStart int, rec map[int]int) *TreeNode {
	if postStart > postEnd {
		return nil
	}
	node := &TreeNode{Val: postorder[postEnd]}
	//根索引
	rootIdx := rec[node.Val]
	//左长度
	lenLeft := rootIdx - inStart

	//右边序列在变化
	node.Left = traversal(postorder, postStart, postStart+lenLeft-1, inStart, rec)
	node.Right = traversal(postorder, postStart+lenLeft, postEnd-1, rootIdx+1, rec)
	return node
}
