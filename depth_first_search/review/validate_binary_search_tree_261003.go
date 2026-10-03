package review

import "math"

/*
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

三问
① dfs(node) 到底返回什么？
② node == nil 返回什么？
③ 左右子树已经帮我算好了，我当前节点要做什么？

① dfs 返回：当前子树是不是合法 BST
② nil：true
③ 检查自己是否合法，然后让左右子树继续检查
*/

func isValidBST(root *TreeNode) bool {
	var dfs func(node *TreeNode, lower, upper int) bool
	dfs = func(node *TreeNode, lower, upper int) bool {
		if node == nil {
			return true
		}
		// todo 检查自己是否合法
		if node.Val < lower || node.Val > upper {
			return false
		}
		left := dfs(node.Left, lower, node.Val)
		right := dfs(node.Right, node.Val, upper)
		return left && right
	}

	// todo math.MinInt64      math.MaxInt64
	return dfs(root, math.MinInt64, math.MaxInt64)
}
