package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	rootVal := preorder[0]
	rootNode := &TreeNode{Val: rootVal}
	idx := findIdx(inorder, rootVal)
	rootNode.Left = buildTree(preorder[1:1+idx], inorder[:idx])
	rootNode.Right = buildTree(preorder[1+idx:], inorder[idx+1:])

	return rootNode
}

func findIdx(arr []int, val int) int {
	for i, v := range arr {
		if v == val {
			return i
		}
	}
	return -1
}
