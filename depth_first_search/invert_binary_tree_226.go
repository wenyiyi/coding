package depth_first_search

/*
https://leetcode.com/problems/invert-binary-tree/
Given the root of a binary tree, invert the tree, and return its root.
Input:

        4
       / \
      2   7
     / \ / \
    1  3 6  9

Output:

        4
       / \
      7   2
     / \ / \
    9  6 3  1

Input: root = [4,2,7,1,3,6,9]
Output: [4,7,2,9,6,3,1]
*/

/*
invertTree(root) 的含义：
→ 翻转以 root 为根的整棵树，并返回 root

什么时候结束
root == nil → return nil

递归：
left  = 翻转后的左子树
right = 翻转后的右子树

当前节点：
root.Left  = right
root.Right = left

最后：
return root
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := invertTree(root.Left)
	right := invertTree(root.Right)
	root.Left, root.Right = right, left
	return root
}
