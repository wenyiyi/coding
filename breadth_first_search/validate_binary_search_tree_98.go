package breadth_first_search

import "math"

/*
*
Given the root of a binary tree, determine if it is a valid binary search tree (BST).

A valid BST is defined as follows:

The left subtree of a node contains only nodes with keys strictly less than the node's key.
The right subtree of a node contains only nodes with keys strictly greater than the node's key.
Both the left and right subtrees must also be binary search trees.

	  5
	 / \
	3   7

左边 < 5，右边 > 5

	  5
	 / \
	3   7
	   / \
	  4   8

因为 4 在根节点 5 的右子树里，却4 < 5

所以这题不能只判断
node.Left.Val < node.Val
node.Right.Val > node.Val

  - Definition for a binary tree node.
  - type TreeNode struct {
  - Val int
  - Left *TreeNode
  - Right *TreeNode
  - }

每往下一层，都把祖先留下来的范围一起传下去
dfs(node, lower, upper)
检查以 node 为根的树，所有节点是否都满足 (lower, upper)
如果 node.Val <= lower → false
如果 node.Val >= upper → false

递归左子树：dfs(node.Left,  lower, node.Val)   范围 (lower,node.Val)  左子树：必须 < 当前节点，所以收紧上界
递归右子树：dfs(node.Right, node.Val, upper)   范围 (node.Val,upper)  右子树：必须 > 当前节点，所以收紧下界

空树应该认为是合法 BST

	if node == nil {
	    return true
	}
*/

/*
三问
① dfs(node) 到底返回什么？
② node == nil 返回什么？
③ 左右子树已经帮我算好了，我当前节点要做什么？

① dfs 返回：当前子树是不是合法 BST
② nil：true
③ 检查自己是否合法，然后让左右子树继续检查
*/
func isValidBST(root *TreeNode) bool {
	var dfs func(node *TreeNode, lower, upper int64) bool

	dfs = func(node *TreeNode, lower, upper int64) bool {
		// 最底层一定会返回true，空树是合法 BST
		if node == nil {
			return true
		}

		// true 返回到上一层后，上一层还要检查自己必须在 (lower, upper) 之间
		if int64(node.Val) <= lower || int64(node.Val) >= upper {
			return false
		}

		// 往左：范围 (lower,node.Val)  左子树：必须 < 当前节点，所以收紧上界
		// 往右：范围 (node.Val,upper)  右子树：必须 > 当前节点，所以收紧下界
		return dfs(node.Left, lower, int64(node.Val)) &&
			dfs(node.Right, int64(node.Val), upper)
	}

	// math.MinInt64 = -9223372036854775808
	// math.MaxInt64 =  9223372036854775807
	return dfs(root, math.MinInt64, math.MaxInt64)
}

/*
    2
   / \
  1   3

dfs(2)
  ↓
2 不是 nil → 检查 2
  ↓
dfs(1)
  ↓
1 不是 nil → 检查 1
  ↓
dfs(nil)
  ↓
这时候才 return true

那既然最底层一定 true，怎么最后还能得到 false
关键是：true 返回到上一层后，上一层还要检查自己
*/
