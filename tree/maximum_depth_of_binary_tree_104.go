package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
当前树的最大深度 = 左子树和右子树最大深度的较大值 + 当前节点这一层

	    root
	   /    \
	left    right

左边深度 = dfs(root.Left)
右边深度 = dfs(root.Right)

当前深度 = max(左边深度, 右边深度) + 1
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	maxLen := max(leftDepth, rightDepth) + 1
	return maxLen
}
