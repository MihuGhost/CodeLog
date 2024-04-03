package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

//递归删除
func deleteNode1(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == key {
		if root.Left == nil && root.Right == nil {
			return nil
		} else if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		} else {
			var node, preNode *TreeNode
			node = root.Right
			for node != nil {
				preNode = node
				node = node.Left
			}
			preNode.Left = root.Left
			return root.Right
		}
	}
	if root.Val > key {
		root.Left = deleteNode1(root.Left, key)
	} else {
		root.Right = deleteNode1(root.Right, key)
	}
	return root
}

//迭代
func deleteNode(root *TreeNode, key int) *TreeNode {
	return nil
}
