package depth_first_search

/**
https://leetcode.com/problems/maximum-depth-of-binary-tree/description/
A binary tree's maximum depth is the number of nodes along the longest
path from the root node down to the farthest leaf node.

      3
   9    20
      15  7

Input: root = [3,9,20,null,null,15,7]
Output: 3

Input: root = [1,null,2]
Output: 2


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

/*
做题时，先回答这三个问题
1 dfs(node) 要完成什么？    -> 返回以 node 为根的树有多少层
2 什么时候直接结束？ -> 空节点返回 0
3 子问题解决后，我怎么得到答案？-> 左右深度取较大值，再加自己这一层
*/

func maxDepth(root *TreeNode) int {
	// 什么时候直接结束？ ->  空节点表示结束，返回 0
	if root == nil {
		return 0
	}
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	// 子问题解决后，我怎么得到答案？ ->  左右深度取较大值，再加自己这一层
	maxLen := max(leftDepth, rightDepth) + 1
	return maxLen
}
