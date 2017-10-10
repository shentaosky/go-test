package main

func main() {
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	res := &[]int{}
	traversal(root, res)
	return *res
}

func traversal(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	traversal(root.Left, res)
	res = append(res, root.Val)
	traversal(root.Right, res)
}

func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	list := []*TreeNode{}
	for root != nil || len(list) != 0 {
		for root != nil {
			res = append(res, root.Val)
			list = append(list, root)
			root = root.Left
		}
		root = list[len(list)-1]
		list = list[:len(list)-1]
		root = root.Right
	}
	return res
}
